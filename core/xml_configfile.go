package core

type XmlConfigFile struct {
	AbstractConfigFile
}

func (config *XmlConfigFile) GetConfigFileFormat() string {
	return XML
}
