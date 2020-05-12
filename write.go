package go_filer

import (
	"io/ioutil"
)

func WriteFile(message string, filePath string) bool{
	messageToByte := []byte (message)
	_ = ioutil.WriteFile(filePath, messageToByte, 0644)
	return true
}
