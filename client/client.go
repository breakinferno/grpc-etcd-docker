package main

import (
	"context"
	"fmt"
	"os"
	pb "reverse/proto"

	grpcUtils "github.com/sundogrd/gopkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	args := os.Args
	endpoints := "docker.for.mac.host.internal:2379"
	r, err := grpcUtils.NewGrpcResolover(endpoints)

	if err != nil {
		grpclog.Fatalf("error is: %v", err)
	}
	b := grpc.RoundRobin(r)

	conn, err := grpc.Dial("sundog.reverse", grpc.WithBalancer(b), grpc.WithInsecure(), grpc.WithBlock())
	// conn, err := grpc.Dial("127.0.0.1:5300", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}

	fmt.Println("client start")

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewReverseClient(conn)
	request := &pb.Request{
		Message: args[1],
	}
	response, err := client.Do(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(response.Message)
}
