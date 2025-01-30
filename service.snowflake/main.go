package main

import (
	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/service.snowflake/domain"
	"github.com/vhall1/shorturl/service.snowflake/handler"
)

func main() {
	httpServer := bootstrap.NewHttpServer(":8080")

	// initialise all services
	snowflakeService, err := domain.NewSnowflakeService(0)
	if err != nil {
		panic(err)
	}

	// register routes
	handler.RegisterRoutes(httpServer.Mux, snowflakeService)

	// listen and serve
	httpServer.Start()
}
