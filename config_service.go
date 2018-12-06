package apollosdk

import (
	"github.com/thinmonkey/apollosdk/core"
	"github.com/thinmonkey/apollosdk/util"
	"sync"
)

const (
	NAMESPACE_APPLICATION       = "application"
	CLUSTER_NAME_DEFAULT        = "default"
	CLUSTER_NAMESPACE_SEPARATOR = "+"
)

var (
	configMap  = make(map[string]*core.Config, 10)
	lock       sync.Mutex
	ConfitUtil core.ConfitUtil
	once       sync.Once
)

//启动默认配置
func init() {
	util.SetDebug(false)
	//默认的配置
	ConfitUtil = core.NewConfigWithConfigFile("config.json")
}

func SetDebug(debug bool) {
	util.SetDebug(debug)
}

//启动动态配置
func Start(appId string, cluster string, metaServer string, dataCenter string) {
	once.Do(func() {
		ConfitUtil = core.NewConfigWithApolloInitConfig(core.ApolloInitConfig{
			AppId:      appId,
			Cluster:    cluster,
			MetaServer: metaServer,
			DataCenter: dataCenter,
		})
	})
}

//自定义配置文件进行配置
func StartWithCusConfig(configFile string) {
	once.Do(func() {
		ConfitUtil = core.NewConfigWithConfigFile(configFile)
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
