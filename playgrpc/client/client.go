package main

import (
	"context"
	"fmt"
	"log"
	pb "playgrpc/github.com/vivek781113/gen/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "hello from grpc"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Msg)
}
