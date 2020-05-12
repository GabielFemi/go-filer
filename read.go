package go_filer

import (
	"io/ioutil"
)
// File path should be absolute e.g. C:\Users\Real Stuff\Desktop\test.txt
func ReadFile(fileLocation string) []byte{
	fileExists := FileExists(fileLocation)
	if fileExists {
		logFatal()
	}
	data, _ := ioutil.ReadFile(fileLocation)
	return data
}