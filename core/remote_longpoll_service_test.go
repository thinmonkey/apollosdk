package core

import (
	"testing"
	"github.com/thinmonkey/apollosdk/util"
)

func TestNewRemoteConfigLongPollService(t *testing.T) {
	configUtil := util.NewConfigUtil("../config.json", "", "", "", "")
	NewRemoteConfigLongPollService(configUtil)
}

func TestRemoteConfigLongPollService_Submit(t *testing.T) {
	configUtil := util.NewConfigUtil("../config.json", "", "", "", "")
	remoteConfig := *NewRemoteConfigLongPollService(configUtil)

	remoteRepository := NewRemoteConfigRepository("application", configUtil)

	remoteConfig.Submit("application", remoteRepository)
}
