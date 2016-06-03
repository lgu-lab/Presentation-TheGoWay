package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//HistoryDb is where the History will be cached, potentially a database
type HistoryDb struct {
	History []int `json:"History"`
}

var historyDb = HistoryDb{History: []int{1, 2, 3, 4, 5}}

func httpHistory(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(historyDb)
}

func httpHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/history", httpHistory)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		httpHello(w, r)
	})

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
