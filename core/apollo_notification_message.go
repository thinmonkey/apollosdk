package core

type ApolloNotificationMessages struct {
	Details map[string]int64 `json:"details,omitempty"`
}

func (apolloNotificationMessages *ApolloNotificationMessages) Put(key string, notificationId int64) {
	apolloNotificationMessages.Details[key] = notificationId
}

func (apolloNotificationMessages *ApolloNotificationMessages) Get(key string) int64 {
	return apolloNotificationMessages.Details[key]
}

func (apolloNotificationMessages *ApolloNotificationMessages) Has(key string) bool {
	_, ok := apolloNotificationMessages.Details[key]
	return ok
}

func (apolloNotificationMessages *ApolloNotificationMessages) MergeFrom(messages ApolloNotificationMessages) {
	for key, value := range messages.Details {
		if _, ok := apolloNotificationMessages.Details[key]; ok && apolloNotificationMessages.Details[key] >= value {
			continue
		}
		apolloNotificationMessages.Details[key] = value
	}
}
