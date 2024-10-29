package main

import grpcmaxclient "patterns/apps/grpc-max-client"

func main() {
	app := grpcmaxclient.NewApp()
	app.Start()
	app.Shutdown()
}
