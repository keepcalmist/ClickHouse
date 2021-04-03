package main

import (
	"github.com/keepcalmist/grpcFibonacci/pkg/config"
	"github.com/keepcalmist/grpcFibonacci/pkg/fibonacci"
	"github.com/keepcalmist/grpcFibonacci/pkg/redis"
	"github.com/keepcalmist/grpcFibonacci/pkg/servers/grpcServer"
	"github.com/keepcalmist/grpcFibonacci/pkg/servers/restServer"
	"os"
	"os/signal"
)

func main() {
	conf := config.New()
	rConn := redis.New(":" + conf.GetString("REDIS_PORT"))
	fiboService := fibonacci.New(rConn)
	grpc := grpcServer.New(fiboService)
	rest := restServer.New(fiboService)
	//
	quitRest := make(chan bool)
	quitGrpc := make(chan bool)
	grpc.Run(conf, quitGrpc)
	rest.Run(conf, quitRest)
	//
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	quitGrpc <- true
	quitRest <- true
}
