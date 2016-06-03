package main

import (
	"fmt"
	"log"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", httpHandler)
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
