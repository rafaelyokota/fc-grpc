package services

import (
	"context"
	"fmt"
	"time"

	"github.com/rafaelyokota/grpc_example/pb"
)

// type UserServiceClient interface {
// 	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
// 	AddUserVebose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVeboseClient, error)
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	//Insert in dt

	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil

}

func (*UserService) AddUserVebose(req *pb.User, stream pb.UserService_AddUserVeboseServer) error {

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "1234",
			Name:  req.Name,
			Email: req.Email,
		},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Request Done",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	return nil
}
