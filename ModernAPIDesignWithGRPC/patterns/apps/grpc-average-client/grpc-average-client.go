package grpcaverageclient

import (
	"context"
	"flag"
	"log"
	"patterns/proto"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var numbers string

func init() {
	flag.StringVar(&numbers, "numbers", "1,2,3,4,5", "comma separated numbers")
}

type App struct {
	clientConn *grpc.ClientConn
}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {
	var err error
	flag.Parse()

	serverAddr := "localhost:50052"

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	app.clientConn, err = grpc.NewClient(serverAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	averageServiceClient := proto.NewAverageServiceClient(app.clientConn)
	numberArr := make([]int, 0)

	strNumbers := strings.Split(strings.TrimSpace(numbers), ",")
	for _, strNumber := range strNumbers {
		intNum, _ := strconv.Atoi(strings.TrimSpace(strNumber))
		numberArr = append(numberArr, intNum)
	}
	doClientStreaming(averageServiceClient, numberArr)
}

func (app *App) Shutdown() {
	app.clientConn.Close()
}

func doClientStreaming(client proto.AverageServiceClient, numbers []int) {
	requests := make([]*proto.AverageRequest, 0)

	for _, n := range numbers {
		req := &proto.AverageRequest{Number: int32(n)}
		requests = append(requests, req)
	}

	stream, err := client.FindAverage(context.Background())
	if err != nil {
		log.Fatalf("error calling FindAverage rpc: %v", err)
	}

	for _, req := range requests {
		log.Printf("sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}

	log.Printf("FindAverage response %v\n", res)
}
