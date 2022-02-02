package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const check_dir string = "/tmp/mfchecks/"

func main() {

	newpath := filepath.Join(check_dir)
	err := os.MkdirAll(newpath, os.ModePerm)

	// https://stackoverflow.com/questions/14668850/list-directory-in-go
	files, err := ioutil.ReadDir(check_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

}
