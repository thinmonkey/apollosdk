package core

import "fmt"

const(
	ADDED ConfigChangeType=iota
	MODIFIED
	DELETED
)

//config change type
type ConfigChangeType int

type ConfigChange struct {
	Namespace    string `json:"namespace"`
	PropertyName string `json:"propertyName"`
	OldValue     string `json:"oldValue"`
	NewValue     string `json:"newValue"`
	ChangeType   ConfigChangeType    `json:"changeType"`
}

func (configChange ConfigChange) String() string {
	return fmt.Sprintf(`{"namespace":%s,"propertyName":%s,"oldValue":%s,"newVaule":%s,"changeType":%d}`,
		configChange.Namespace, configChange.PropertyName, configChange.OldValue, configChange.NewValue, configChange.ChangeType)
}
