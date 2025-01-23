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
	server *http.Server
}

type httpServerConf struct {
	HttpAddr string `envconfig:"optional"`
}

// Creates a new HTTP server struct, it also exposes a new mux
// that can be used for routing. The server address is read from the HTTP_ADDR env variable
func NewHttpServer() (*HttpServer, error) {
	conf := &httpServerConf{}
	if err := config.Load(conf); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	addr := conf.HttpAddr
	// bind to port 80 by default
	if addr == "" {
		log.Printf("[HTTP] no address set. Using fallback :80\n")
		addr = ":80"
	}

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &HttpServer{
		Mux:    mux,
		server: server,
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

		log.Println("[HTTP] Shutdown signal received. Gracefully shutting down server")

		ctx, release := context.WithTimeout(context.Background(), time.Second*10)
		defer release()

		if err := s.server.Shutdown(ctx); err != nil {
			ch <- fmt.Errorf("got an error while trying to gracefully shutdown http server: %v", err)
		}
	}()

	go func() {
		log.Printf("[HTTP] Server listening on %s\n", s.server.Addr)
		ch <- s.server.ListenAndServe()
	}()

	err := <-ch
	// `http.ErrServerClosed` is always returned when the Shutdown function is called
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
