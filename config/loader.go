package config

import "os"

type Config struct {
	Server struct {
		Address string
		Port    int
	}

	Proxy struct {
		Timeout int
	}

	Auth struct {
		Tokens []string
	}
}

func fileIsExist() bool {
	if _, err := os.Stat("./config.toml"); err == nil {
		return true
	} else {
		return false
	}
}
