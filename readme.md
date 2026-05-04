
site project.

api-proxy lets you use a simple HTTP request to get target content. 

Set listen addess and ports in `config.toml`:
```toml
[server]
address="127.0.0.1"
port=8080
```

add access token which is allowed:
```toml
[auth]
allowAccessTokens = [
    "first-secret-token",
    "another-secret-token",
]
```

then start the proxy server:
```shell
./api-proxy-go
```

example using curl to make a request:
```shell
curl -H 'Authorization: Bearer first-secret-token' \
    http://127.0.0.1:8080 \
    -d '{
        "url":"https://example.org/"
        }'
```

the proxy server will response the same HTTP status code and response from target web site:
```
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080
> POST / HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/8.5.0
> Accept: */*
> Authorization: Bearer 123456
> Content-Length: 31
> Content-Type: application/x-www-form-urlencoded
> 
< HTTP/1.1 200 OK
< Date: Mon, 04 May 2026 09:06:11 GMT
< Content-Length: 528
< Content-Type: text/html; charset=utf-8
< 
<!doctype html><html lang="en"><head><title>Example Domain</title><meta name="viewport" content="width=device-width, initial-scale=1"><style>body{background:#eee;width:60vw;margin:15vh auto;font-family:system-ui,sans-serif}h1{font-size:1.5em}div{opacity:0.8}a:link,a:visited{color:#348}</style></head><body><div><h1>Example Domain</h1><p>This domain is for use in documentation examples without needing permission. Avoid use in operations.</p><p><a href="https://iana.org/domains/example">Learn more</a></p></div></body></html>
* Connection #0 to host 127.0.0.1 left intact
```

otherwise, HTTP 500 with an error message will return:
```
< HTTP/1.1 500 Internal Server Error
< Date: Mon, 04 May 2026 09:07:22 GMT
< Content-Length: 20
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 127.0.0.1 left intact
invalid access token
```
