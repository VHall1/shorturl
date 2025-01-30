package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vhall1/shorturl/service.shortener/handler"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()

	handler.RegisterRoutes(mux)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("🚀 Server listening on %v", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-done
	log.Printf("💀 Shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}

	log.Println("👋 Done")
}
