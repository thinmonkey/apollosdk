package core

import (
	"testing"
)

func TestNewRemoteConfigRepository(t *testing.T) {
	configUtil := NewConfigWithConfigFile("../config.json")
	//success
	NewRemoteConfigRepository("application", configUtil)
	//fail
	//NewRemoteConfigRepository("application_a", configUtil)
}

func TestRemoteConfigRepository_GetConfig(t *testing.T) {
	configUtil := NewConfigWithConfigFile("../config.json")

	remoteRepository := NewRemoteConfigRepository("application", configUtil)
	t.Log(*remoteRepository.GetConfig())
}
