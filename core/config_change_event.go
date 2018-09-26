package core

type ConfigChangeListener interface {
	OnChange(changeEvent ConfigChangeEvent)
}

//定义一个函数类型
type OnChangeFunc func(changeEvent ConfigChangeEvent)

//函数实现接口，这样既可以用方法，也可以用接口，灵活性增加
func (f OnChangeFunc) OnChange(changeEvent ConfigChangeEvent) {
	f(changeEvent)
}

type ConfigChangeEvent struct {
	Namespace string                   `json:"namespace"`
	Changes   map[string]ConfigChange `json:"changes"`
}

func (configChangeEvent ConfigChangeEvent) ChangeKeys() []string {
	keys := make([]string, 0)
	for key := range configChangeEvent.Changes {
		keys = append(keys, key)
	}
	return keys
}

func (configChangeEvent ConfigChangeEvent) GetChanges(key string) ConfigChange {
	configChange, ok := configChangeEvent.Changes[key]
	if ok {
		return configChange
	} else {
		return ConfigChange{}
	}
}

func (configChangeEvent ConfigChangeEvent) IsChanged(key string) bool {
	_, ok := configChangeEvent.Changes[key]
	if ok {
		return true
	} else {
		return false
	}
}
