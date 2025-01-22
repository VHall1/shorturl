package main

import (
	"log"

	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/services/shortener/handler"
	"github.com/vhall1/shorturl/services/shortener/service"
	"github.com/vhall1/shorturl/services/shortener/store"
)

func main() {
	httpServer, err := bootstrap.NewHttpServer()
	if err != nil {
		log.Fatal(err)
	}

	psql, err := bootstrap.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewUrlStore(psql)
	svc := service.NewShortenerService(store)
	h := handler.NewShortenerHttpHandler(svc)
	h.RegisterRoutes(httpServer.Mux)

	if err := httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
