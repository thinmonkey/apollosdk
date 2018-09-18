package core

import (
	"fmt"
)

type ServiceDto struct {
	AppName     string `json:"appName"`
	InstanceId  string `json:"instanceId"`
	HomePageUrl string `json:"homePageUrl"`
}

func (serviceDto *ServiceDto) String() string {
	return fmt.Sprintf(`ServiceDto{appName=%s,instanceId=%s,homePageUrl=%s}`, serviceDto.AppName, serviceDto.InstanceId, serviceDto.HomePageUrl)
}
