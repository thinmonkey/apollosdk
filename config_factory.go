package apollosdk

import (
	"fmt"
	"github.com/thinmonkey/apollosdk/core"
	"strings"
	"sync"
)

var ConfigFactoryInstance ConfigFactory
var onceConfigFactory sync.Once

type ConfigFactory interface {
	CreateConfig(namespace string) core.Config
	CreateConfigFile(namespace string, configFileFormat string) core.ConfigFile
}

type DefaultConfigFactory struct {
	manager ConfigManager
}

func NewConfigFactory() ConfigFactory {
	return &DefaultConfigFactory{
		manager: NewConfigManager(),
	}
}

func GetConfigFactory() ConfigFactory {
	onceConfigFactory.Do(func() {
		ConfigFactoryInstance = NewConfigFactory()
	})
	return ConfigFactoryInstance
}

func (*DefaultConfigFactory) CreateConfig(namespace string) core.Config {
	defaultConfig := core.NewDefaultConfig(namespace, core.NewRemoteConfigRepository(namespace, ConfitUtil), ConfitUtil)
	return defaultConfig
}

func (*DefaultConfigFactory) CreateConfigFile(namespace string, configFileFormat string) core.ConfigFile {
	configRepository := core.NewRemoteConfigRepository(namespace, ConfitUtil)
	switch configFileFormat {
	case core.JSON:
		return core.NewJsonConfigFile(namespace, configRepository)
	case core.XML:
		return core.NewXmlConfigFile(namespace, configRepository)
	case core.TXT:
		return core.NewTxtConfigFile(namespace, configRepository)
	case core.YAML:
		return core.NewYamlConfigFile(namespace, configRepository)
	case core.YML:
		return core.NewYmlConfigFile(namespace, configRepository)
	}
	return nil
}

func (*DefaultConfigFactory) determineFileFormat(namespaceName string) string {
	lowerCase := strings.ToLower(namespaceName)
	for _, format := range core.FILE_FORMAT {
		if strings.HasSuffix(lowerCase, fmt.Sprintf(".%s", format)) {
			return format
		}
	}
	return core.PROPERTIES
}

func (*DefaultConfigFactory) trimNamespaceFormat(namespaceName string, format string) string {
	extension := fmt.Sprintf(".%s", format)
	if !strings.HasSuffix(strings.ToLower(namespaceName), extension) {
		return namespaceName
	}
	return namespaceName[:len(namespaceName)-len(extension)]
}
