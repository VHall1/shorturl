package main

import (
	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/service.shortener/domain"
	"github.com/vhall1/shorturl/service.shortener/handler"
	"github.com/vhall1/shorturl/service.shortener/store"
)

func main() {
	addr := ":8080"
	db, err := bootstrap.NewMysqlConn("root:pw@/shorturl")
	if err != nil {
		panic(err)
	}

	httpServer := bootstrap.NewHttpServer(addr)

	// initialise all services
	urlStore := store.NewUrlStore(db)
	urlService := domain.NewUrlService(urlStore)

	// register routes
	handler.RegisterRoutes(httpServer.Mux, urlService)

	// listen and serve
	httpServer.Start()
}
