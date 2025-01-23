package main

import (
	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/services/snowflake/handler"
	"github.com/vhall1/shorturl/services/snowflake/types"
)

type HttpServer struct {
	snowflakeService types.Snowflake
}

func NewHttpServer(snowflakeService types.Snowflake) *HttpServer {
	return &HttpServer{
		snowflakeService: snowflakeService,
	}
}

func (s *HttpServer) Start() error {
	httpServer, err := bootstrap.NewHttpServer()
	if err != nil {
		return err
	}

	h := handler.NewSnowflakeHttpHandler(s.snowflakeService)
	h.RegisterRoutes(httpServer.Mux)

	return httpServer.Start()
}
