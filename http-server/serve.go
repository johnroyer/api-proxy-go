package http_server

import "net/http"

type UserInput struct {
	Url string `json:"url"`
}

func getBearerToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if token == "" {
		return ""
	}
	return token[7:]
}
