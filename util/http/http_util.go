package http

import (
	"github.com/thinmonkey/apollosdk/util"
	"io/ioutil"
	"net/http"
)

func Request(request HttpRequest) (*HttpResponse, error) {
	client := &http.Client{
		Timeout:request.ConnectTimeout,
	}

	var responseBody []byte

	res, err := client.Get(request.Url)

	if res == nil || err != nil {
		util.DebugPrintf("Connect Apollo Server Fail,Error:%v", err)
		return nil, util.ApolloConfigError{Message: err.Error()}
	}

	responseBody, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		util.DebugPrintf("Read Apollo Server Body Fail,Error:", err)
		return nil, util.ApolloConfigError{Message: err.Error()}
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNotModified {
		return &HttpResponse{res.StatusCode, responseBody}, nil
	}
	err = util.ApolloConfigStatusCodeError{StatusCode: res.StatusCode, Message: string(responseBody)}
	util.DebugPrintf("Apollo Server httpResponse error:", err.Error())
	return nil, err
}
