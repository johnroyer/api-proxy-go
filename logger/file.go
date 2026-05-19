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

func IsAvailable(filePath string) bool {
	return IsExists(filePath) && IsWritable(filePath)
}

func CreateLogChannel(bufferSize int) chan LogData {
	if bufferSize <= 10 {
		return make(chan LogData, 10)
	}
	return make(chan LogData, bufferSize)
}

// 從 channel 中讀取 logData 並寫入指定的檔案
// 此 func 預計放進 goroutine
func Handle(logChannel chan LogData, filePath string) {
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return false
	}
	defer logFile.Close()

	return true
}
