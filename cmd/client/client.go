package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/rafaelyokota/grpc_example/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	AddBatchUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "rafael",
		Email: "rafael.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "rafael",
		Email: "rafael.com",
	}
	responseStream, err := client.AddUserVebose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not received the message: %v", err)
		}
		fmt.Println("Status:", stream.Status, "-", stream.GetUser())
	}
}

func AddBatchUser(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "Rafael",
			Email: "rafael.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "Rafael 2",
			Email: "rafael.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "Rafael 3",
			Email: "rafael.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "Rafael 4",
			Email: "rafael.com",
		},
	}
	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		fmt.Println("Sending Information:", req)
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error Receiving  response: %v", err)
	}

	fmt.Println(res)
}
