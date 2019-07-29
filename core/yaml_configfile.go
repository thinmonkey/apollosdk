package core

type YamlConfigFile struct {
	PlainTextConfigFile
}

func NewYamlConfigFile(namespace string, configRepository ConfigRepository) ConfigFile {
	return &YamlConfigFile{
		PlainTextConfigFile: NewPlainTextConfigFile(namespace, configRepository),
	}
}

func (config *YamlConfigFile) GetConfigFileFormat() string {
	return YAML
}
