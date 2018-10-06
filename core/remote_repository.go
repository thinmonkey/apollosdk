package core

import (
	"sync"
	"time"
	"fmt"
	"net/url"
	"strings"
	"encoding/json"
	"github.com/thinmonkey/apollosdk/util/schedule"
	"github.com/thinmonkey/apollosdk/util/http"
	"github.com/thinmonkey/apollosdk/util"
)

type RemoteConfigRepository struct {
	AbstractConfigRepository
	Namespace                   string
	ApolloConfig                *ApolloConfig
	RemoteMessages              ApolloNotificationMessages
	LongPollServiceDto          ServiceDto
	schedulePolicy              schedule.ExponentialSchedulePolicy
	lock                        sync.Mutex
	ConfigNeedForceRefresh      bool
	remoteConfigLongPollService *RemoteConfigLongPollService
	configUtil                  util.ConfitUtil
}

func NewRemoteConfigRepository(Namespace string, configUtil util.ConfitUtil) *RemoteConfigRepository {
	remoteConfigRepository := &RemoteConfigRepository{
		Namespace:  Namespace,
		configUtil: configUtil,
	}
	remoteConfigRepository.remoteConfigLongPollService = NewRemoteConfigLongPollService(configUtil)
	remoteConfigRepository.schedulePolicy = schedule.NewExponentialSchedulePolicy(configUtil.HttpOnErrorRetryInterval, configUtil.HttpOnErrorRetryInterval*8)
	remoteConfigRepository.trySync()
	remoteConfigRepository.schedulePeriodicRefresh()
	remoteConfigRepository.scheduleLongPollingRefresh()
	return remoteConfigRepository
}

func (remoteConfigRepository *RemoteConfigRepository) GetConfig() *Properties {
	if remoteConfigRepository.ApolloConfig == nil {
		remoteConfigRepository.sync()
	}
	return remoteConfigRepository.transformApolloConfigToProperties(remoteConfigRepository.ApolloConfig)
}

func (remoteConfigRepository *RemoteConfigRepository) sync() {
	currentApolloConfig := remoteConfigRepository.loadApolloConfig()
	if currentApolloConfig != remoteConfigRepository.ApolloConfig {
		remoteConfigRepository.ApolloConfig = currentApolloConfig
		remoteConfigRepository.FireRepositoryChange(remoteConfigRepository.Namespace, remoteConfigRepository.GetConfig())
	}
}

func (remoteConfigRepository *RemoteConfigRepository) getSourceType() ConfigSourceType {
	return REMOTE
}

func (remoteConfigRepository *RemoteConfigRepository) transformApolloConfigToProperties(config *ApolloConfig) *Properties {
	result := Properties{}
	for key, value := range config.Configurations {
		result[key] = value
	}
	return &result
}

func (remoteConfigRepository *RemoteConfigRepository) loadApolloConfig() (*ApolloConfig) {

	appId := remoteConfigRepository.configUtil.AppId
	cluster := remoteConfigRepository.configUtil.Cluster
	dataCenter := remoteConfigRepository.configUtil.DataCenter
	var maxRetry int
	if remoteConfigRepository.ConfigNeedForceRefresh {
		maxRetry = 2
	} else {
		maxRetry = 1
	}

	var onErrorSleepTime time.Duration
	configServices := remoteConfigRepository.getConfigServices()

	for i := 0; i < maxRetry; i++ {
		for _, serviceDto := range configServices {
			if onErrorSleepTime > 0 {
				time.Sleep(onErrorSleepTime)
			}
			url := assembleQueryConfigUrl(serviceDto.HomePageUrl, appId, cluster, remoteConfigRepository.Namespace, dataCenter,
				remoteConfigRepository.RemoteMessages, remoteConfigRepository.ApolloConfig)
			httpRequest := http.HttpRequest{
				Url:            url,
				ConnectTimeout: remoteConfigRepository.configUtil.HttpTimeout,
			}
			httpResponse, err := http.Request(httpRequest)
			if err != nil {
				onErrorSleepTime = remoteConfigRepository.calErrorSleepTime()
				util.Logger.Error("loadApolloConfig http err:",err)
				continue
			}
			util.Logger.Infof("remote_repository response,statusCode:%d,body:%s,url: %s", httpResponse.StatusCode,httpResponse.ReponseBody,url)
			remoteConfigRepository.ConfigNeedForceRefresh = false
			remoteConfigRepository.schedulePolicy.Success()
			if httpResponse.StatusCode == 304 {
				return remoteConfigRepository.ApolloConfig
			}

			var newApolloConfig ApolloConfig
			err = json.Unmarshal(httpResponse.ReponseBody, &newApolloConfig)
			if err != nil {
				util.Logger.Error("loadApolloConfig http response json unmarshal ApolloConfig err:",err)
				continue
			}
			util.Logger.Infof("remote_repository request success:%s",newApolloConfig)
			return &newApolloConfig
		}
	}
	return nil
}

