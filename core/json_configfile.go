package core

type JsonConfigFile struct {
	PlainTextConfigFile
}

func (config *JsonConfigFile) GetConfigFileFormat() string {
	return JSON
}



