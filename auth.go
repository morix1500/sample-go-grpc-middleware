package main

import (
	"context"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	pb "github.com/morix1500/sample-go-grpc-middleware/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net"
)

type HelloService struct{}

func (h HelloService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, " + in.Name,
	}, nil
}

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authFunc)),
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

func authFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	fmt.Printf("receive token: %s\n", token)
	if token != "hoge" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid token")
	}
	newCtx := context.WithValue(ctx, "result", "ok")
	return newCtx, nil
}
