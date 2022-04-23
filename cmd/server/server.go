package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rafaelyokota/grpc_example/pb"
	"github.com/rafaelyokota/grpc_example/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not Connect: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)
	fmt.Println("Running Server...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not Serve: %v", err)
	}
}
