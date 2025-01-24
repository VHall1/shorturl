package main

import (
	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/services/snowflake/handler"
	"github.com/vhall1/shorturl/services/snowflake/service"
)

func main() {
	// TODO: dinamically generate machine IDs
	svc, err := service.NewSnowflake(int64(0))
	if err != nil {
		panic(err)
	}

	http, err := bootstrap.NewHttpServer()
	if err != nil {
		panic(err)
	}

	h := handler.NewSnowflakeHttpHandler(svc)
	h.RegisterRoutes(http.Mux)

	if err := http.Start(); err != nil {
		panic(err)
	}
}
