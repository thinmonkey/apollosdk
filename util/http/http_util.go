package http

import (
	"net/http"
	"io/ioutil"
	"errors"
	"github.com/zhhao226/apollosdk/util"
)

func Request(request HttpRequest) (*HttpResponse, error) {
	client := &http.Client{
		Timeout: request.ConnectTimeout,
	}

	var responseBody []byte

	res, err := client.Get(request.Url)

	if res == nil || err != nil {
		util.Logger.Error("Connect Apollo Server Fail,Error:", err)
		return nil, err
	}

	//not modified break
	switch res.StatusCode {
	case http.StatusOK:
		responseBody, err = ioutil.ReadAll(res.Body)
		if err != nil {
			util.Logger.Error("Connect Apollo Server Fail,Error:", err)
			return nil, err
		} else {
			return &HttpResponse{http.StatusOK, responseBody}, nil
		}
	case http.StatusNotModified:
		return &HttpResponse{http.StatusNotModified, nil}, nil
	default:
		util.Logger.Error("Connect Apollo Server Fail,Error:", err)
		if res != nil {
			util.Logger.Error("Connect Apollo Server Fail,StatusCode:", res.StatusCode)
		}
		err = errors.New("Connect Apollo Server Fail!")
		// if error then sleep
		return nil, err
	}
}
