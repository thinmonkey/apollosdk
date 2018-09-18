package core

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
