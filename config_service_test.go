package apollosdk

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	config := GetAppConfig()

	chanEvent := (*config).GetChangeKeyNotify()

	configEvent := <-chanEvent

	t.Log(configEvent)
	configNew := GetConfig("app.tc.mat.disable")
	t.Log((*configNew).GetStringProperty("mats", ""))
}

func TestGetAppConfig(t *testing.T) {

}
