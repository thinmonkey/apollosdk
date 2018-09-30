package apollosdk

import (
	"github.com/thinmonkey/apollosdk/core"
	"sync"
	"github.com/thinmonkey/apollosdk/util"
)

const (
	NAMESPACE_APPLICATION       = "application"
	CLUSTER_NAME_DEFAULT        = "default"
	CLUSTER_NAMESPACE_SEPARATOR = "+"
)

var (
	configMap  = make(map[string]*core.Config, 10)
	lock       sync.Mutex
	ConfitUtil util.ConfitUtil
	once       sync.Once
)

//启动默认配置
func init() {
	ConfitUtil = util.NewConfigUtil("config.json", "", "", "", "")
}

//启动动态配置
func Start(configFilename string, appId string, cluster string, metaServer string, dataCenter string) {
	once.Do(func() {
		ConfitUtil = util.NewConfigUtil(configFilename, appId, cluster, metaServer, dataCenter)
	})
}

func GetConfig(namespace string) core.Config {
	lock.Lock()
	defer lock.Unlock()
	config, ok := configMap[namespace]

	if !ok {
		remoteRepository := core.NewRemoteConfigRepository(namespace, ConfitUtil)

		repository := core.ConfigRepository(remoteRepository)

		defaultConfig := core.NewDefaultConfig(namespace, repository, ConfitUtil)

		config := core.Config(defaultConfig)
		configMap[namespace] = &config
		return config
	}
	return *config

}

func GetAppConfig() core.Config {
	return GetConfig(NAMESPACE_APPLICATION)
}
