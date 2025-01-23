package bootstrap

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/vhall1/shorturl/lib/config"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Server   *grpc.Server
	listener *net.Listener
	addr     string
}

type grpcServerConf struct {
	GrpcAddr string `envconfig:"optional"`
}

func NewGrpcServer() (*GrpcServer, error) {
	conf := &grpcServerConf{}
	if err := config.Load(conf); err != nil {
		return nil, err
	}

	addr := conf.GrpcAddr
	if addr == "" {
		log.Printf("[gRPC] no address set. Using fallback :50051\n")
		addr = ":50051"
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()

	return &GrpcServer{
		Server:   s,
		listener: &lis,
		addr:     addr,
	}, nil
}

func (s *GrpcServer) Start() error {
	// listen for SIGINT and SIGTERM shutdown signals
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		log.Println("[gRPC] Shutdown signal received. Gracefully shutting down server")
		s.Server.GracefulStop()
	}()

	ch := make(chan error)
	go func() {
		log.Printf("[gRPC] Server listening on %s\n", s.addr)
		ch <- s.Server.Serve(*s.listener)
	}()

	err := <-ch
	return err
}
