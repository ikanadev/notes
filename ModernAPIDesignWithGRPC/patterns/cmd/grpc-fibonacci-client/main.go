package main

import grpcfibonacciclient "patterns/apps/grpc-fibonacci-client"

func main() {
	app := grpcfibonacciclient.NewApp()
	app.Start()
}
