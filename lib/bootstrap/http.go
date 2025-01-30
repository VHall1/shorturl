package bootstrap

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vhall1/shorturl/lib/middleware"
)

type HttpServer struct {
	Mux    *http.ServeMux
	Server *http.Server
}

func NewHttpServer(addr string) *HttpServer {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr: addr,
		// middleware applied to all services
		Handler: middleware.ApplyMiddleware(mux, middleware.Logger),
	}

	return &HttpServer{
		Mux:    mux,
		Server: server,
	}
}

func (s *HttpServer) Start() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("🚀 Server listening on %v", s.Server.Addr)
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-done
	log.Printf("💀 Shutting down")

	if err := s.Server.Shutdown(context.Background()); err != nil {
		panic(err)
	}

	log.Println("👋 Done")
}
