package main

import grpcaverageserver "patterns/apps/grpc-average-server"

func main() {
	app := grpcaverageserver.NewApp()
	app.Start()
}
