/**
 * @author Created by zhanghao on 2019-07-30 11:24, [haozhang@ecarx.com.cn]
 */
package core

import "sync/atomic"

type PropertiesConfigFile struct {
	AbstractConfigFile
	content atomic.Value
}

func NewPropertiesConfigFile(namespace string, configRepository ConfigRepository) ConfigFile {
	return &PropertiesConfigFile{
		AbstractConfigFile:NewAbstractConfigFile(namespace,configRepository),
	}
}

func (config *PropertiesConfigFile) update(newProperties Properties) {
	config.AbstractConfigFile.update(newProperties)
	config.content.Store(nil)
}

func (config *PropertiesConfigFile) GetContent() string {
	if config.content.Load() == nil {
		config.content.Store(config.doContent())
	}
	return config.content.Load().(string)
}

func (config *PropertiesConfigFile) doContent() string {
	if !config.HasContent() {
		return ""
	}
	return config.Properties.ToString()
}

func (config *PropertiesConfigFile) HasContent() bool {
	return config.Properties != nil && len(config.Properties) != 0
}

func (config *PropertiesConfigFile) GetConfigFileFormat() string {
	return PROPERTIES
}
