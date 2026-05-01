package config

import "os"

func fileIsExist() bool {
	if _, err := os.Stat("./config.toml"); err == nil {
		return true
	} else {
		return false
	}
}
