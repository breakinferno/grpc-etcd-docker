package main

import (
	"fmt"
	"net"
	pb "reverse/proto"
	"time"

	grpcUtils "github.com/sundogrd/gopkg/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	endpoints := "docker.for.mac.host.internal:2379"
	resolver, err := grpcUtils.NewGrpcResolover(endpoints)
	if err != nil {
		fmt.Printf("[reverse-grpc] NewGrpcResolover err: %s\n", err.Error())
		panic(err)
	}

	err = grpcUtils.ResgiterServer(*resolver, "sundog.reverse", "127.0.0.1:5300", 5*time.Second, 5)
	if err != nil {
		fmt.Printf("[reverse-grpc] RegisterServer err: %s\n", err.Error())
		panic(err)
	}

	pb.RegisterReverseServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	n := 0

	rune := make([]rune, len(request.Message))
	for _, r := range request.Message {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	output := string(rune)
	response = &pb.Response{
		Message: output,
	}
	return response, nil
}
