package main

//go:generate go run github.com/rakyll/statik -f -src=. -include=main.go

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/rakyll/statik/fs"

	_ "github.com/byo/go-s390x-test/statik" // TODO: Replace with the absolute import path
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("Can't open statik FS: ", err)
	}

	// Access individual files by their paths.
	r, err := statikFS.Open("/main.go")
	if err != nil {
		log.Fatal("Can't open data: ", err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("Cant' read data: ", err)
	}

	fmt.Println(string(contents))
}
