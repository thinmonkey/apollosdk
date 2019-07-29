package core

type XmlConfigFile struct {
	PlainTextConfigFile
}

func NewXmlConfigFile(namespace string, configRepository ConfigRepository) ConfigFile {
	return &XmlConfigFile{
		PlainTextConfigFile: NewPlainTextConfigFile(namespace, configRepository),
	}
}

func (config *XmlConfigFile) GetConfigFileFormat() string {
	return XML
}
