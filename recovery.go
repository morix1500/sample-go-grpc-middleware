package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	pb "github.com/morix1500/sample-go-grpc-middleware/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net"
)

type HelloService struct{}

func (h HelloService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	if in.Name == "panic" {
		panic("failed")
	}

	return &pb.HelloResponse{
		Message: "Hello, " + in.Name,
	}, nil
}

func main() {
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	pb.RegisterHelloServiceServer(s, HelloService{})

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func recoveryFunc(p interface{}) error {
	fmt.Printf("p: %+v\n", p)
	return grpc.Errorf(codes.Internal, "Unexpected error")
}
