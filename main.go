package main

import http_client "api-proxy-go/http-client"

func main() {
	response, err := http_client.SendRequest(
		http_client.ProxyRequest{
			Method: "GET",
			Url:    "https://zeroplex.tw/ip",
		})
	if err != nil {
		println("error occures: " + err.Error())
	}

	print(string(response.Body))
}
