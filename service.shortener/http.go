package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vhall1/shorturl/service.shortener/common"
)

type httpServer struct {
	mux    *http.ServeMux
	server *http.Server
}

func newHttpServer(addr string) *httpServer {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr: addr,
		// middleware applied to all routes
		Handler: common.ApplyMiddleware(mux, common.Logger),
	}

	return &httpServer{
		mux:    mux,
		server: server,
	}
}

func (s *httpServer) start() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("🚀 Server listening on %v", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-done
	log.Printf("💀 Shutting down")

	if err := s.server.Shutdown(context.Background()); err != nil {
		panic(err)
	}

	log.Println("👋 Done")
}
