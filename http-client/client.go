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
