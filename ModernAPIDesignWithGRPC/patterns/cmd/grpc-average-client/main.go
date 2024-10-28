package main

import grpcaverageclient "patterns/apps/grpc-average-client"

func main() {
	app := grpcaverageclient.NewApp()
	app.Start()
	app.Shutdown()
}
