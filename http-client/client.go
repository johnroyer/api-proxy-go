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

func SendRequest(req ProxyRequest) (*ProxyResponse, error) {
	httpReq, err := http.NewRequest(req.Method, req.Url, nil)
	if err != nil {
		return &ProxyResponse{StatusCode: 500, Body: []byte("Error creating request")}, err
	}

	client := &http.Client{}
	response, err := client.Do(httpReq)
	if err != nil {
		return &ProxyResponse{StatusCode: 500, Body: []byte("Error sending request")}, err
	}

	statusCode := response.StatusCode
	htmlBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return &ProxyResponse{StatusCode: 500, Body: []byte("Error reading response body")}, err
	}

	return &ProxyResponse{StatusCode: statusCode, Body: htmlBytes}, nil
}
