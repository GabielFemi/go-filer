package go_filer

import (
	"fmt"
	"io/ioutil"
)

func WriteFile(message string, filePath string) bool{
	fileExists := FileExists(filePath)
	if fileExists {
		fmt.Println("The file already exists!")
		return false
	}else {
		messageToByte := []byte (message)
		_ = ioutil.WriteFile(filePath, messageToByte, 0644)

	}
	return true
}
