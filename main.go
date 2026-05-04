package main

import (
	"api-proxy-go/config"
	http_server "api-proxy-go/http-server"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// read config
	userConfig, err := config.LoadConfig("./config.toml")
	if err != nil || userConfig == nil {
		println("failed to read config")
		return
	}

	availableTokens := map[string]bool{}
	for i := 0; i < len(userConfig.Auth.Tokens); i++ {
		availableTokens[userConfig.Auth.Tokens[i]] = true
	}

	handler := http_server.Handler{
		AvailableTokens: &availableTokens,
	}

	listen := userConfig.Server.Address + ":" + strconv.Itoa(userConfig.Server.Port)
	println("server starting on " + listen)
	err = http.ListenAndServe(listen, &handler)
	if err != nil {
		log.Fatal("server failed: ", err)
	}
}
