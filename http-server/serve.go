package http_server

import (
	"api-proxy-go/auth"
	"api-proxy-go/http-client"
	"encoding/json"
	"net/http"
	"strings"
)

type UserInput struct {
	Url string `json:"url"`
}

type Handler struct {
	AvailableTokens *map[string]bool
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

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 先檢查 bearer token
	accessToken := getBearerToken(r)
	if accessToken == "" {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("access token required"))
		if err != nil {
			print(err.Error())
			return
		}
		return
	}
	if !auth.IsValid(accessToken, h.AvailableTokens) {
		// return HTTP 500
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("invalid access token"))
		if err != nil {
			print(err.Error())
			return
		}
		return
	}

	// 尋找 request 中的 url 參數
	url := getTargetUrl(r)
	if url == "" {
		// return HTTP 500
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("invalid target url"))
		if err != nil {
			print(err.Error())
			return
		}
		return
	}
	// must HTTP
	if strings.HasPrefix(strings.ToLower(url), "http") == false {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("URL must start with http"))
		if err != nil {
			print(err.Error())
		}
		return
	}

	// fetch URL
	request := http_client.ProxyRequest{
		Method: "GET",
		Url:    url,
	}
	response, err := http_client.Fetch(request)
	if err != nil {
		// return HTTP 500
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("invalid target url"))
		if err != nil {
			print(err.Error())
			return
		}
		return
	}

	// 建立 response
	statusCode := response.StatusCode
	html := response.Body
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(html)
	if err != nil {
		print(err.Error())
		return
	}
}
