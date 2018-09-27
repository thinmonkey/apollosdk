package core

import (
	"testing"
	"github.com/zhhao226/apollosdk/util"
)

func TestNewRemoteConfigRepository(t *testing.T) {
	configUtil := util.NewConfigUtil("../config.properties", "", "", "", "")
	//success
	NewRemoteConfigRepository("application", configUtil)
	//fail
	//NewRemoteConfigRepository("application_a", configUtil)
}

func TestRemoteConfigRepository_GetConfig(t *testing.T) {
	configUtil := util.NewConfigUtil("../config.properties", "", "", "", "")

	remoteRepository := NewRemoteConfigRepository("application", configUtil)
	t.Log(*remoteRepository.GetConfig())
}
