package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	urlStore := store.NewUrlStore(db)
	urlService := domain.NewUrlService(urlStore)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux, urlService)

	srv := &http.Server{
		Addr: addr,
		// middleware applied to all routes
		Handler: common.ApplyMiddleware(mux, common.Logger),
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
