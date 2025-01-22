package main

import (
	"log"

	"github.com/vhall1/shorturl/lib/bootstrap"
)

func main() {
	httpServer, err := bootstrap.NewHttpServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
