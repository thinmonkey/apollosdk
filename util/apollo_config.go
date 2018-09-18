package util

import (
	"time"
	"os"
	"net"
	"io/ioutil"
	"encoding/json"
)

var (
	RefreshInterval      = 5 * time.Minute
	ConnectTimeout       = 30 * time.Second
	OnErrorRetryInterval = 1 * time.Second

	MaxConfigCacheSize             = 50 * 1024 * 1024
	ConfigCacheExpireTime          = 1 * 60
	LongPollingInitialDelayInMills = 2 * time.Second

	configStartFile = make(map[string]interface{}, 10)
)

func init() {
	fs, err := ioutil.ReadFile("config.properties")
	if err != nil {
		Logger.Info("Fail to find config file:" + err.Error())
	}
	err = json.Unmarshal(fs, &configStartFile)
	if err != nil {
		Logger.Info("Fail to read json config file:" + err.Error())
	}
	refreshInterval := configStartFile["refreshInterval"].(string)
	if refreshInterval != "" {
		RefreshInterval, _ = time.ParseDuration(refreshInterval + "s")
	}
	connectTimeout := configStartFile["connectTimeout"].(string)
	if connectTimeout != "" {
		ConnectTimeout, _ = time.ParseDuration(connectTimeout + "s")
	}
	onErrorRetryInterval := configStartFile["onErrorRetryInterval"].(string)
	if onErrorRetryInterval != "" {
		OnErrorRetryInterval, _ = time.ParseDuration(onErrorRetryInterval + "s")
	}
	configCacheExpireTime := configStartFile["configCacheExpireTime"].(float64)
	if configCacheExpireTime != 0 {
		ConfigCacheExpireTime = int(configCacheExpireTime)
	}
	maxConfigCacheSize := configStartFile["maxConfigCacheSize"].(float64)
	if maxConfigCacheSize != 0 {
		MaxConfigCacheSize = int(maxConfigCacheSize)
	}
	longPollingInitialDelayInMills := configStartFile["longPollingInitialDelayInMills"].(string)
	if longPollingInitialDelayInMills != "" {
		LongPollingInitialDelayInMills, _ = time.ParseDuration(longPollingInitialDelayInMills + "s")
	}
}

func GetAppId() string {
	appId := os.Getenv("apollo.appId")
	if appId != "" {
		return appId
	}
	appId = configStartFile["appId"].(string)
	if appId != "" {
		return appId
	}
	return "application"
}

func GetCluster() string {
	cluster := os.Getenv("apollo.Cluster")
	if cluster != "" {
		return cluster
	}
	cluster = configStartFile["cluster"].(string)
	if cluster != "" {
		return cluster
	}
	return "default"
}

func GetDateCenter() string {
	return os.Getenv("apollo.dataCenter")
}

func GetMetaServer() string {
	meteCenter := os.Getenv("DOCKER_SERVER")
	if meteCenter != "" {
		return meteCenter
	}
	return configStartFile["metaServer"].(string)
}

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				internalIp := ipnet.IP.To4().String()
				return internalIp
			}
		}
	}
	return ""
}
