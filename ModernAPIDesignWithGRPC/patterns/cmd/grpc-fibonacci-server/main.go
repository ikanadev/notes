package main

import grpcfibonacciserver "patterns/apps/grpc-fibonacci-server"

func main() {
	app := grpcfibonacciserver.NewApp()
	app.Start()
}
