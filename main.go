package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	cache "siper92/mansion-assessment/internal"

	"github.com/gorilla/mux"
)

const ListenUrl = "localhost:8080"

func getLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	location := cache.GetCacheContext("location")
	fmt.Print(location)
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getLocations).Methods(http.MethodPost)

	fmt.Printf("\n\t>>> Listening for POST requests on %s\n\n", ListenUrl)
	log.Fatal(http.ListenAndServe(ListenUrl, r))
}
