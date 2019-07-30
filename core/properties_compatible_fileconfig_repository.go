/**
 * @author Created by zhanghao on 2019-07-30 15:04, [haozhang@ecarx.com.cn]
 */
package core

import (
	"log"
	"reflect"
)

type PropertiesCompatibleFileConfigRepository struct {
	AbstractConfigRepository
	configFile PropertiesCompatibleConfigFile
	cachedProperties Properties
}

func NewPropertiesCompatibleFileConfigRepository(configFile PropertiesCompatibleConfigFile) *PropertiesCompatibleFileConfigRepository {
	p :=  &PropertiesCompatibleFileConfigRepository{configFile: configFile}
	p.configFile.AddChangeListener(p)
	p.trySync()
	return p
}

func (p *PropertiesCompatibleFileConfigRepository) sync() {
	current := p.configFile.AsProperties()

	if current == nil {
		log.Printf("PropertiesCompatibleConfigFile.AsProperties should never return null")
		return
	}

	if !reflect.DeepEqual(current,p.cachedProperties) {
		p.cachedProperties = current
		p.FireRepositoryChange(p.configFile.GetNamespace(),p.cachedProperties)
	}
}

func (p *PropertiesCompatibleFileConfigRepository) GetConfig() Properties {
	if p.cachedProperties == nil {
		p.sync()
	}
	return p.cachedProperties
}

func (p *PropertiesCompatibleFileConfigRepository) GetSourceType() ConfigSourceType {
	return p.configFile.GetSourceType()
}

func (p *PropertiesCompatibleFileConfigRepository) OnChange(changeEvent ConfigFileChangeEvent) {
	p.trySync()
}
