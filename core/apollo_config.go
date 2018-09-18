package core

import "fmt"

type ApolloConfig struct {
	AppId string `json:"appId"`

	Cluster string `json:"cluster"`

	NamespaceName string `json:"namespaceName"`

	Configurations map[string]string `json:"configurations"`

	ReleaseKey string `json:"releaseKey"`
}

func (apolloConfig *ApolloConfig) String() string {
	return fmt.Sprintf(`ApolloConfig{appId=%s,cluster=%s,namespaceName=%s,configurations=%s,releaseKey=%s}`,
		apolloConfig.AppId, apolloConfig.Cluster, apolloConfig.NamespaceName, apolloConfig.Configurations, apolloConfig.ReleaseKey)
}
