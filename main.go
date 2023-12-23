package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID : %d", id)
}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("allow", http.MethodPost)

		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("Create a new Snippet"))
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