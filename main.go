package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// links all the routes to the related functions
	http.HandleFunc("/create", create)
	http.HandleFunc("/open/", open)
	http.HandleFunc("/draw/", draw)
	fmt.Println("Server initialized")
	// serve the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
