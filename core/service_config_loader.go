package core

import (
	"net/url"
	"strings"
	"github.com/thinmonkey/apollosdk/util/http"
	"encoding/json"
	"sync"
	"github.com/thinmonkey/apollosdk/util"
)

var once sync.Once
var configServiceLoader *ConfigServiceLoader

type ConfigServiceLoader struct {
	ServiceDtoList []ServiceDto
	configUtil     util.ConfitUtil
}

func NewConfigServiceLoad(configUtil util.ConfitUtil) *ConfigServiceLoader {
	once.Do(func() {
		configServiceLoader = &ConfigServiceLoader{
			configUtil: configUtil,
		}
		configServiceLoader.tryUpdateConfigServices()
		configServiceLoader.schedulePeriodicRefresh()
	})
	return configServiceLoader
}

func (serviceLoader *ConfigServiceLoader) tryUpdateConfigServices() {
	serviceLoader.updateConfigServices()
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

	httpReponse, err := http.Request(httpRequest)
	if err != nil {
		util.Logger.Error(err)
		return
	}
	if httpReponse.StatusCode == 200 && httpReponse.ReponseBody != nil {
		var serviceConfig = make([]ServiceDto, 1)
		err := json.Unmarshal(httpReponse.ReponseBody, &serviceConfig)
		if err != nil {
			util.Logger.Error("json unmarshal err ", err)
		}
		serviceLoader.setConfigServices(serviceConfig)
	}

	util.Logger.Infof("Get service config response,statusCode:%d,body:%s,url: %s", httpReponse.StatusCode,httpReponse.ReponseBody,url)
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
	httpPath := host+path
	rawUrl,_ := url.PathUnescape(httpPath)
	util.Logger.Infof("service_config request rawUrl:%s",rawUrl)
	return httpPath
}
