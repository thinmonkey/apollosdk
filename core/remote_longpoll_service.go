package core

import (
	"time"
	"github.com/zhhao226/apollosdk/util/schedule"
	"sync"
	"math/rand"
	"github.com/zhhao226/apollosdk/util/http"
	"encoding/json"
	"net/url"
	"strings"
	"fmt"
	"github.com/zhhao226/apollosdk/util"
)

const (
	INIT_NOTIFICATION_ID = -1
)

type RemoteConfigLongPollService struct {
	longPollServiceStarted     bool
	longPollServiceStoped      bool
	schedulePolicy             schedule.ExponentialSchedulePolicy
	longPollNamespace          map[string]*RemoteConfigRepository
	notifications              map[string]int64
	remoteNotificationMessages map[string]*ApolloNotificationMessages
	configUtil                 util.ConfitUtil
	sync.RWMutex
}

func NewRemoteConfigLongPollService(configUtil util.ConfitUtil) *RemoteConfigLongPollService {
	return &RemoteConfigLongPollService{
		schedulePolicy:             schedule.NewExponentialSchedulePolicy(configUtil.HttpOnErrorRetryInterval, configUtil.HttpOnErrorRetryInterval*8),
		longPollNamespace:          make(map[string]*RemoteConfigRepository, 8),
		notifications:              make(map[string]int64, 8),
		remoteNotificationMessages: make(map[string]*ApolloNotificationMessages, 8),
		configUtil:                 configUtil,
	}
}

func (remoteConfigLongPollService *RemoteConfigLongPollService) Submit(Namespace string, repository *RemoteConfigRepository) {
	remoteConfigLongPollService.Lock()
	remoteConfigLongPollService.longPollNamespace[Namespace] = repository
	remoteConfigLongPollService.notifications[Namespace] = INIT_NOTIFICATION_ID
	if !remoteConfigLongPollService.longPollServiceStarted {
		remoteConfigLongPollService.startLongPoll()
	}
	remoteConfigLongPollService.Unlock()
}

func (remoteConfigLongPollService *RemoteConfigLongPollService) startLongPoll() {
	appId := remoteConfigLongPollService.configUtil.AppId
	cluster := remoteConfigLongPollService.configUtil.Cluster
	dataCenter := remoteConfigLongPollService.configUtil.DataCenter
	longPollingInitialDelayInMills := remoteConfigLongPollService.configUtil.LongPollingRefreshInterval
	go func() {
		if longPollingInitialDelayInMills > 0 {
			time.Sleep(time.Duration(longPollingInitialDelayInMills))
		}
		remoteConfigLongPollService.doLongPollingRefresh(appId, cluster, dataCenter)
	}()
}

func (remoteConfigLongPollService *RemoteConfigLongPollService) doLongPollingRefresh(appId string, cluster string, dataCenter string) {
	var lastServiceDto *ServiceDto
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		if !remoteConfigLongPollService.longPollServiceStoped {
			if lastServiceDto == nil {
				serviceDtos := remoteConfigLongPollService.getConfigServices()
				lastServiceDto = &serviceDtos[rand.Intn(len(serviceDtos))]
			}
			url := assembleLongPollRefreshUrl(lastServiceDto.HomePageUrl, appId, cluster, dataCenter,
				remoteConfigLongPollService.notifications)

			httpRequest := http.HttpRequest{
				Url:            url,
				ConnectTimeout: remoteConfigLongPollService.configUtil.LongPollingTimeout,
			}

			httpReponse, err := http.Request(httpRequest)
			if err != nil {
				util.Logger.Error(err)
				lastServiceDto = nil
				sleepTime := remoteConfigLongPollService.schedulePolicy.Fail()
				time.Sleep(sleepTime)
				continue
			}
			util.Logger.Debugf("Long polling response: %s, url: %s", httpReponse.StatusCode, url);
			if httpReponse.StatusCode == 200 && httpReponse.ReponseBody != nil {
				var apolloNotifications []ApolloConfigNotification
				err := json.Unmarshal(httpReponse.ReponseBody, &apolloNotifications)
				if err != nil {
					util.Logger.Error("json unmarshal err ", err)
					lastServiceDto = nil
					sleepTime := remoteConfigLongPollService.schedulePolicy.Fail()
					time.Sleep(sleepTime)
					continue
				}
				remoteConfigLongPollService.updateNotifications(apolloNotifications)
				remoteConfigLongPollService.updateRemoteNotifications(apolloNotifications)
				remoteConfigLongPollService.notify(lastServiceDto, apolloNotifications)
			}

			if httpReponse.StatusCode == 304 {
				lastServiceDto = nil
			}

			remoteConfigLongPollService.schedulePolicy.Success()
		}
	}
}
func assembleLongPollRefreshUrl(host string, appId string, cluster string, dataCenter string, notifications map[string]int64) string {

	queryParam := ""
	if appId != "" {
		releaseQuery := "appId=" + url.QueryEscape(appId) + "&"
		queryParam = queryParam + releaseQuery
	}
	if cluster != "" {
		dataQuery := "cluster=" + url.QueryEscape(cluster) + "&"
		queryParam = queryParam + dataQuery
	}
	if notifications != nil {
		notificationList := make([]ApolloConfigNotification, 0)
		index := 0
		for key, value := range notifications {
			notificationList = append(notificationList, ApolloConfigNotification{NamespaceName: key, NotificationId: value})
			index++
		}
		notifications, err := json.Marshal(notificationList)
		if err != nil {
			util.Logger.Error(err)
		}
		util.Logger.Info(string(notifications))
		notificationsQuery := "notifications=" + url.QueryEscape(string(notifications)) + "&"
		queryParam = queryParam + notificationsQuery
	}
	if dataCenter != "" {
		messageQuery := "dataCenter=" + url.QueryEscape(dataCenter) + "&"
		queryParam = queryParam + messageQuery
	}
	if util.GetLocalIp() != "" {
		ipQuery := "ip=" + url.QueryEscape(util.GetLocalIp())
		queryParam = queryParam + ipQuery
	}
	if !strings.HasSuffix(host, "/") {
		host = host + "/"
	}
	url := host + "notifications/v2?" + queryParam
	util.Logger.Info(url)
	return url
}

