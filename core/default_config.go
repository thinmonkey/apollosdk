package core

import (
	"os"
	"github.com/zhhao226/apollosdk/util"
)

type DefaultConfig struct {
	AbstractConfig
	Properties       *Properties
	ConfigRepository ConfigRepository
	Namespace        string
	SourceType       ConfigSourceType
}

func NewDefaultConfig(nameSpace string, configReposity *ConfigRepository, configUtil util.ConfitUtil) *DefaultConfig {
	defaultConfig := DefaultConfig{
		Namespace:        nameSpace,
		ConfigRepository: *configReposity,
		Properties:       (*configReposity).GetConfig(),
	}
	defaultConfig.AbstractConfig = AbstractConfig{
		configChangeListeners: make([]ConfigChangeListener, 0),
		InterestKeyMap:        make(map[ConfigChangeListener][]string, 0),
		GetProperty:           defaultConfig.GetDefaultProterty,
		configUtil:            configUtil,
	}

	configChangeListener := RepositoryChangeListener(&defaultConfig)

	defaultConfig.ConfigRepository.AddChangeListener(&configChangeListener)
	return &defaultConfig
}

func (defaultConfig *DefaultConfig) GetDefaultProterty(key string, defaultValue string) []byte {
	value := defaultConfig.Properties.getProperty(key)
	if value == "" {
		value = os.Getenv(key)
	}
	if value == "" {
		value = defaultValue
	}
	return []byte(value)
}

/**
 * Return a set of the property names
 *
 * @return the property names
 */
func (defaultConfig *DefaultConfig) GetPropertyNames() []string {
	return []string{}
}

func (defaultConfig *DefaultConfig) OnRepositoryChange(namespace string, newProperties *Properties) () {
	if defaultConfig.Properties == newProperties {
		return
	}

	actualChanges := defaultConfig.updateAndCalcConfigChanges(*newProperties, defaultConfig.ConfigRepository.GetSourceType())
	if actualChanges == nil || len(actualChanges) == 0 {
		return
	}
	defaultConfig.fireConfigChange(ConfigChangeEvent{namespace, actualChanges})
}

func (defaultConfig *DefaultConfig) updateConfig(newProperties Properties, sourceType ConfigSourceType) {
	defaultConfig.Properties = &newProperties
	defaultConfig.SourceType = sourceType
}

func (defaultConfig *DefaultConfig) updateAndCalcConfigChanges(properties Properties, sourceType ConfigSourceType) map[string]ConfigChange {
	configChanges := defaultConfig.calcPropertyChanges(defaultConfig.Namespace, *defaultConfig.Properties, properties)

	for _, change := range configChanges {
		change.OldValue = string(defaultConfig.GetDefaultProterty(change.PropertyName, change.OldValue))
	}

	actualChanges := make(map[string]ConfigChange, 10)

	defaultConfig.updateConfig(properties, sourceType)
	defaultConfig.clearConfigCache()

	for _, change := range configChanges {
		change.NewValue = string(defaultConfig.GetDefaultProterty(change.Namespace, change.NewValue))
		switch change.ChangeType {
		case ADDED:
			if change.NewValue == change.OldValue {
				continue
			}
			if change.OldValue != "" {
				change.ChangeType = MODIFIED
			}
			actualChanges[change.PropertyName] = change
		case MODIFIED:
			if change.OldValue != change.NewValue {
				actualChanges[change.PropertyName] = change
			}
		case DELETED:
			if change.OldValue == change.NewValue {
				continue
			}
			if change.NewValue != "" {
				change.ChangeType = MODIFIED
			}
			actualChanges[change.PropertyName] = change
		}
	}
	return actualChanges
}
