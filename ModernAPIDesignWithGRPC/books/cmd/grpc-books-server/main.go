package main

import grpcbooksserver "books/internal/apps/grpc-books-server"

func main() {
	app := grpcbooksserver.NewApp()
	app.Start()
	app.Shutdown()
}
