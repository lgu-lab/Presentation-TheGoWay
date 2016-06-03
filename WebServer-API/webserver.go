package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SECTION: Start API Handlers Section

func addAPIRoutes(apiRouter *mux.Router) {
	apiRouter.Path("/history").Methods("GET").HandlerFunc(httpHistory)
	apiRouter.HandleFunc("/compute/{n}/{m}", httpCompute)

	//Version is the structure for the Version API
	type Version struct {
		Version string
	}
	version := Version{Version: "v1"}
	apiRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(version)
	})

}

// END SECTION

func httpHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", httpHello)

	addAPIRoutes(router.PathPrefix("/api").Subrouter())

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
