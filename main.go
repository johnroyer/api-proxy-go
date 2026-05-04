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
	server := http.ListenAndServe(
		listen,
		&handler,
	)
	println("server start on " + listen)
	log.Fatal(server)
}
