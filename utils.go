package main

import (
	"io/ioutil"
	"log"
	"os"
)

func getFilesContents(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	filebyt, err := ioutil.ReadAll(file)
	return string(filebyt)
}
