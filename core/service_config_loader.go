package core

import (
	"net/url"
	"strings"
	"github.com/zhhao226/apollosdk/util/http"
	"encoding/json"
	"time"
	"sync"
	"github.com/zhhao226/apollosdk/util"
)

var once sync.Once
var configServiceLoader *ConfigServiceLoader

type ConfigServiceLoader struct {
	ServiceDtoList []ServiceDto
}

func NewConfigServiceLoad() *ConfigServiceLoader {
	once.Do(func() {
		configServiceLoader = &ConfigServiceLoader{}
		configServiceLoader.tryUpdateConfigServices()
		configServiceLoader.schedulePeriodicRefresh()
	})
	return configServiceLoader
}

func (serviceLoader *ConfigServiceLoader) tryUpdateConfigServices() {
	serviceLoader.updateConfigServices()
}

func (serviceLoader *ConfigServiceLoader) schedulePeriodicRefresh() {
	go func() {
		t2 := time.NewTimer(util.RefreshInterval)
		//long poll for sync
		for {
			select {
			case <-t2.C:
				serviceLoader.tryUpdateConfigServices()
				t2.Reset(util.RefreshInterval)
			}
		}
	}()
}

func (serviceLoader *ConfigServiceLoader) updateConfigServices() {
	url := serviceLoader.assembleQueryConfigUrl(util.GetMetaServer(), util.GetAppId())

	httpRequest := http.HttpRequest{
		Url:            url,
		ConnectTimeout: util.ConnectTimeout,
	}

	httpReponse, err := http.Request(httpRequest)
	if err != nil {
		util.Logger.Error(err)
	}
	if httpReponse.StatusCode == 200 && httpReponse.ReponseBody != nil {
		var serviceConfig = make([]ServiceDto, 1)
		err := json.Unmarshal(httpReponse.ReponseBody, &serviceConfig)
		if err != nil {
			util.Logger.Error("json unmarshal err ", err)
		}
		serviceLoader.setConfigServices(serviceConfig)
	}

	util.Logger.Debugf("Get service config response: %s, url: %s", httpReponse.StatusCode, url)
}

func (serviceLoader *ConfigServiceLoader) setConfigServices(serviceDtoList []ServiceDto) {
	serviceLoader.ServiceDtoList = serviceDtoList
}

func (serviceLoader *ConfigServiceLoader) assembleQueryConfigUrl(host string, appId string) string {
	path := "services/config"

	queryParam := ""
	if util.GetAppId() != "" {
		appIdQuery := "appId=" + url.QueryEscape(util.GetAppId()) + "&"
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
	util.Logger.Info(host, path)
	return host + path
}
