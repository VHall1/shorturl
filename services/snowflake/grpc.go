package main

import (
	"github.com/vhall1/shorturl/lib/bootstrap"
	"github.com/vhall1/shorturl/services/snowflake/types"
)

type GrpcServer struct {
	snowflakeService types.Snowflake
}

func NewGrpcServer(snowflakeService types.Snowflake) *GrpcServer {
	return &GrpcServer{
		snowflakeService: snowflakeService,
	}
}

func (s *GrpcServer) Start() error {
	grpcServer, err := bootstrap.NewGrpcServer()
	if err != nil {
		return err
	}

	// h := handler.NewSnowflakeHttpHandler(s.snowflakeService)
	// h.RegisterRoutes(httpServer.Mux)

	return grpcServer.Start()
}
