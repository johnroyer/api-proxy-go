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

func SendRequest(req ProxyRequest) ProxyResponse {
	httpReq, err := http.NewRequest(req.Method, req.Url, nil)
	if err != nil {
		return ProxyResponse{StatusCode: 500, Body: []byte("Error creating request")}
	}

	client := &http.Client{}
	response, err := client.Do(httpReq)
	if err != nil {
		return ProxyResponse{StatusCode: 500, Body: []byte("Error sending request")}
	}

	statusCode := response.StatusCode
	htmlBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return ProxyResponse{StatusCode: 500, Body: []byte("Error reading response body")}
	}

	return ProxyResponse{StatusCode: statusCode, Body: htmlBytes}
}
