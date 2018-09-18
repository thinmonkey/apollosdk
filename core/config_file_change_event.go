package core

import "fmt"

type ConfigFileChangeEvent struct {
	Namespace  string `json:"namespace"`
	OldValue   string `json:"oldValue"`
	NewValue   string `json:"newValue"`
	ChangeType int    `json:"changeType"`
}

func (configFileChange ConfigFileChangeEvent) String() string {
	return fmt.Sprintf(`{"namespace":%s,"oldValue":%s,"newVaule":%s,"changeType":%d}`,
		configFileChange.Namespace, configFileChange.OldValue, configFileChange.NewValue, configFileChange.ChangeType)
}
