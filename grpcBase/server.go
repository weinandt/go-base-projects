package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/weinandt/go-base-projects/grpcBase/user"
	"google.golang.org/grpc"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserServiceServer) GetUser(ctx context.Context, userRequest *pb.UserRequest) (*pb.UserReponse, error) {
	fmt.Println(userRequest.Id)

	userName := "Nick"

	return &pb.UserReponse{
		Id:   userRequest.GetId(),
		Name: userName,
	}, nil
}

func main() {
	hostPort := "localhost:50000"
	listener, err := net.Listen("tcp", hostPort)
	if err != nil {
		fmt.Println("Failed to set up listener.")
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, &UserServiceServer{})
	fmt.Println("Starting server on", hostPort)
	if err := server.Serve(listener); err != nil {
		panic("Could not serve grpc endpoint")
	}
}
