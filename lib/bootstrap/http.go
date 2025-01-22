package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vhall1/shorturl/lib/config"
)

type HttpServer struct {
	Mux    *http.ServeMux
	Server *http.Server
}

type httpServerConf struct {
	HttpAddr string
}

// Creates a new HTTP server struct. The server address is read
// from the HTTP_ADDR env variable
func NewHttpServer() (*HttpServer, error) {
	conf := &httpServerConf{}
	if err := config.Load(conf); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    conf.HttpAddr,
		Handler: mux,
	}

	return &HttpServer{
		Mux:    mux,
		Server: server,
	}, nil
}

func (s *HttpServer) Start() error {
	ch := make(chan error)

	// listen for SIGINT and SIGTERM shutdown signals
	// and attempt to gracefully shutdown http server
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs

		fmt.Println("Shutdown signal received. Gracefully shutting down http server")

		ctx, release := context.WithTimeout(context.Background(), time.Second*10)
		defer release()

		if err := s.Server.Shutdown(ctx); err != nil {
			ch <- fmt.Errorf("got an error while trying to gracefully shutdown http server: %v", err)
		}
	}()

	go func() {
		log.Printf("HTTP server listening on %s\n", s.Server.Addr)
		ch <- s.Server.ListenAndServe()
	}()

	err := <-ch
	// `http.ErrServerClosed` is always returned when the Shutdown function is called
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
