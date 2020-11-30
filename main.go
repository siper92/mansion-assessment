package main

import (
	"fmt"
	"log"
	"net/http"

	"siper92/mansion-assessment/internal/handler"

	"github.com/gorilla/mux"
)

const ListenUrl = "localhost:8080"

func main() {
	fmt.Printf("\n\t>>> Initializing service")
	// fasters load, but possible a period with incomplete results
	go handler.InitCache()
	fmt.Printf("\n\t>>> Service initialized")

	r := mux.NewRouter()
	r.HandleFunc("/", handler.FilterNearbyLocations).Methods(http.MethodPost)

	fmt.Printf("\n\t>>> Listening for POST requests on %s\n\n", ListenUrl)
	log.Fatal(http.ListenAndServe(ListenUrl, r))
}
