package http_client

type Request struct {
	Method string
	Url    string
	Header map[string]string
}

type Response struct {
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
