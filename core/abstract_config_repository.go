package core

import "sync"

const (
	REMOTE ConfigSourceType = iota
	LOCAL
)

type ConfigSourceType int

type AbstractConfigRepository struct {
	Listeners []*RepositoryChangeListener
	rwMutex   sync.RWMutex
}

func (abstractConfigRepository *AbstractConfigRepository) AddChangeListener(listener *RepositoryChangeListener) {
	isAdd := false
	for _, value := range abstractConfigRepository.Listeners {
		if value == listener {
			isAdd = true
			break
		}
	}
	if !isAdd {
		abstractConfigRepository.Listeners = append(abstractConfigRepository.Listeners, listener)
	}
}

func (abstractConfigRepository *AbstractConfigRepository) RemoveChangeListener(listener *RepositoryChangeListener) {
	index := -1
	abstractConfigRepository.rwMutex.Lock()
	defer abstractConfigRepository.rwMutex.Unlock()
	for key, value := range abstractConfigRepository.Listeners {
		if value == listener {
			index = key
			break
		}
	}
	if index != -1 {
		abstractConfigRepository.Listeners = append(abstractConfigRepository.Listeners[:index], abstractConfigRepository.Listeners[index+1:]...)
	}
}

func (abstractConfigRepository *AbstractConfigRepository) FireRepositoryChange(namespace string, newProperties *Properties) {
	for _, value := range abstractConfigRepository.Listeners {
		(*value).OnRepositoryChange(namespace, newProperties)
	}
}
