package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing response body

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	}
	w.Write([]byte("Hello world from snippet box"))

}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet...."))
}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {

	// Create mux and attach with the handler function
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