func (remoteConfigRepository *RemoteConfigRepository) calErrorSleepTime() time.Duration {
	if remoteConfigRepository.ConfigNeedForceRefresh {
		return remoteConfigRepository.configUtil.HttpOnErrorRetryInterval
	} else {
		return remoteConfigRepository.schedulePolicy.Fail()
	}
}

func assembleQueryConfigUrl(host string, appId string, cluster string, namespace string,
	dateCenter string, remoteMessages ApolloNotificationMessages, previousConfig *ApolloConfig) string {
	pathFormat := "configs/%s/%s/%s"
	path := fmt.Sprintf(pathFormat, appId, cluster, namespace)

	queryParam := ""
	if previousConfig != nil && previousConfig.ReleaseKey != "" {
		releaseQuery := "releaseKey=" + url.QueryEscape(previousConfig.ReleaseKey) + "&"
		queryParam = queryParam + releaseQuery
	}
	if dateCenter != "" {
		dataQuery := "dataCenter=" + url.QueryEscape(dateCenter) + "&"
		queryParam = queryParam + dataQuery
	}
	if remoteMessages.Details != nil {
		message, _ := json.Marshal(remoteMessages)
		messageQuery := "messages=" + url.QueryEscape(string(message)) + "&"
		queryParam = queryParam + messageQuery
	}
	if util.GetLocalIp() != "" {
		ipQuery := "ip=" + url.QueryEscape(util.GetLocalIp())
		queryParam = queryParam + ipQuery
	}
	if !strings.HasSuffix(host, "/") {
		host = host + "/"
	}
	if queryParam != "" {
		path = path + "?" + queryParam
	}
	httpPath := host + path
	rawUrl,_ := url.PathUnescape(httpPath)
	util.Logger.Infof("remote_repository request rawUrl:%s",rawUrl)
	return httpPath
}

func (remoteConfigRepository *RemoteConfigRepository) getConfigServices() []ServiceDto {
	configServiceLoad := NewConfigServiceLoad(remoteConfigRepository.configUtil)
	return configServiceLoad.ServiceDtoList
}

func (remoteConfigRepository *RemoteConfigRepository) scheduleLongPollingRefresh() {
	remoteConfigRepository.remoteConfigLongPollService.Submit(remoteConfigRepository.Namespace, remoteConfigRepository)
}

func (remoteConfigRepository *RemoteConfigRepository) schedulePeriodicRefresh() {
	util.ScheduleIntervalExecutor(remoteConfigRepository.configUtil.HttpRefreshInterval, remoteConfigRepository.trySync)
}

func (remoteConfigRepository *RemoteConfigRepository) trySync() {
	remoteConfigRepository.sync()
}

func (remoteConfigRepository *RemoteConfigRepository) onLongPollNotified(longPollNotifiedServiceDto ServiceDto, remoteMessages ApolloNotificationMessages) {
	remoteConfigRepository.LongPollServiceDto = longPollNotifiedServiceDto
	remoteConfigRepository.RemoteMessages = remoteMessages
	go func() {
		remoteConfigRepository.ConfigNeedForceRefresh = true
		remoteConfigRepository.trySync()
	}()
}
