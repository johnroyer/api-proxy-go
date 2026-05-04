package http_client

import "testing"

func TestMakeHttpRequest(t *testing.T) {
	var request ProxyRequest
	request = ProxyRequest{
		Method: "GET",
		Url:    "https://zeroplex.tw/ip",
	}

	response, err := SendRequest(request)
	if err != nil {
		t.Error("send request failed: " + err.Error())
	}

	if response.StatusCode != 200 {
		t.Error("unexpected status code: " + string(rune(response.StatusCode)))
	}

	if response.Body == nil {
		t.Error("response body should not be empty")
	}
}
