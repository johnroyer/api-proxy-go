package http_server

import (
	"encoding/json"
	"net/http"
)

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

func getTargetUrl(httpRequest *http.Request) string {
	if httpRequest.Method != "POST" {
		return ""
	}

	var userInput UserInput
	err := json.NewDecoder(httpRequest.Body).Decode(&userInput)
	if err != nil {
		return ""
	}
	return userInput.Url
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// 先檢查 bearer token

	// 尋找 request 中的 url 參數

	// fetch URL

	// 建立 response
}
