package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var sizeIndex map[int64][]string = make(map[int64][]string)

func index(path string, info fs.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	size := info.Size()
	sizeIndex[size] = append(sizeIndex[size], path)
	fmt.Println(path)

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}
	filepath.Walk(os.Args[1], index)

	for _, v := range sizeIndex {
		if len(v) > 1 {
			fmt.Println("potential duplicates:", v)
		}
	}
}
