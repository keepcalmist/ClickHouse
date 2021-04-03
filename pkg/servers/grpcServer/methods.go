package grpcServer

import (
	"context"
	pb "github.com/keepcalmist/grpcFibonacci/pkg/grpc"
	"github.com/keepcalmist/grpcFibonacci/pkg/usefullFunctions"
	"log"
)

func (s *server) Get(c context.Context, request *pb.Request) (*pb.Response, error) {
	log.Println(request)
	response := &pb.Response{}
	if !usefullFunctions.Validate(request.X, request.Y) {
		response.Error = "Write valid digits "
		return response, nil
	}
	resp, err := s.fibo.CalculateDigits(c, request.X, request.Y)
	if err != nil {
		response.Error = err.Error()
	}
	response.Arr = resp

	return response, nil
}
