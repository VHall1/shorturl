package main

import (
	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/services/shortener/handler"
	"github.com/vhall1/shorturl/services/shortener/service"
	"github.com/vhall1/shorturl/services/shortener/store"
)

func main() {
	http, err := bootstrap.NewHttpServer()
	if err != nil {
		panic(err)
	}

	psql, err := bootstrap.NewPostgres()
	if err != nil {
		panic(err)
	}

	store := store.NewUrlStore(psql)
	svc := service.NewShortenerService(store)
	h := handler.NewShortenerHttpHandler(svc)
	h.RegisterRoutes(http.Mux)

	if err := http.Start(); err != nil {
		panic(err)
	}
}
