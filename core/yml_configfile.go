package core

type YmlConfigFile struct {
	YamlConfigFile
}

func NewYmlConfigFile(namespace string, configRepository ConfigRepository) ConfigFile {
	yamlConfigFile := NewYamlConfigFile(namespace, configRepository).(*YamlConfigFile)
	return &YmlConfigFile{
		YamlConfigFile: *yamlConfigFile,
	}
}

func (config *YmlConfigFile) GetConfigFileFormat() string {
	return YML
}
