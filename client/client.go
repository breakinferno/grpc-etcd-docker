package main

import (
	"context"
	"fmt"
	"os"
	pb "reverse/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {

	args := os.Args
	conn, err := grpc.Dial("127.0.0.1:5300", grpc.WithInsecure())
	fmt.Println("client ")

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewReverseClient(conn)
	request := &pb.Request{
		Message: args[1],
	}
	fmt.Println("client here")
	response, err := client.Do(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(response.Message)
}
