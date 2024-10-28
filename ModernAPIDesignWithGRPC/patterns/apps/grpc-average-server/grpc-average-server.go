package grpcaverageserver

import (
	"fmt"
	"io"
	"log"
	"net"
	"patterns/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	proto.UnimplementedAverageServiceServer
}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {
	serverPort := 50052
	serverAddr := fmt.Sprintf("0.0.0.0:%d", serverPort)

	fmt.Println("starting average gRPC server at", serverAddr)

	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterAverageServiceServer(server, app)
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

// FindAverage implements proto.AverageServiceServer.
func (app *App) FindAverage(stream grpc.ClientStreamingServer[proto.AverageRequest, proto.AverageResponse]) error {
	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&proto.AverageResponse{
					Average: int32(sum / count),
				},
			)
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}
		count += 1
		sum += int(req.GetNumber())
	}
}
