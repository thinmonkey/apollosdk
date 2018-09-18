package apollosdk

import (
	"testing"
	"github.com/zhhao226/apollosdk/core"
	"github.com/zhhao226/apollosdk/util"
)

func TestGetConfig(t *testing.T) {
	config := GetAppConfig()
	configLis := ConfigListener{}

	configListener := core.ConfigChangeListener(&configLis)
	(*config).AddChangeListener(&configListener)
	t.Log((*config).GetStringProperty("test", ""))
	configNew := GetConfig("app.tc.mat.disable")
	t.Log((*configNew).GetStringProperty("mats", ""))
}

type ConfigListener struct {
}

func (configListener *ConfigListener) OnChange(changeEvent core.ConfigChangeEvent) () {
	util.Logger.Info(changeEvent)
}

func TestGetAppConfig(t *testing.T) {

}
