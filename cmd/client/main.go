package main

import (
	"context"
	"fmt"
	pb "github.com/keepcalmist/grpcFibonacci/pkg/grpc"
	"google.golang.org/grpc"
	"log"
)

const address = "127.0.0.1:5300"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("error:", err)
	}

	defer conn.Close()

	client := pb.NewFibonacciClient(conn)

	req, err := client.Get(context.Background(), &pb.Request{
		X: 1,
		Y: 10,
	})

	fmt.Println(req)
}
