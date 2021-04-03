package grpcServer

import (
	"fmt"
	"github.com/keepcalmist/grpcFibonacci/pkg/fibonacci"
	pb "github.com/keepcalmist/grpcFibonacci/pkg/grpc"
	"github.com/keepcalmist/grpcFibonacci/pkg/servers"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

type server struct {
	pb.FibonacciServer
	fibo fibonacci.FiboService
}

func New(fibo fibonacci.FiboService) servers.Server {
	return &server{
		fibo: fibo,
	}
}

func (s *server) Run(conf *viper.Viper, quit chan bool) {
	listener, err := net.Listen("tcp", fmt.Sprint(":", conf.GetString("GRPC_SERVER_PORT")))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFibonacciServer(grpcServer, s)
	log.Println("grpc server starting on port: ", conf.GetString("GRPC_SERVER_PORT"))
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Println(err)
		}
	}()
	<-quit
	grpcServer.GracefulStop()
}
