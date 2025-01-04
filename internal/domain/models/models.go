package models

import "net/http"

type Response struct {
	Data       any   `json:"data,omitempty"`
	StatusCode int   `json:"status_code"`
	Error      error `json:"error,omitempty"`
}

type EndPoint struct {
	Target  string
	Handler http.HandlerFunc
}

func NewEndPoint(target string, handler func(http.ResponseWriter, *http.Request)) EndPoint {
	return EndPoint{
		Target:  target,
		Handler: handler,
	}
}
