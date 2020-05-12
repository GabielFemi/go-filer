package go_filer

import (
	"os"
)

func FileExists(filePath string) bool{
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}


