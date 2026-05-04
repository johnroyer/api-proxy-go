package config

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

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

func fileIsExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func LoadConfig(path string) (*Config, error) {
	if !fileIsExist(path) {
		return nil, errors.New("Config file not found")
	}

	configText, _ := os.ReadFile(path)
	if configText == nil {
		return nil, errors.New("Config file is not readable")
	}

	var config Config
	err := toml.Unmarshal(configText, &config)
	if nil == err {
		return nil, errors.New("Failed to parse config file")
	} else {
		return &config, nil
	}
}
