package config

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server struct {
		Address string `toml:"address"`
		Port    int    `toml:"port"`
	}

	Proxy struct {
		Timeout int `toml:"timeout"`
	}

	Auth struct {
		Tokens []string `toml:"allowAccessTokens"`
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

	var userConfig Config
	err := toml.Unmarshal(configText, &userConfig)
	if err != nil {
		println(err.Error())
		return nil, errors.New("failed to parse config")
	}
	return &userConfig, nil
}
