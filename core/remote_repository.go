package core

import (
	"sync"
	"time"
	"fmt"
	"net/url"
	"strings"
	"encoding/json"
	"github.com/zhhao226/apollosdk/util/schedule"
	"github.com/zhhao226/apollosdk/util/http"
	"github.com/zhhao226/apollosdk/util"
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
}

func NewRemoteConfigRepository(Namespace string) *RemoteConfigRepository {
	remoteConfigRepository := &RemoteConfigRepository{
		Namespace: Namespace,
	}
	remoteConfigRepository.remoteConfigLongPollService = NewRemoteConfigLongPollService()
	remoteConfigRepository.schedulePolicy = schedule.NewExponentialSchedulePolicy(util.OnErrorRetryInterval, util.OnErrorRetryInterval*8)
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
	currentApolloConfig, _ := remoteConfigRepository.loadApolloConfig()
	if currentApolloConfig != remoteConfigRepository.ApolloConfig {
		remoteConfigRepository.ApolloConfig = currentApolloConfig
		remoteConfigRepository.FireRepositoryChange(remoteConfigRepository.Namespace, remoteConfigRepository.GetConfig())
	}
}

func (remoteConfigRepository *RemoteConfigRepository) GetSourceType() ConfigSourceType {
	return REMOTE
}

func (remoteConfigRepository *RemoteConfigRepository) transformApolloConfigToProperties(config *ApolloConfig) *Properties {
	result := Properties{}
	for key, value := range config.Configurations {
		result[key] = value
	}
	return &result
}

func (remoteConfigRepository *RemoteConfigRepository) loadApolloConfig() (*ApolloConfig, error) {

	appId := util.GetAppId()
	cluster := util.GetCluster()
	dataServer := util.GetDateCenter()
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
			url := assembleQueryConfigUrl(serviceDto.HomePageUrl, appId, cluster, remoteConfigRepository.Namespace, dataServer,
				remoteConfigRepository.RemoteMessages, remoteConfigRepository.ApolloConfig)
			httpRequest := http.HttpRequest{
				Url:            url,
				ConnectTimeout: util.ConnectTimeout,
			}
			httpResponse, err := http.Request(httpRequest)
			if err != nil {
				onErrorSleepTime = remoteConfigRepository.calErrorSleepTime()
				return nil, err
			}
			remoteConfigRepository.ConfigNeedForceRefresh = false
			remoteConfigRepository.schedulePolicy.Success()
			if httpResponse.StatusCode == 304 {
				return remoteConfigRepository.ApolloConfig, nil
			}

			var newApolloConfig ApolloConfig
			err = json.Unmarshal(httpResponse.ReponseBody, &newApolloConfig)
			if err != nil {
				return nil, err
			}
			return &newApolloConfig, nil
		}
	}
	return nil, nil
}

func (remoteConfigRepository *RemoteConfigRepository) calErrorSleepTime() time.Duration {
	if remoteConfigRepository.ConfigNeedForceRefresh {
		return util.OnErrorRetryInterval
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
	util.Logger.Info(host, path)
	return host + path
}

func (remoteConfigRepository *RemoteConfigRepository) getConfigServices() []ServiceDto {
	configServiceLoad := NewConfigServiceLoad()
	return configServiceLoad.ServiceDtoList
}

func (remoteConfigRepository *RemoteConfigRepository) scheduleLongPollingRefresh() {
	remoteConfigRepository.remoteConfigLongPollService.Submit(remoteConfigRepository.Namespace, remoteConfigRepository)
}

func (remoteConfigRepository *RemoteConfigRepository) schedulePeriodicRefresh() {
	go func() {
		t2 := time.NewTimer(util.RefreshInterval)
		//long poll for sync
		for {
			select {
			case <-t2.C:
				remoteConfigRepository.trySync()
				t2.Reset(util.RefreshInterval)
			}
		}
	}()
}

func (remoteConfigRepository *RemoteConfigRepository) trySync() bool {
	remoteConfigRepository.sync()
	return true
}

func (remoteConfigRepository *RemoteConfigRepository) onLongPollNotified(longPollNotifiedServiceDto ServiceDto, remoteMessages ApolloNotificationMessages) {
	remoteConfigRepository.LongPollServiceDto = longPollNotifiedServiceDto
	remoteConfigRepository.RemoteMessages = remoteMessages
	go func() {
		remoteConfigRepository.ConfigNeedForceRefresh = true
		remoteConfigRepository.trySync()
	}()
}
