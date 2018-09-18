package http

import "time"

type HttpRequest struct {
	Url            string        `json:"url"`
	ConnectTimeout time.Duration `json:"connectTimeout"`
}

type HttpResponse struct {
	StatusCode  int    `json:"statusCode"`
	ReponseBody []byte `json:"ReponseBody"`
}
