package core

type JsonConfigFile struct {
	PlainTextConfigFile
}

func NewJsonConfigFile(namespace string, configRepository ConfigRepository) ConfigFile {
	return &JsonConfigFile{
		PlainTextConfigFile: NewPlainTextConfigFile(namespace, configRepository),
	}
}

func (config *JsonConfigFile) GetConfigFileFormat() string {
	return JSON
}



