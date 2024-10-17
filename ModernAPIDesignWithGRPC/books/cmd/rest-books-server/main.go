package main

import restbooksserver "books/internal/apps/rest-books-server"

func main() {
	app := restbooksserver.NewApp()
	app.Start()
	app.Shutdown()
}
