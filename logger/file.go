package logger

import "os"

func IsExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func IsWritable(filePath string) bool {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return false
	}
	defer file.Close()
	return true
}
