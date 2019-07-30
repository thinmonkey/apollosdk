package apollosdk

import (
	"fmt"
	"github.com/thinmonkey/apollosdk/core"
	"sync"
)

var ConfigManagerInstance ConfigManager
var onceConfigManager sync.Once

type ConfigManager interface {
	GetConfig(namespace string) core.Config
	GetConfigFile(namespace string, configFileFormat string) core.ConfigFile
}

type DefaultConfigManager struct {
	configs     map[string]core.Config
	configFiles map[string]core.ConfigFile

	lock sync.Mutex
}

func GetConfigManager() ConfigManager {
	onceConfigManager.Do(func() {
		ConfigManagerInstance = NewConfigManager()
	})
	return ConfigManagerInstance
}

func NewConfigManager() ConfigManager {
	return &DefaultConfigManager{
		configs:     map[string]core.Config{},
		configFiles: map[string]core.ConfigFile{},
	}
}

func (manager *DefaultConfigManager) GetConfig(namespace string) core.Config {
	if config, ok := manager.configs[namespace]; ok {
		return config
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	config := core.NewDefaultConfig(namespace, core.NewRemoteConfigRepository(namespace, ConfitUtil), ConfitUtil)
	if config != nil {
		manager.configs[namespace] = config
	}
	return config
}

func (manager *DefaultConfigManager) GetConfigFile(namespace string, configFileFormat string) core.ConfigFile {
	namespaceFileName := fmt.Sprintf("%s.%s", namespace, configFileFormat)

	if configFile, ok := manager.configFiles[namespaceFileName]; ok {
		return configFile
	}

	manager.lock.Lock()
	defer manager.lock.Unlock()

	configFile := GetConfigFactory().CreateConfigFile(namespaceFileName, configFileFormat)
	if configFile != nil {
		manager.configFiles[namespaceFileName] = configFile
	}

	return configFile
}
