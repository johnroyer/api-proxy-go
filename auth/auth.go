package auth

import "api-proxy-go/config"

type Tokens struct {
	token string
}

func readFromConfig(config config.Config) error {
	size := len(config.Auth.Tokens)

	tokens := make([]string, size)
	for i := 0; i < size; i++ {
		tokens[i] = config.Auth.Tokens[i]
	}
	return nil
}
