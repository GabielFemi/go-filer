package go_filer

import (
	"log"
	"os"
)


func logFatal() {
	log.Fatal(os.ErrExist)
}

func logDoesNotExist() {
	log.Fatal(os.ErrNotExist)
}
