package apollosdk

import (
	"time"
	"os"
	"net"
)

var (
	RefreshInterval      = 5 * time.Minute
	ConnectTimeout       = 30 * time.Second
	ReadTimeout          = 5 * time.Second
	Cluster              = "default"
	LoadConfigQPS        = 2 * time.Second //2 times per second
	LongPollQPS          = 2 * time.Second //2 times per second
	OnErrorRetryInterval = 1 * time.Second

	MaxConfigCacheSize             = 50 * 1024 * 1024
	ConfigCacheExpireTime          = 1 * 60
	LongPollingInitialDelayInMills = 2 * time.Second
)

func GetAppId() string {
	//return os.Getenv("apollo.appId")
	return "app-capability"
}

func GetCluster() string {
	//return os.Getenv("apollo.Cluster")
	return Cluster
}

func GetApolloEnv() string {
	return os.Getenv("docker_env")
}

func GetDateServer() string {
	return os.Getenv("docker_env")
}

func GetMetaServer() string {
	return "http://10.160.2.153:8083"
	//return os.Getenv("docker_server")
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
