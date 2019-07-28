package core

type TxtConfigFile struct {
	AbstractConfigFile
}

func (config *TxtConfigFile) GetConfigFileFormat() string {
	return TXT
}

