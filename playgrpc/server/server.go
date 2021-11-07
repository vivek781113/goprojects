package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "playgrpc/github.com/vivek781113/gen/proto"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {

	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	listener, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("grpc server is running on port %d\n", 5000)
	}

}
