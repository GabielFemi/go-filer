package go_filer

import (
	"io/ioutil"
)

func WriteFile(message string, filePath string) bool{
	fileExists := FileExists(filePath)
	if fileExists {
		logFatal()
		return false
	}else {
		messageToByte := []byte (message)
		_ = ioutil.WriteFile(filePath, messageToByte, 0644)

	}
	return true
}
