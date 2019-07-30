package core

import (
	"encoding/json"
	"github.com/thinmonkey/apollosdk/util"
	"github.com/thinmonkey/apollosdk/util/http"
	"net/url"
	"strings"
	"sync"
)

var once sync.Once
var configServiceLoader *ConfigServiceLoader

type ConfigServiceLoader struct {
	ServiceDtoList []ServiceDto
	configUtil     ConfigUtil
}

func NewConfigServiceLoad(configUtil ConfigUtil) *ConfigServiceLoader {
	once.Do(func() {
		configServiceLoader = &ConfigServiceLoader{
			configUtil: configUtil,
		}
		configServiceLoader.tryUpdateConfigServices()
		configServiceLoader.schedulePeriodicRefresh()
	})
	return configServiceLoader
}

func (serviceLoader *ConfigServiceLoader) tryUpdateConfigServices() bool {
	serviceLoader.updateConfigServices()
	return true
}

func (serviceLoader *ConfigServiceLoader) schedulePeriodicRefresh() {
	util.ScheduleIntervalExecutor(serviceLoader.configUtil.HttpRefreshInterval, serviceLoader.tryUpdateConfigServices)
}

func (serviceLoader *ConfigServiceLoader) updateConfigServices() {
	url := serviceLoader.assembleQueryConfigUrl(serviceLoader.configUtil.MetaServer, serviceLoader.configUtil.AppId)

	httpRequest := http.HttpRequest{
		Url:            url,
		ConnectTimeout: serviceLoader.configUtil.HttpTimeout,
	}

	httpResponse, err := http.Request(httpRequest)
	if err != nil {
		util.DebugPrintf("updateConfigServices http err %v ", err)
		return
	}
	if httpResponse.StatusCode == 200 && httpResponse.ReponseBody != nil {
		var serviceConfig = make([]ServiceDto, 1)
		err := json.Unmarshal(httpResponse.ReponseBody, &serviceConfig)
		if err != nil {
			util.DebugPrintf("json unmarshal err：%v ", err)
		}
		serviceLoader.setConfigServices(serviceConfig)
	}

	util.DebugPrintf("Get service config response,statusCode:%d,body:%s,url: %s", httpResponse.StatusCode, httpResponse.ReponseBody, url)
}

func (serviceLoader *ConfigServiceLoader) setConfigServices(serviceDtoList []ServiceDto) {
	serviceLoader.ServiceDtoList = serviceDtoList
}

func (serviceLoader *ConfigServiceLoader) assembleQueryConfigUrl(host string, appId string) string {
	path := "services/config"

	queryParam := ""
	if appId := serviceLoader.configUtil.AppId; appId != "" {
		appIdQuery := "appId=" + url.QueryEscape(appId) + "&"
		queryParam = queryParam + appIdQuery
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
	rawUrl, _ := url.PathUnescape(httpPath)
	util.DebugPrintf("service_config request rawUrl:%s", rawUrl)
	return httpPath
}
