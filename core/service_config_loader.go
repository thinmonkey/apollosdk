package core

import (
	"net/url"
	"strings"
	"github.com/zhhao226/apollosdk"
	"github.com/zhhao226/apollosdk/util/http"
	"encoding/json"
	"time"
	"sync"
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
		t2 := time.NewTimer(apollosdk.RefreshInterval)
		//long poll for sync
		for {
			select {
			case <-t2.C:
				serviceLoader.tryUpdateConfigServices()
				t2.Reset(apollosdk.RefreshInterval)
			}
		}
	}()
}

func (serviceLoader *ConfigServiceLoader) updateConfigServices() {
	url := serviceLoader.assembleQueryConfigUrl(apollosdk.GetMetaServer(), apollosdk.GetAppId())

	httpRequest := http.HttpRequest{
		Url:            url,
		ConnectTimeout: apollosdk.ConnectTimeout,
	}

	httpReponse, err := http.Request(httpRequest)
	if err != nil {
		apollosdk.Logger.Error(err)
	}
	if httpReponse.StatusCode == 200 && httpReponse.ReponseBody != nil {
		var serviceConfig = make([]ServiceDto, 1)
		err := json.Unmarshal(httpReponse.ReponseBody, &serviceConfig)
		if err != nil {
			apollosdk.Logger.Error("json unmarshal err ", err)
		}
		serviceLoader.setConfigServices(serviceConfig)
	}

	apollosdk.Logger.Debugf("Get service config response: %s, url: %s", httpReponse.StatusCode, url)
}

func (serviceLoader *ConfigServiceLoader) setConfigServices(serviceDtoList []ServiceDto) {
	serviceLoader.ServiceDtoList = serviceDtoList
}

func (serviceLoader *ConfigServiceLoader) assembleQueryConfigUrl(host string, appId string) string {
	path := "services/config"

	queryParam := ""
	if apollosdk.GetAppId() != "" {
		appIdQuery := "appId=" + url.QueryEscape(apollosdk.GetAppId()) + "&"
		queryParam = queryParam + appIdQuery
	}
	if apollosdk.GetLocalIp() != "" {
		ipQuery := "ip=" + url.QueryEscape(apollosdk.GetLocalIp())
		queryParam = queryParam + ipQuery
	}
	if !strings.HasSuffix(host, "/") {
		host = host + "/"
	}
	if queryParam != "" {
		path = path + "?" + queryParam
	}
	apollosdk.Logger.Info(host, path)
	return host + path
}
