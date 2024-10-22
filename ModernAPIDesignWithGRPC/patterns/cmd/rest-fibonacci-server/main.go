package main

import restfibbonacciserver "patterns/apps/rest-fibonacci-server"

func main() {
	app := restfibbonacciserver.NewApp()
	app.Start()
}
