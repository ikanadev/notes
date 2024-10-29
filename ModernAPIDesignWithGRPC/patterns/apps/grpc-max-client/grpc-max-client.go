package grpcmaxclient

import (
	"context"
	"flag"
	"io"
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
	flag.StringVar(&numbers, "numbers", "", "provide comma separated numbers to compute max")
}

type App struct {
	clientConn *grpc.ClientConn
}

func NewApp() *App {
	return &App{}
}

func (app *App) Shutdown() {
	if app.clientConn != nil {
		app.clientConn.Close()
	}
}

func (app *App) Start() {
	var err error
	flag.Parse()

	serverAddr := "localhost:50053"

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	app.clientConn, err = grpc.NewClient(serverAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}

	maxServiceClient := proto.NewMaxServiceClient(app.clientConn)
	numberArr := make([]int, 0)

	strNumbers := strings.Split(strings.TrimSpace(numbers), ",")
	for _, strNum := range strNumbers {
		intNum, _ := strconv.Atoi(strings.TrimSpace(strNum))
		numberArr = append(numberArr, intNum)
	}

	doBidirectionalStreaming(maxServiceClient, numberArr)
}

func doBidirectionalStreaming(client proto.MaxServiceClient, numbers []int) {
	requests := make([]*proto.MaxRequest, 0)

	for _, n := range numbers {
		req := &proto.MaxRequest{Number: int32(n)}
		requests = append(requests, req)
	}

	stream, err := client.FindMax(context.Background())
	if err != nil {
		log.Fatalf("error while calling FindMax: %v\n", err)
	}

	waitChan := make(chan struct{})

	go func(reqs []*proto.MaxRequest) {
		for _, req := range reqs {
			log.Printf("sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}(requests)

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving from FindMax rpc response: %v\n", err)
				break
			}
			log.Printf("response received: %v\n", res)
		}
	}()

	<-waitChan
}
