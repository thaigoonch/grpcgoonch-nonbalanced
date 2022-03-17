package main

import (
	"fmt"
	"net"

	grpcgoonch "github.com/thaigoonch/grpcgoonch-nonbalanced/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	fmt.Println("grpcgoonch waiting for client requests...")
	port := 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	s := grpcgoonch.Server{}

	grpcServer := grpc.NewServer()
	grpcgoonch.RegisterServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Fatalf("Failed to serve gRPC server over port %d: %v", port, err)
	}
}
