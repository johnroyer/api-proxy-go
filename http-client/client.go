package http_client

import (
	"io"
	"net/http"
)

type ProxyRequest struct {
	Method string
	Url    string
}

type ProxyResponse struct {
	StatusCode int
	Body       []byte
}

func SendRequest(method, url string, header map[string]string) *Request {
	return &Request{
		Method: method,
		Url:    url,
		Header: header,
	}
}
