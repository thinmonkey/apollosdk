package core

type Config interface {
	//GetProperty(key string, defaultValue string) []byte
	//GetLongProperty(key string, defaultValue int64) int64
	GetStringProperty(key string, defaultValue string) string
	GetIntProperty(key string, defaultValue int) int
	GetFloatProperty(key string, defaultValue float32) float32
	GetDoubleProperty(key string, defaultValue float64) float64
	//GetByteProperty(key string, defaultValue byte) byte
	GetBoolProperty(key string, defaultValue bool) bool
	GetArrayProperty(key string, delimiter string, defaultValue []string) []string
	/**
	   * Return the Date property value with the given name, or {@code defaultValue} if the name doesn't exist.
	   * Will try to parse the date with Locale.US and formats as follows: yyyy-MM-dd HH:mm:ss.SSS,
	   * yyyy-MM-dd HH:mm:ss and yyyy-MM-dd
	   *
	   * @param key          the property name
	   * @param defaultValue the default value when name is not found or any error occurred
	   * @return the property value
	   */
	//GetDateProperty(key string, defaultValue time.Time) time.Time
	//GetDateFormatProperty(key string, format string, defaultValue time.Time) time.Time
	//GetDateFormatLocaleProperty(key string, format string, locale time.Location, defalutValue time.Time) time.Time

	/**
	 * Return the duration property value(in milliseconds) with the given name, or {@code
	 * defaultValue} if the name doesn't exist. Please note the format should comply with the follow
	 * example (case insensitive). Examples:
	 * <pre>
	 *    "123MS"          -- parses as "123 milliseconds"
	 *    "20S"            -- parses as "20 seconds"
	 *    "15M"            -- parses as "15 minutes" (where a minute is 60 seconds)
	 *    "10H"            -- parses as "10 hours" (where an hour is 3600 seconds)
	 *    "2D"             -- parses as "2 days" (where a day is 24 hours or 86400 seconds)
	 *    "2D3H4M5S123MS"  -- parses as "2 days, 3 hours, 4 minutes, 5 seconds and 123 milliseconds"
	 * </pre>
	 *
	 * @param key          the property name
	 * @param defaultValue the default value when name is not found or any error occurred
	 * @return the parsed property value(in milliseconds)
	 */
	//GetDurationProperty(key string, defaultValue int64) int64


	//GetChangeKeyNotify() <-chan ConfigChangeEvent

	AddChangeListener(listener ConfigChangeListener)

	AddChangeListenerFunc(listenerFunc OnChangeFunc)

	AddChangeListenerInterestedKeys(listener ConfigChangeListener, interestedKeys []string)

	AddChangeListenerFuncInterestedKeys(listenerFunc OnChangeFunc, interestedKeys []string)

	//GetChangeInterestedKeysNotify(interestedKeys []string) <-chan ConfigChangeEvent

	/**
	 * Remove the change listener
	 *
	 * @param listener the specific config change listener to remove
	 * @return true if the specific config change listener is found and removed
	 */
	RemoveChangeListener(listener ConfigChangeListener) bool



	/**
	 * Return a set of the property names
	 *
	 * @return the property names
	 */
	GetPropertyNames() []string
}
