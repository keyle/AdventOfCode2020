package main

import (
	"io/ioutil"
	"log"
)

func getFilesContents(filename string) string {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(contents)
}
