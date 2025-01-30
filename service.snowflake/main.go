package main

import "github.com/vhall1/shorturl/lib/bootstrap"

func main() {
	addr := ":8080"

	httpServer := bootstrap.NewHttpServer(addr)

	// // initialise all services
	// urlStore := store.NewUrlStore(db)
	// urlService := domain.NewUrlService(urlStore)

	// // register routes
	// handler.RegisterRoutes(httpServer.Mux, urlService)

	// listen and serve
	httpServer.Start()
}
