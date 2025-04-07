package main

import (
	"log"

	"shorturl/internal"
)

func main() {
	server := internal.Server{}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
