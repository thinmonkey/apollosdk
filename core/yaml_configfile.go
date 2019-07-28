package core

type YamlConfigFile struct {
	AbstractConfigFile
}

func (config *YamlConfigFile) GetConfigFileFormat() string {
	return YAML
}
