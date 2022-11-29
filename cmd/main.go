package main

import (
	"errors"
	"log"
)

func main() {
	err := runWebServer()
	if err != nil {
		log.Fatal("Shutting down, error: ", err)
	}
}

func runWebServer() error {
	return errors.New("fail run web server")
}
