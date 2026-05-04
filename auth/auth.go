package auth

import "api-proxy-go/config"

type Tokens struct {
	token string
}

func readFromConfig(config config.Config) (map[string]bool, error) {
	size := len(config.Auth.Tokens)

	tokens := make(map[string]bool, size)
	for _, token := range config.Auth.Tokens {
		tokens[token] = true
	}
	return tokens, nil
}
