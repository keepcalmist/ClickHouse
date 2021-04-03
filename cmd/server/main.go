package main

import (
	"context"
	"fmt"
	"github.com/keepcalmist/grpcFibonacci/pkg/config"
	pb "github.com/keepcalmist/grpcFibonacci/pkg/grpc"
	"github.com/keepcalmist/grpcFibonacci/pkg/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {
	pb.FibonacciServer
	rConn redis.Repo
}

const port = "5300"

func main() {
	conf := config.New()

	srv := &server{
		rConn: redis.New("localhost:" + conf.GetString("REDIS_PORT")),
	}

	listener, err := net.Listen("tcp", fmt.Sprint(":", port))

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterFibonacciServer(grpcServer, srv)
	log.Println("Server starting on port: ", port)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Println(err)
		}
	}()
	chanStop := make(chan os.Signal)
	signal.Notify(chanStop, os.Interrupt)
	<-chanStop
	grpcServer.GracefulStop()
}

func (s *server) Get(c context.Context, request *pb.Request) (*pb.Response, error) {
	log.Println(request)
	response := &pb.Response{}
	if !validate(request.X, request.Y) {
		response.Error = "Write valid digits "
		return response, nil
	}
	resp, err := s.calculateDigits(c, request.X, request.Y)
	if err != nil {
		response.Error = err.Error()
	}
	response.Arr = resp

	return response, nil
}

func (s *server) calculateDigits(ctx context.Context, from, to int32) ([]int32, error) {
	maxKeyInStorage, err := s.rConn.GetMax(ctx)
	if err != nil {
		return nil, err
	}
	if maxKeyInStorage > int(from) && maxKeyInStorage > int(to) {
		fiboSlice, err := s.rConn.GetSliceDigits(ctx, int(from), int(to))
		if err != nil {
			return nil, err
		}
		reSlice := make([]int32, 0, len(fiboSlice))
		for _, value := range fiboSlice {
			reSlice = append(reSlice, int32(value))
		}
		return reSlice, nil
	}

	if maxKeyInStorage > int(from) && maxKeyInStorage < int(to) {
		fiboSlice, err := s.rConn.GetSliceDigits(ctx, int(from), maxKeyInStorage)
		if err != nil {
			return nil, err
		}
		calculatedDigits := calculate(fiboSlice[len(fiboSlice)-1], fiboSlice[len(fiboSlice)-2], to-int32(maxKeyInStorage))
		err = s.rConn.SetArray(ctx, int32(maxKeyInStorage+1), to, calculatedDigits)
		if err != nil {
			return nil, err
		}
		reSlice := make([]int32, 0, len(fiboSlice))
		for _, value := range fiboSlice {
			reSlice = append(reSlice, int32(value))
		}
		return append(reSlice, calculatedDigits...), nil
	}

	if maxKeyInStorage < int(from) && from >= 2 {
		fiboSlice, err := s.rConn.GetSliceDigits(ctx, maxKeyInStorage-1, maxKeyInStorage)
		if err != nil {
			return nil, err
		}
		calculatedDigits := calculate(fiboSlice[len(fiboSlice)-1],
			fiboSlice[len(fiboSlice)-2], to-int32(maxKeyInStorage))
		reSlice := make([]int32, 0, len(fiboSlice))
		for _, value := range fiboSlice {
			reSlice = append(reSlice, int32(value))
		}
		err = s.rConn.SetArray(ctx, int32(maxKeyInStorage+1), to, calculatedDigits)
		if err != nil {
			return nil, err
		}
		return calculatedDigits[from-int32(maxKeyInStorage)-1:], nil
	}

	return nil, err
}

func calculate(x1, x2 int, count int32) []int32 {
	reSlice := make([]int32, count)
	for i := 0; int32(i) < count; i++ {
		reSlice[i] = int32(x1 + x2)
		x1, x2 = x2, x2+x1
	}
	return reSlice
}

func validate(from, to int32) bool {
	if from > to || to < 0 || from < 0 {
		return false
	} else {
		return true
	}
}

//func (s *server) findDigitsInStorage(ctx context.Context,from, to int32) ([]int32, error){
//	maxInRedis, err := s.rConn.GetMax()
//	if err != nil {
//		log.Println(err)
//		maxInRedis = 0
//	}
//	if maxInRedis > int(from) && maxInRedis > int(to) {
//		reSlice := make([]int32,0,to - from + 1)
//		for ;to < from;from++{
//			digit, err := s.rConn.GetDigit(int(from))
//			if err != nil {
//				return nil, err
//			}
//			reSlice = append(reSlice,int32(digit))
//		}
//		return reSlice, nil
//	}
//
//	if maxInRedis > int(from) && maxInRedis < int(to){
//		reSlice := make([]int32,0,to - from + 1)
//		for ;int32(maxInRedis) < from;from++{
//			digit, err := s.rConn.GetDigit(int(from))
//			if err != nil {
//				return nil, err
//			}
//			reSlice = append(reSlice,int32(digit))
//		}
//		return reSlice, nil
//	}
//
//	if maxInRedis >= 2 {
//		second, err := s.rConn.GetDigit(maxInRedis)
//		if err != nil {
//			return nil, err
//		}
//		first, err := s.rConn.GetDigit(maxInRedis-1)
//		if err != nil {
//			return nil, err
//		}
//		reSlice := make([]int32,2)
//		reSlice[0]=int32(first)
//		reSlice[1]=int32(second)
//		return reSlice, nil
//	}
//	if maxInRedis == 1 {
//		first, err := s.rConn.GetDigit(maxInRedis)
//		if err != nil {
//			return nil, err
//		}
//		reSlice := make([]int32,1)
//		reSlice[0]=int32(first)
//	}
//
//	return nil, nil
//}
