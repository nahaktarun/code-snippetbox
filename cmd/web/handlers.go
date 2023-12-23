package main

import (
	"fmt"
	"html/template"
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
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
	}

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
