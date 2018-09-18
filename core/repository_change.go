package core

type RepositoryChangeListener interface {
	/**
   * Invoked when config repository changes.
   * @param namespace the namespace of this repository change
   * @param newProperties the properties after change
   */
	OnRepositoryChange(namespace string, newProperties *Properties)
}
