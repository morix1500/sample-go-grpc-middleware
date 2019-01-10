package main

import (
	"context"
	"fmt"
	pb "github.com/morix1500/sample-go-grpc-middleware/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	in := &pb.HelloRequest{
		Name: "Morix",
	}
	res, err := client.Hello(context.Background(), in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}
