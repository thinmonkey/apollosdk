package core

type ConfigRepository interface {
	/**
	 * Get the config from this repository.
	 * @return config
	 */
	GetConfig() *Properties

	/**
	 * Add change listener.
	 * @param listener the listener to observe the changes
	 */
	AddChangeListener(listener *RepositoryChangeListener)

	/**
	 * Remove change listener.
	 * @param listener the listener to remove
	 */
	RemoveChangeListener(listener *RepositoryChangeListener)

	/**
   * Return the config's source type, i.e. where is the config loaded from
   *
   * @return the config's source type
   */
	getSourceType() ConfigSourceType
}
