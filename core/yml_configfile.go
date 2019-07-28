package core

type YmlConfigFile struct {
	AbstractConfigFile
}

func (config *YmlConfigFile) GetConfigFileFormat() string {
	return YML
}
