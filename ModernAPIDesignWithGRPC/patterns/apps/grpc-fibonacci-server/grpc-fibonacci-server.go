package grpcfibonacciserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"patterns/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	proto.UnimplementedFibonacciServiceServer
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	serverPort := 50051
	serverAddr := fmt.Sprintf("0.0.0.0:%d", serverPort)

	fmt.Println("starting fibonacci gRPC server at", serverAddr)

	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterFibonacciServiceServer(server, a)

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {}

// SyncFibonacci implements proto.FibonacciServiceServer.
func (a *App) SyncFibonacci(ctx context.Context, req *proto.FibonacciRequest) (*proto.SyncFibonacciResponse, error) {
	numFibonaccci := int(req.GetNumber())
	fibonacciNumbers := make([]int32, numFibonaccci)

	now := time.Now()
	for i := 0; i < numFibonaccci; i++ {
		fibonacciNumbers[i] = int32(fib(i))
	}
	timeTaken := time.Since(now).Milliseconds()

	resp := proto.SyncFibonacciResponse{
		TimeTaken:        fmt.Sprintf("%v milliseconds", timeTaken),
		FibonacciNumbers: fibonacciNumbers,
	}
	return &resp, nil
}

// AsyncFibonacci implements proto.FibonacciServiceServer.
func (a *App) AsyncFibonacci(req *proto.FibonacciRequest, stream grpc.ServerStreamingServer[proto.AsyncFibonacciResponse]) error {
	numFibonacci := int(req.GetNumber())
	for i := 0; i < numFibonacci; i++ {
		fibI := int32(fib(i))
		resp := &proto.AsyncFibonacciResponse{
			Sequence:        int32(i + 1),
			FibonacciNumber: fibI,
		}
		stream.Send(resp)
	}
	return nil
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
