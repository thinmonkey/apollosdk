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
	ConfitUtil core.ConfigUtil
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

/**
@param namespace
获取namespace配置
 */
func GetConfig(namespace string) core.Config {
	return ConfigManagerInstance.GetConfig(namespace)
}

/**
获取默认的配置
 */
func GetAppConfig() core.Config {
	return GetConfig(NAMESPACE_APPLICATION)
}

func GetConfigFile(namespace string,configFileFormat string) core.ConfigFile {
	return ConfigManagerInstance.GetConfigFile(namespace,configFileFormat)
}
