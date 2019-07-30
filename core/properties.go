package core

import "fmt"

type Properties map[string]string

func (properties *Properties) StringPropertyNames() []string {
	keys := make([]string, 0)
	for key := range *properties {
		keys = append(keys, key)
	}
	return keys
}

func (properties *Properties) getProperty(key string) string {
	return (*properties)[key]
}

func (properties *Properties) PutAll(newProperties Properties) {
	for key, value := range newProperties {
		(*properties)[key] = value
	}
}

func (properties *Properties) Contain(key string) bool {
	if _, ok := (*properties)[key]; ok {
		return ok
	}
	return false
}

func (properties *Properties) ToString() string {
	var properString string
	for key, value := range *properties {
		properString = fmt.Sprintf("%s%s=%s\n", properString, key, value)
	}
	return properString[:len(properString)-1]
}
