package http_client

type Request struct {
	Method string
	URL    string
	Header map[string]string
}

type Response struct {
	StatusCode int
	Body       []byte
}

func NewRequest(method, url string, header map[string]string) *Request {
	return &Request{
		Method: method,
		URL:    url,
		Header: header,
	}
}
