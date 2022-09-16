package main

import (
	"fmt"
	"log"
	"net"
	"trade-engine/internal/model/config"
	"trade-engine/internal/model/server"
	"trade-engine/pkg/protocol"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	s := server.Server{}
	protocol.RegisterExchangeServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
