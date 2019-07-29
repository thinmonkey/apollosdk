package core

//PlainTextConfigFile
type PlainTextConfigFile struct {
	AbstractConfigFile
}

func NewPlainTextConfigFile(namespace string, configRepository ConfigRepository) PlainTextConfigFile {
	return PlainTextConfigFile{
		AbstractConfigFile: NewAbstractConfigFile(namespace, configRepository),
	}
}

/**
  * Get file content of the namespace
  * @return file content, {@code null} if there is no content
  */
func (config *PlainTextConfigFile) GetContent() string {
	if !config.HasContent() {
		return ""
	}
	return config.Properties.getProperty(CONFIG_FILE_CONTENT_KEY)
}

func (config *PlainTextConfigFile) GetNamespace() string {
	return config.Namespace
}

func (config *PlainTextConfigFile) GetConfigFileFormat() string {
	panic("implement me")
}

/**
 * Whether the config file has any content
 * @return true if it has content, false otherwise.
 */
func (config *PlainTextConfigFile) HasContent() bool {
	if config.Properties == nil {
		return false
	}
	return config.Properties.Contain(CONFIG_FILE_CONTENT_KEY)
}

func (config *PlainTextConfigFile)  update(newProperties *Properties) {
	config.rwMutex.Lock()
	defer config.rwMutex.Unlock()
	config.Properties = newProperties
}