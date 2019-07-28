package core

import (
	"reflect"
	"sync"
)

/**

 */
type AbstractConfigFile struct {
	ConfigRepository ConfigRepository
	Namespace string
	Listeners []ConfigFileChangeListener
	Properties *Properties
	SourceType ConfigSourceType

	rwMutex       sync.Mutex
}

func (config *AbstractConfigFile) GetContent() string {
	panic("implement me")
}

func (config *AbstractConfigFile) HasContent() bool {
	panic("implement me")
}

func (config *AbstractConfigFile) GetNamespace() string {
	return config.Namespace
}

func (config *AbstractConfigFile) GetConfigFileFormat() string {
	panic("implement me")
}

func (config *AbstractConfigFile) update(newProperties *Properties) {
	panic("implement me")
}

func (config *AbstractConfigFile) AddChangeListener(listener ConfigFileChangeListener){
	var isAdded  bool
	for _, value := range config.Listeners {
		if reflect.DeepEqual(value, listener) {
			isAdded = true
			break
		}
	}
	if !isAdded {
		config.Listeners = append(config.Listeners,listener)
	}
}

func (config *AbstractConfigFile) RemoveChangeListener(listener ConfigFileChangeListener) bool{
	index := -1
	config.rwMutex.Lock()
	defer config.rwMutex.Unlock()
	for key, value := range config.Listeners {
		if reflect.DeepEqual(value, listener) {
			index = key
			break
		}
	}
	if index != -1 {
		config.Listeners = append(config.Listeners[:index], config.Listeners[index+1:]...)
		return true
	}
	return false
}

func (config *AbstractConfigFile) OnRepositoryChange(namespace string, newProperties *Properties){
	if newProperties == config.Properties {
		return
	}
	var newConfigProperties Properties
	newConfigProperties.PutAll(*newProperties)

	oldValue := config.GetContent()

	config.update(newProperties)
	config.SourceType = config.ConfigRepository.getSourceType()

	 newValue := config.GetContent()

	 changeType := MODIFIED

	if oldValue == "" {
		changeType = ADDED
	} else if newValue == "" {
		changeType = DELETED
	}

	config.fireConfigChange(ConfigFileChangeEvent{
		Namespace:  config.Namespace,
		OldValue:   oldValue,
		NewValue:   newValue,
		ChangeType: int(changeType),
	})
}

func (config *AbstractConfigFile) fireConfigChange(changeEvent ConfigFileChangeEvent) ()  {
	for _,listener := range config.Listeners{
		go func() {
			listener.OnChange(changeEvent)
		}()
	}
}