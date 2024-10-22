package main

import restfibonacciclient "patterns/apps/rest-fibonacci-client"

func main() {
	app := restfibonacciclient.NewApp()
	app.Start()
}
