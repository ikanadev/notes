package grpcmaxserver

import (
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"patterns/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	proto.UnimplementedMaxServiceServer
}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {
	serverPort := 50053
	serverAddr := fmt.Sprintf("0.0.0.0:%d", serverPort)

	fmt.Println("starting max gRPC server at", serverAddr)

	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterMaxServiceServer(server, app)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (app *App) Shutdown() {}

// FindMax implements proto.MaxServiceServer.
func (app *App) FindMax(stream grpc.BidiStreamingServer[proto.MaxRequest, proto.MaxResponse]) error {
	max := int32(math.MinInt32)
	mu := sync.Mutex{}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			mu.Lock()
			currentMax := max
			mu.Unlock()
			err := stream.Send(&proto.MaxResponse{Max: currentMax})
			if err != nil {
				log.Fatalf("Error while sending to client: %v", err)
				return
			}
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
			return err
		}

		num := req.GetNumber()
		mu.Lock()
		if max < num {
			max = num
		}
		mu.Unlock()
	}
}
