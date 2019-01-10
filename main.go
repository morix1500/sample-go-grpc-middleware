package main

import (
	"context"
	//"fmt"
	pb "github.com/morix1500/sample-go-grpc-middleware/proto"
	"google.golang.org/grpc"
	"net"
)

type HelloService struct{}

func (h HelloService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, " + in.Name,
	}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, HelloService{})

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
