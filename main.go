package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/open/", open)
	http.HandleFunc("/draw/", draw)
	fmt.Println("Server initialized")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
