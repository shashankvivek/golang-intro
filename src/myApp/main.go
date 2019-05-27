package main

import (
	"log"
	"net/http"

	app "myApp/routes"
)

func main() {
	handleRequest()
}

func handleRequest() {
	r := app.NewRouter()
	// run using : sudo -E go run main.go
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", r))
}
