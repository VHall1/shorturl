package main

import (
	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/service.shortener/common"
	"github.com/vhall1/shorturl/service.shortener/domain"
	"github.com/vhall1/shorturl/service.shortener/handler"
	"github.com/vhall1/shorturl/service.shortener/store"
)

func main() {
	// TODO: pull these from a config somewhere else
	addr := ":8080"
	db, err := common.NewMysqlConn("root:pw@/shorturl")
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
