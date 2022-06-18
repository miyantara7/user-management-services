package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func RunServer() {

	opts := []grpc.ServerOption{}

	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionAge:      5 * time.Minute,
		MaxConnectionIdle:     1 * time.Hour,
		MaxConnectionAgeGrace: 5 * time.Second,
	}))

	grpcServer := grpc.NewServer(opts...)

	reflection.Register(grpcServer)

	svcHost := "localhost"
	svcPort := 9901

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svcHost, svcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start notpool Service gRPC server: %v", err)
		}
	}()

	fmt.Printf("notpool Service gRPC server is running at %s:%d\n", svcHost, svcPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())

}
