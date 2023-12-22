package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing response body

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {

	// Create mux and attach with the handler function
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
