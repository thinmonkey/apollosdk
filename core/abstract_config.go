package core

import (
	"github.com/coocood/freecache"
	"github.com/thinmonkey/apollosdk/util"
	"github.com/thinmonkey/apollosdk/util/set"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type AbstractConfig struct {
	ConfigVersion int64 `json:"configVersion"`
	cache         *freecache.Cache
	rwMutex       sync.Mutex
	//configChangeListeners []chan ConfigChangeEvent
	//InterestKeyMap        map[chan ConfigChangeEvent][]string
	GetProperty func(key string, defaultValue string) []byte

	//add changeListener
	configChangeListeners []ConfigChangeListener
	InterestKeyMap        map[ConfigChangeListener][]string
	configUtil            ConfitUtil
}

/**
 * Return the property value with the given key, or {@code defaultValue} if the key doesn't exist.
 *
 * @param key          the property name
 * @param defaultValue the default value when key is not found or any error occurred
 * @return the property value
 */
//func (config *AbstractConfig) GetProperty {
//	return ""
//}

/**
 * Return the long property value with the given key, or {@code defaultValue} if the key doesn't
 * exist.
 *
 * @param key          the property name
 * @param defaultValue the default value when key is not found or any error occurred
 * @return the property value as long
 */
//func (config *AbstractConfig) GetLongProperty(key string, defaultValue int64) int64 {
//	if config.cache == nil {
//		config.cache = newCache()
//	}
//	result := config.getValueFromCache(key)
//	if result != nil {
//		return strconv.ParseInt()
//	}
//
//	return config.getValueFromCache(key, defaultValue)
//}

/**
 * Return the property value with the given key, or {@code defaultValue} if the key doesn't exist.
 *
 * @param key          the property name
 * @param defaultValue the default value when key is not found or any error occurred
 * @return the property value
 */
func (config *AbstractConfig) GetStringProperty(key string, defaultValue string) string {
	if config.cache == nil {
		config.cache = config.newCache()
	}
	result := config.getValueFromCache(key)
	if result != nil && len(result) > 0 {
		value := string(result)
		return value
	}

	return defaultValue
}

/**
 * Return the integer property value with the given key, or {@code defaultValue} if the key
 * doesn't exist.
 *
 * @param key          the property name
 * @param defaultValue the default value when key is not found or any error occurred
 * @return the property value as integer
 */
func (config *AbstractConfig) GetIntProperty(key string, defaultValue int) int {
	if config.cache == nil {
		config.cache = config.newCache()
	}
	result := config.getValueFromCache(key)
	if result != nil && len(result) > 0 {
		value, err := strconv.Atoi(string(result))
		if err != nil {
			util.DebugPrintf("string convert to int err:", err)
		}
		return value
	}

	return defaultValue
}

/**
 * Return the float property value with the given key, or {@code defaultValue} if the key doesn't
 * exist.
 *
 * @param key          the property name
 * @param defaultValue the default value when key is not found or any error occurred
 * @return the property value as float
 */
func (config *AbstractConfig) GetFloatProperty(key string, defaultValue float32) float32 {
	if config.cache == nil {
		config.cache = config.newCache()
	}
	result := config.getValueFromCache(key)
	if result != nil && len(result) > 0 {
		value, err := strconv.ParseFloat(string(result), 32)
		if err != nil {
			util.DebugPrintf("string convert to float32 err:", err)
		}
		return float32(value)
	}

	return defaultValue
}

func (config *AbstractConfig) GetPropertyNames() []string {
	return []string{}
}

/**
 * Return the double property value with the given key, or {@code defaultValue} if the key doesn't
 * exist.
 *
 * @param key          the property name
 * @param defaultValue the default value when key is not found or any error occurred
 * @return the property value as double
 */
func (config *AbstractConfig) GetDoubleProperty(key string, defaultValue float64) float64 {
	if config.cache == nil {
		config.cache = config.newCache()
	}
	result := config.getValueFromCache(key)
	if result != nil && len(result) > 0 {
		value, err := strconv.ParseFloat(string(result), 64)
		if err != nil {
			util.DebugPrintf("string convert to float32 err:", err)
		}
		return value
	}

	return defaultValue
}

/**
 * Return the byte property value with the given key, or {@code defaultValue} if the key doesn't
 * exist.
 *
 * @param key          the property name
 * @param defaultValue the default value when key is not found or any error occurred
 * @return the property value as byte
 */
//func (config *AbstractConfig) GetByteProperty(key string, defaultValue byte) byte {
//	if config.cache == nil {
//		config.cache = newCache()
//	}
//	return 0
//}

/**
 * Return the boolean property value with the given key, or {@code defaultValue} if the key
 * doesn't exist.
 *
 * @param key          the property name
 * @param defaultValue the default value when key is not found or any error occurred
 * @return the property value as boolean
 */
func (config *AbstractConfig) GetBoolProperty(key string, defaultValue bool) bool {
	if config.cache == nil {
		config.cache = config.newCache()
	}
	result := config.getValueFromCache(key)
	if result != nil && len(result) > 0 {
		value, err := strconv.ParseBool(string(result))
		if err != nil {
			util.DebugPrintf("string convert to bool err:", err)
		}
		return value
	}

	return defaultValue
}

/**
 * Return the array property value with the given key, or {@code defaultValue} if the key doesn't exist.
 *
 * @param key          the property name
 * @param delimiter    the delimiter regex
 * @param defaultValue the default value when key is not found or any error occurred
 */
func (config *AbstractConfig) GetArrayProperty(key string, delimiter string, defaultValue []string) []string {
	if config.cache == nil {
		config.cache = config.newCache()
	}
	result := config.getValueFromCache(key)
	if result != nil && len(result) > 0 {
		value := strings.Split(string(result), delimiter)
		return value
	}

	return defaultValue
}

/**
 * Return the Date property value with the given name, or {@code defaultValue} if the name doesn't exist.
 * Will try to parse the date with Locale.US and formats as follows: yyyy-MM-dd HH:mm:ss.SSS,
 * yyyy-MM-dd HH:mm:ss and yyyy-MM-dd
 *
 * @param key          the property name
 * @param defaultValue the default value when name is not found or any error occurred
 * @return the property value
 */
//func (config *AbstractConfig) GetDateProperty(key string, defaultValue time.Time) time.Time {
//	if config.cache == nil {
//		config.cache = newCache()
//	}
//	return time.Time{}
//}

/**
 * Return the Date property value with the given name, or {@code defaultValue} if the name doesn't exist.
 *
 * @param key          the property name
 * @param format       the date format, see {@link java.text.SimpleDateFormat} for more
 *                     information
 * @param locale       the locale to use
 * @param defaultValue the default value when name is not found or any error occurred
 * @return the property value
 */
//func (config *AbstractConfig) GetDateFormatProperty(key string, format string, defaultValue time.Time) time.Time {
//	if config.cache == nil {
//		config.cache = newCache()
//	}
//	return time.Time{}
//}

func (config *AbstractConfig) newCache() *freecache.Cache {
	cache := freecache.NewCache(config.configUtil.MaxConfigCacheSize)
	return cache
}

func (config *AbstractConfig) clearConfigCache() {
	config.rwMutex.Lock()
	defer config.rwMutex.Unlock()
	if config.cache != nil {
		config.cache.Clear()
	} else {
		config.newCache()
	}
	config.ConfigVersion = config.ConfigVersion + 1

}

/**
 * Return the Date property value with the given name, or {@code defaultValue} if the name doesn't exist.
 *
 * @param key          the property name
 * @param format       the date format, see {@link java.text.SimpleDateFormat} for more
 *                     information
 * @param locale       the locale to use
 * @param defaultValue the default value when name is not found or any error occurred
 * @return the property value
 */
//func (config *AbstractConfig) GetDateFormatLocaleProperty(key string, format string, locale time.Location, defalutValue time.Time) time.Time {
//	if config.cache == nil {
//		config.cache = newCache()
//	}
//	return time.Time{}
//}

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
//func (config *AbstractConfig) GetDurationProperty(key string, defaultValue int64) int64 {
//	if config.cache == nil {
//		config.cache = newCache()
//	}
//	return 0
//}

/**
* Add change listener to this config instance, will be notified when any key is changed in this namespace.
*
* @param listener the config change listener
 */
func (config *AbstractConfig) AddChangeListener(listener ConfigChangeListener) {
	if config.cache == nil {
		config.cache = config.newCache()
	}
	config.AddChangeListenerInterestedKeys(listener, nil)
}

func (config *AbstractConfig) AddChangeListenerFunc(listenerFunc OnChangeFunc) {
	config.AddChangeListener(listenerFunc)
}

//func (config *AbstractConfig) GetChangeKeyNotify() <-chan ConfigChangeEvent {
//	if config.cache == nil {
//		config.cache = newCache()
//	}
//	config.rwMutex.Lock()
//	defer config.rwMutex.Unlock()
//	chanNotify := make(chan ConfigChangeEvent, 1)
//	config.configChangeListeners = append(config.configChangeListeners, chanNotify)
//	config.addChangeListener(chanNotify)
//	return chanNotify
//}

//func (config *AbstractConfig) GetChangeInterestedKeysNotify(interestedKeys []string) <-chan ConfigChangeEvent {
//	if config.cache == nil {
//		config.cache = newCache()
//	}
//	config.rwMutex.Lock()
//	defer config.rwMutex.Unlock()
//	chanNotify := make(chan ConfigChangeEvent, 1)
//	config.configChangeListeners = append(config.configChangeListeners, chanNotify)
//	config.addChangeListenerInterestedKeys(chanNotify, interestedKeys)
//	return chanNotify
//}

/**
 * Add change listener to this config instance, will only be notified when any of the interested keys is changed in this namespace.
 *
 * @param listener the config change listener
 * @param interestedKeys the keys interested by the listener
 */
func (config *AbstractConfig) AddChangeListenerInterestedKeys(listener ConfigChangeListener, interestedKeys []string) {
	isAdd := false
	for _, value := range config.configChangeListeners {
		if reflect.DeepEqual(value, listener) {
			isAdd = true
			break
		}
	}
	if !isAdd {
		config.configChangeListeners = append(config.configChangeListeners, listener)
		if interestedKeys != nil && len(interestedKeys) > 0 {
			//go-sdk升级后，原先代码在当前版本有问题
			//目前未用到该功能， 暂时注释掉,修复待定
			//config.InterestKeyMap[listener] = interestedKeys
		}
	}
}

func (config *AbstractConfig) AddChangeListenerFuncInterestedKeys(listenerFunc OnChangeFunc, interestedKeys []string) {
	config.AddChangeListenerInterestedKeys(listenerFunc, interestedKeys)
}

/**
 * Remove the change listener
 *
 * @param listener the specific config change listener to remove
 * @return true if the specific config change listener is found and removed
 */
func (config *AbstractConfig) RemoveChangeListener(listener ConfigChangeListener) bool {
	index := -1
	config.rwMutex.Lock()
	defer config.rwMutex.Unlock()
	for key, value := range config.configChangeListeners {
		if reflect.DeepEqual(value, listener) {
			index = key
			break
		}
	}
	if index != -1 {
		config.configChangeListeners = append(config.configChangeListeners[:index], config.configChangeListeners[index+1:]...)
		return true
	}
	return false
}

func (config *AbstractConfig) getValueFromCache(key string) []byte {
	result, _ := config.cache.Get([]byte(key))
	if result != nil {
		return result
	}
	return config.getValueFromPropertiesAndSaveCache(key)

}

func (config *AbstractConfig) getValueFromPropertiesAndSaveCache(key string) []byte {
	currentVersion := config.ConfigVersion
	result := config.GetProperty(key, "")
	if result != nil {
		config.rwMutex.Lock()
		if config.ConfigVersion == currentVersion {
			config.cache.Set([]byte(key), result, config.configUtil.ConfigCacheExpireTime)
		}
		config.rwMutex.Unlock()
		return result
	}
	return nil
}

func (config *AbstractConfig) fireConfigChange(changeEvent ConfigChangeEvent) {
	for _, listener := range config.configChangeListeners {
		if !config.isConfigChangeListenerInterested(listener, changeEvent) {
			continue
		}
		go func(l ConfigChangeListener) {
			l.OnChange(changeEvent)
		}(listener)

	}

}

func (config *AbstractConfig) isConfigChangeListenerInterested(listener ConfigChangeListener, changeEvent ConfigChangeEvent) bool {
	interestedKeys := config.InterestKeyMap[listener]
	if interestedKeys == nil || len(interestedKeys) == 0 {
		return true
	}
	for _, interestedKey := range interestedKeys {
		if changeEvent.IsChanged(interestedKey) {
			return true
		}
	}

	return false
}

func (config *AbstractConfig) calcPropertyChanges(namespace string, previous Properties,
	current Properties) []ConfigChange {
	previousKeys := previous.StringPropertyNames()
	currentKeys := current.StringPropertyNames()

	commonKeys := set.Intersection(previousKeys, currentKeys)
	newKeys := set.Difference(currentKeys, commonKeys)
	removeKeys := set.Difference(previousKeys, commonKeys)

	changes := make([]ConfigChange, 10)

	for _, newKey := range newKeys {
		changeConfig := ConfigChange{
			Namespace:    namespace,
			PropertyName: newKey,
			OldValue:     "",
			NewValue:     current.getProperty(newKey),
			ChangeType:   ADDED,
		}
		changes = append(changes, changeConfig)
	}

	for _, removeKey := range removeKeys {
		changeConfig := ConfigChange{
			Namespace:    namespace,
			PropertyName: removeKey,
			OldValue:     previous.getProperty(removeKey),
			NewValue:     "",
			ChangeType:   DELETED,
		}
		changes = append(changes, changeConfig)
	}

	for _, commonKey := range commonKeys {
		preValue := previous.getProperty(commonKey)
		curvalue := current.getProperty(commonKey)

		if preValue == curvalue {
			continue
		}

		changeConfig := ConfigChange{
			Namespace:    namespace,
			PropertyName: commonKey,
			NewValue:     curvalue,
			OldValue:     preValue,
			ChangeType:   MODIFIED,
		}
		changes = append(changes, changeConfig)
	}
	return changes
}
