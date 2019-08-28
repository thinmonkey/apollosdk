package core

import (
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	configUtil := NewConfigWithConfigFile("../config.json")

	remoteRepository := NewRemoteConfigRepository("application", configUtil)

	configRepo := ConfigRepository(remoteRepository)
	NewDefaultConfig("application", configRepo, configUtil)
}

func TestDefaultConfig_GetPropertyNames(t *testing.T) {
	configUtil := NewConfigWithConfigFile("../config.json")

	remoteRepository := NewRemoteConfigRepository("application", configUtil)

	configRepo := ConfigRepository(remoteRepository)
	config := NewDefaultConfig("application", configRepo, configUtil)
	t.Log(config.GetPropertyNames())
}

func TestDefaultConfig_OnRepositoryChange(t *testing.T) {
	configUtil := NewConfigWithConfigFile("../config.json")

	remoteRepository := NewRemoteConfigRepository("application", configUtil)

	configRepo := ConfigRepository(remoteRepository)
	config := NewDefaultConfig("application", configRepo, configUtil)

	newProperties := Properties{}
	newProperties["aaa"] = "content"
	config.OnRepositoryChange("application", &newProperties)
}
