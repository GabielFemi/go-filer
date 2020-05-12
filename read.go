package go_filer

import (
	"io/ioutil"
	"log"
	"os"
)
// File path should be absolute e.g. C:\Users\Real Stuff\Desktop\test.txt
func ReadFile(fileLocation string) []byte{
	_, err := os.Stat(fileLocation)
	if os.IsNotExist(err) {
		log.Fatal("Folder does not exist")
	}
	data, _ := ioutil.ReadFile(fileLocation)
	return data
}