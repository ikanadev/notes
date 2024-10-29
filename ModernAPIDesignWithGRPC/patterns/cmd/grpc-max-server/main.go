package main

import grpcmaxserver "patterns/apps/grpc-max-server"

func main() {
	app := grpcmaxserver.NewApp()
	app.Start()
	app.Shutdown()
}
