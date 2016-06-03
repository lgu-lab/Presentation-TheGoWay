package main

import (
	"fmt"
	"log"
	"net/http"
)

var historyDb = []int{1, 2, 3, 4, 5}

func httpHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ History: %v }", historyDb)
}

func httpHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {

	http.HandleFunc("/history", httpHistory)
	http.HandleFunc("/", httpHello)
	{
		// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 	if r.URL.Path != "/" {
		// 		http.NotFound(w, r)
		// 		return
		// 	}
		// 	httpHello(w, r)
		// })
	}
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
