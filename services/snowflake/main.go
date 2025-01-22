package main

import (
	"log"

	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/services/snowflake/handler"
	"github.com/vhall1/shorturl/services/snowflake/service"
)

func main() {
	httpServer, err := bootstrap.NewHttpServer()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: dinamically generate machine IDs
	s, err := service.NewSnowflake(int64(0))
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewSnowflakeHttpHandler(s)
	h.RegisterRoutes(httpServer.Mux)

	if err := httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
