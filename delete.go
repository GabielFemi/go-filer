package go_filer

import (
	"log"
	"os"
)

func DeleteFile(fileName string) bool{
	fileExists := FileExists(fileName)
	if !fileExists {
		logDoesNotExist()
	}else {
		err := os.Remove(fileName)

		if err != nil {
			log.Panic("Some thing extreme occurred!")
		}

	}
	return true
}