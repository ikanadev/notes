package grpcfibonacciclient

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"patterns/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var typeOfCall string
var number int

func init() {
	flag.StringVar(&typeOfCall, "typeOfCall", "sync", "sync or async call?")
	flag.IntVar(&number, "number", 10, "for what number do you want the sequence?")
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

	serverAddr := "localhost:50051"

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	app.clientConn, err = grpc.NewClient(serverAddr, opts)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	fibonacciServiceClient := proto.NewFibonacciServiceClient(app.clientConn)

	if typeOfCall == "sync" {
		syncFibonacci(fibonacciServiceClient, number)
	}
	if typeOfCall == "async" {
		asyncFibonacci(fibonacciServiceClient, number)
	}
}

func syncFibonacci(client proto.FibonacciServiceClient, number int) {
	req := &proto.FibonacciRequest{Number: int32(number)}
	res, err := client.SyncFibonacci(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling fibonacci rpc: %v", err)
	}

	fibResp := ""
	for _, n := range res.FibonacciNumbers {
		fibResp = fmt.Sprintf("%s,%d", fibResp, n)
	}

	log.Printf("Sync fibonacci sequence result for %d: %v, and time taken was: %s\n", number, fibResp, res.TimeTaken)
}

func asyncFibonacci(client proto.FibonacciServiceClient, number int) {
	req := &proto.FibonacciRequest{Number: int32(number)}

	resStream, err := client.AsyncFibonacci(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling fibonacci stream rpc: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream from fibonacci rpc: %v", err)
		}
		log.Printf("Response from fibo stream: sequence(%d), fibonacci number(%d)", msg.Sequence, msg.FibonacciNumber)
	}
}
