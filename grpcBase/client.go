package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/weinandt/go-base-projects/grpcBase/user"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:50000"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userResponse, err := client.GetUser(ctx, &pb.UserRequest{Id: "myId"})
	if err != nil {
		fmt.Println(err)
		panic("Could not get response for getUser request.")
	}

	fmt.Printf("%+v\n", userResponse)
}
