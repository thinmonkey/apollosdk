package core

import (
	"gopkg.in/yaml.v2"
	"log"
)

type YamlConfigFile struct {
	PlainTextConfigFile
	Properties *Properties
}

func NewYamlConfigFile(namespace string, configRepository ConfigRepository) ConfigFile {
	yamlConfigFile := &YamlConfigFile{
		PlainTextConfigFile: NewPlainTextConfigFile(namespace, configRepository),
	}
	yamlConfigFile.tryTransformToProperties()
	return yamlConfigFile
}

func (config *YamlConfigFile) GetConfigFileFormat() string {
	return YAML
}

func (config *YamlConfigFile) tryTransformToProperties() bool {
	config.transformToProperties()
	return true
}

func (config *YamlConfigFile) update(newProperties *Properties) {
	config.PlainTextConfigFile.update(newProperties)
	config.transformToProperties()
}

func (config *YamlConfigFile) transformToProperties() {
	config.Properties = config.toProperties()
}

func (config *YamlConfigFile) toProperties() *Properties {
	if !config.HasContent() {
		return &Properties{}
	}
	var properties *Properties
	err := yaml.Unmarshal([]byte(config.GetContent()), properties)
	if err != nil {
		log.Printf("Parse yaml file content failed for namespace:%err", err)
	}
	return properties
}

func (config *YamlConfigFile) AsProperties() *Properties {
	if config == nil {
		config.transformToProperties()
	}
	return config.Properties
}
