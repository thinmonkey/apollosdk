package core

import "fmt"

type ApolloConfigNotification struct {
	NamespaceName  string                     `json:"namespaceName"`
	NotificationId int64                      `json:"notificationId"`
	Messages       ApolloNotificationMessages `json:"messages,omitempty"`
}

func (apolloConfigNotification *ApolloConfigNotification) AddMessage(key string, notificationId int64) {
	apolloConfigNotification.Messages.Put(key, notificationId)
}

func (apolloConfigNotification *ApolloConfigNotification) String() string {
	return fmt.Sprintf(`ApolloConfigNotification{namespaceName=%s,notificationId=%s}`,
		apolloConfigNotification.NamespaceName, apolloConfigNotification.NotificationId)
}
