package main

import (
	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/service.shortener/domain"
	"github.com/vhall1/shorturl/service.shortener/handler"
	"github.com/vhall1/shorturl/service.shortener/store"
)

func main() {
	db, err := bootstrap.NewMysqlConn("root:pw@/shorturl")
	if err != nil {
		panic(err)
	}

	httpServer := bootstrap.NewHttpServer(":8080")

	// initialise all services
	urlStore := store.NewUrlStore(db)
	urlService := domain.NewUrlService(urlStore)

	// register routes
	handler.RegisterRoutes(httpServer.Mux, urlService)

	// listen and serve
	httpServer.Start()
}
