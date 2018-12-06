package core

import (
	"testing"
)

func TestNewRemoteConfigLongPollService(t *testing.T) {
	configUtil := NewConfigUtil("../config.json", "", "", "", "")
	NewRemoteConfigLongPollService(configUtil)
}

func TestRemoteConfigLongPollService_Submit(t *testing.T) {
	configUtil := NewConfigUtil("../config.json", "", "", "", "")
	remoteConfig := *NewRemoteConfigLongPollService(configUtil)

	remoteRepository := NewRemoteConfigRepository("application", configUtil)

	remoteConfig.Submit("application", remoteRepository)
}
