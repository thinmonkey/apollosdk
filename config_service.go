package apollosdk

import (
	"github.com/zhhao226/apollosdk/core"
	"sync"
)

const (
	NAMESPACE_APPLICATION       = "application"
	CLUSTER_NAME_DEFAULT        = "default"
	CLUSTER_NAMESPACE_SEPARATOR = "+"
)

var configMap = make(map[string]*core.Config, 10)
var lock sync.Mutex

func GetConfig(namespace string) core.Config {
	lock.Lock()
	defer lock.Unlock()
	config, ok := configMap[namespace]

	if !ok {
		remoteRepository := core.NewRemoteConfigRepository(namespace)

		repository := core.ConfigRepository(remoteRepository)

		defaultConfig := core.NewDefaultConfig(namespace, &repository)

		config := core.Config(defaultConfig)
		configMap[namespace] = &config
		return config
	}
	return *config

}

func GetAppConfig() core.Config {
	return GetConfig(NAMESPACE_APPLICATION)
}
