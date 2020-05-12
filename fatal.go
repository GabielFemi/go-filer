package go_filer

import (
	"log"
	"os"
)


func logFatal() {
	log.Fatal(os.ErrExist)
}


