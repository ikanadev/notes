package main

import grpcbooksclient "books/internal/apps/grpc-books-client"

func main() {
	app := grpcbooksclient.NewApp()
	app.Start()
	app.Shutdown()
}
