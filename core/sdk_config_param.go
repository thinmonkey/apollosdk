package core

import (
	"github.com/sirupsen/logrus"
	"time"
	"os"
	"io/ioutil"
	"encoding/json"
)

type ConfitUtil struct {
	ApolloInitConfig
	CacheInitConfig
	HttpRefreshInterval      time.Duration
	HttpTimeout              time.Duration
	HttpOnErrorRetryInterval time.Duration
	LongPollingInitDelay     time.Duration
	LongPollingTimeout       time.Duration
	configStartFile          map[string]interface{}
}

type ApolloInitConfig struct {
	AppId      string
	Cluster    string
	DataCenter string
	MetaServer string
}

type CacheInitConfig struct {
	MaxConfigCacheSize    int
	ConfigCacheExpireTime int
}

func NewConfigUtil(configFile string, appId string, cluster string, metaServer string, dataCenter string) ConfitUtil {
	configUtil := ConfitUtil{
		HttpRefreshInterval:      5 * time.Minute,
		HttpTimeout:              30 * time.Second,
		HttpOnErrorRetryInterval: 1 * time.Second,
		LongPollingInitDelay:     2 * time.Second,
		LongPollingTimeout:       60 * time.Second,
		CacheInitConfig: CacheInitConfig{
			MaxConfigCacheSize:    50 * 1024 * 1024,
			ConfigCacheExpireTime: 1 * 60,
		},
		ApolloInitConfig: ApolloInitConfig{
			AppId:      appId,
			Cluster:    cluster,
			MetaServer: metaServer,
			DataCenter: dataCenter,
		},
	}
	configUtil.LoadConfigFile(configFile)
	return configUtil
}

func (util *ConfitUtil) LoadConfigFile(filename string) {
	if filename == "" {
		filename = "config.json"
	}
	fs, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error("Fail to find config file:" + err.Error())
		return
	}
	util.configStartFile = make(map[string]interface{}, 10)
	err = json.Unmarshal(fs, &util.configStartFile)
	if err != nil {
		logrus.Error("Fail to read json config file:" + err.Error())
		return
	}
	initRefreshTime(util)
	initHttpTimeout(util)
	initErrorRetry(util)
	initCacheExpireTime(util)
	initMaxCacheSize(util)
	initLongPollInitDelay(util)
	initLongpollTimeout(util)
	initAppId(util)
	initCluster(util)
	initDataServer(util)
	initMetaServer(util)
}
func initMetaServer(util *ConfitUtil) {
	//优先选择用户运行时代码设置的
	if util.MetaServer != "" {
		return
	}
	//其次选择系统环境变量配置
	metaCenter := os.Getenv("DOCKER_SERVER")
	if metaCenter != "" {
		util.MetaServer = metaCenter
		return
	}
	//最后选择配置文件配置
	metaCenter, _ = util.configStartFile["metaServer"].(string)
	if metaCenter != "" {
		util.MetaServer = metaCenter
	}
}

func initDataServer(util *ConfitUtil) {
	//优先选择用户运行时代码设置的
	if util.DataCenter != "" {
		return
	}
	//其次选择系统环境变量配置
	util.DataCenter = os.Getenv("apollo.dataCenter")
}

func initCluster(util *ConfitUtil) {
	//优先选择用户运行时代码设置的
	if util.Cluster != "" {
		return
	}
	//其次选择系统环境变量配置
	cluster := os.Getenv("apollo.Cluster")
	if cluster != "" {
		util.Cluster = cluster
		return
	}
	//最后选择配置文件配置
	cluster, _ = util.configStartFile["cluster"].(string)
	if cluster != "" {
		util.Cluster = cluster
		return
	}
}

func initAppId(util *ConfitUtil) {
	//优先选择用户运行时代码设置的
	if util.AppId != "" {
		return
	}
	//其次选择系统环境变量配置
	appId := os.Getenv("apollo.appId")
	if appId != "" {
		util.AppId = appId
		return
	}
	//最后选择配置文件配置
	appId, _ = util.configStartFile["appId"].(string)
	if appId != "" {
		util.AppId = appId
		return
	}
}

func initLongpollTimeout(util *ConfitUtil) {
	longPollingTimeout, _ := util.configStartFile["longPollingTimeout"].(string)
	if longPollingTimeout != "" {
		util.LongPollingTimeout, _ = time.ParseDuration(longPollingTimeout)
	}
}

func initLongPollInitDelay(util *ConfitUtil) {
	longPollingInitDelay, _ := util.configStartFile["longPollingInitDelay"].(string)
	if longPollingInitDelay != "" {
		util.LongPollingInitDelay, _ = time.ParseDuration(longPollingInitDelay)
	}
}

func initMaxCacheSize(util *ConfitUtil) {
	maxConfigCacheSize, _ := util.configStartFile["maxConfigCacheSize"].(float64)
	if maxConfigCacheSize != 0 {
		util.MaxConfigCacheSize = int(maxConfigCacheSize)
	}
}

func initCacheExpireTime(util *ConfitUtil) {
	configCacheExpireTime, _ := util.configStartFile["configCacheExpireTime"].(float64)
	if configCacheExpireTime != 0 {
		util.ConfigCacheExpireTime = int(configCacheExpireTime)
	}
}

func initErrorRetry(util *ConfitUtil) {
	onErrorRetryInterval, _ := util.configStartFile["onErrorRetryInterval"].(string)
	if onErrorRetryInterval != "" {
		util.HttpOnErrorRetryInterval, _ = time.ParseDuration(onErrorRetryInterval)
	}
}

func initHttpTimeout(util *ConfitUtil) {
	connectTimeout, _ := util.configStartFile["httpTimeout"].(string)
	if connectTimeout != "" {
		util.HttpTimeout, _ = time.ParseDuration(connectTimeout)
	}
}

func initRefreshTime(util *ConfitUtil) {
	refreshInterval, _ := util.configStartFile["httpRefreshInterval"].(string)
	if refreshInterval != "" {
		util.HttpRefreshInterval, _ = time.ParseDuration(refreshInterval)
	}
}
