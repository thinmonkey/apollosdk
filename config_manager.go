package apollosdk

import "github.com/thinmonkey/apollosdk/core"

type ConfigManager interface {
	GetConfig(namespace string) core.Config
	GetConfigFile(namespace string,configSourceType core.ConfigSourceType) core.ConfigFile
}

type DefaultConfigManager struct {
	configs map[string]core.Config
	configFiles map[string]core.ConfigFile
}

func (*DefaultConfigManager) GetConfig(namespace string) core.Config {

}

func (*DefaultConfigManager) GetConfigFile(namespace string, configSourceType core.ConfigSourceType) core.ConfigFile {
	panic("implement me")
}

