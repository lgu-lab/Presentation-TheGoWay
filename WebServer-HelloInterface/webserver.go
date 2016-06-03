package main

import (
	"fmt"
	"log"
	"net/http"
)

type historyDB []int

func (db historyDB) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/history":
		fmt.Fprintf(w, "History:\n")
		for _, t := range db {
			fmt.Fprintf(w, "%v\n", t)
		}
	default:
		fmt.Fprintf(w, "Unknown request\n")
	}
}

func main() {
	db := historyDB{1, 2, 3, 4, 5}
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
