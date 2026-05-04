package http_server

import "net/http"

func getBearerToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if token == "" {
		return ""
	}
	return token[7:]
}
