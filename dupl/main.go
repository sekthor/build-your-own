package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var test filepath.WalkFunc

func index(path string, info fs.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	fmt.Println(path)

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}
	filepath.Walk(os.Args[1], index)
}
