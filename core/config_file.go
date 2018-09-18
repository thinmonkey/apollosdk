package core

const (
	PROPERTIES = "properties"
	XML        = "xml"
	JSON       = "json"
	YML        = "yml"
	YAML       = "yaml"
)

type ConfigFile interface {
	/**
	  * Get file content of the namespace
	  * @return file content, {@code null} if there is no content
	  */
	GetContent() string

	/**
	 * Whether the config file has any content
	 * @return true if it has content, false otherwise.
	 */
	HasContent() bool

	/**
	 * Get the namespace of this config file instance
	 * @return the namespace
	 */
	GetNamespace() string

	/**
	 * Get the file format of this config file instance
	 * @return the config file format enum
	 */
	GetConfigFileFormat() string

	/**
	 * Add change listener to this config file instance.
	 *
	 * @param listener the config file change listener
	 */

	//AddChangeListener(listener ConfigFileChangeListener)
}
