package core

type YmlConfigFile struct {
	PlainTextConfigFile
}

func NewYmlConfigFile(namespace string, configRepository ConfigRepository) ConfigFile {
	return &YmlConfigFile{
		PlainTextConfigFile: NewPlainTextConfigFile(namespace, configRepository),
	}
}

func (config *YmlConfigFile) GetConfigFileFormat() string {
	return YML
}

func (config *YmlConfigFile) ToProperties() *Properties{
	return &Properties{}
}
