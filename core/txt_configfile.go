package core

type TxtConfigFile struct {
	PlainTextConfigFile
}

func NewTxtConfigFile(namespace string, configRepository ConfigRepository) ConfigFile {
	return &TxtConfigFile{
		PlainTextConfigFile: NewPlainTextConfigFile(namespace, configRepository),
	}
}

func (config *TxtConfigFile) GetConfigFileFormat() string {
	return TXT
}
