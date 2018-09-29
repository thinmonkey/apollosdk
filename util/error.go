package util

import "fmt"

type ApolloConfigError struct {
	Message string
}

func (e ApolloConfigError) Error() string {
	return e.Message
}

type ApolloConfigStatusCodeError struct {
	Message string
	StatusCode int
}

func (e ApolloConfigStatusCodeError) Error() string  {
	return fmt.Sprintf("[status code: %d] %s",e.StatusCode,e.Message)
}

func (e ApolloConfigStatusCodeError) GetStatusCode() int  {
	return e.StatusCode
}
