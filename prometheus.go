package main

import (
	"context"
	"fmt"
	grpc_prome "github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/morix1500/sample-go-grpc-middleware/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	grpcMetrics = grpc_prome.NewServerMetrics()

	reg = prometheus.NewRegistry()

	customizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "demo_server_say_hello_method_handle_count",
		Help: "Total number of RPCs handled on the server.",
	}, []string{"name"})
)

type HelloService struct{}

func (h HelloService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Custom Metrics Count up
	customizedCounterMetric.WithLabelValues("Test").Inc()
	return &pb.HelloResponse{
		Message: "Hello, " + in.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	// カスタムメトリクス登録
	reg.MustRegister(grpcMetrics, customizedCounterMetric)
	customizedCounterMetric.WithLabelValues("Test")

	// create http server
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    fmt.Sprintf("0.0.0.0:%d", 5001),
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)
	pb.RegisterHelloServiceServer(s, HelloService{})

	// メトリクス初期化
	grpcMetrics.InitializeMetrics(s)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