func (remoteConfigLongPollService *RemoteConfigLongPollService) getConfigServices() []ServiceDto {
	configServiceLoad := NewConfigServiceLoad(remoteConfigLongPollService.configUtil)
	return configServiceLoad.ServiceDtoList
}

func (remoteConfigLongPollService *RemoteConfigLongPollService) updateNotifications(messages []ApolloConfigNotification) {
	for _, value := range messages {
		if value.NamespaceName == "" {
			continue
		}
		namespace := value.NamespaceName
		if _, ok := remoteConfigLongPollService.notifications[namespace]; ok {
			remoteConfigLongPollService.notifications[namespace] = value.NotificationId
		}
		namespaceNameWithPropertiesSuffix :=
			fmt.Sprintf("%s.%s", namespace, PROPERTIES)
		if _, ok := remoteConfigLongPollService.notifications[namespaceNameWithPropertiesSuffix]; ok {
			remoteConfigLongPollService.notifications[namespaceNameWithPropertiesSuffix] = value.NotificationId
		}
	}
}

func (remoteConfigLongPollService *RemoteConfigLongPollService) updateRemoteNotifications(messages []ApolloConfigNotification) {
	for _, value := range messages {
		if value.NamespaceName == "" {
			continue
		}
		namespace := value.NamespaceName
		if len(value.Messages.Details) == 0 {
			continue
		}
		if _, ok := remoteConfigLongPollService.remoteNotificationMessages[namespace]; !ok {
			notificationMessage := new(ApolloNotificationMessages)
			notificationMessage.Details = make(map[string]int64, 0)
			remoteConfigLongPollService.remoteNotificationMessages[namespace] = notificationMessage
		}
		remoteConfigLongPollService.remoteNotificationMessages[namespace].MergeFrom(value.Messages)
	}
}

func (remoteConfigLongPollService *RemoteConfigLongPollService) notify(serviceDto *ServiceDto, messages []ApolloConfigNotification) {
	if messages == nil {
		return
	}
	for _, notification := range messages {
		namespaceName := notification.NamespaceName
		remoteRepositoryList := make([]*RemoteConfigRepository, 0)
		remoteRepositoryList = append(remoteRepositoryList, remoteConfigLongPollService.longPollNamespace[namespaceName])

		originalMessages := remoteConfigLongPollService.remoteNotificationMessages[namespaceName]
		var remoteMessages ApolloNotificationMessages

		if originalMessages != nil && originalMessages.Details != nil {
			remoteMessages.Details = originalMessages.Details
		}

		nameSpaceSuffix := fmt.Sprintf("%s.%s", namespaceName, PROPERTIES)
		remoteRepositorySuffix := remoteConfigLongPollService.longPollNamespace[nameSpaceSuffix]
		if remoteRepositorySuffix != nil {
			remoteRepositoryList = append(remoteRepositoryList, remoteRepositorySuffix)
		}
		for _, remoteRepository := range remoteRepositoryList {
			remoteRepository.onLongPollNotified(remoteRepository.LongPollServiceDto, remoteMessages)
		}
	}
}
