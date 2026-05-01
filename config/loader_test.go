package config

import "testing"

func TestOnFileNotExists(t *testing.T) {
	exists := fileIsExist("/path/to/file/that/does/not/exist.toml")

	if exists {
		t.Error("File expected not to exist, but it does")
	}
}
