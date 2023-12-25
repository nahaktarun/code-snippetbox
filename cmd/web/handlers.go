package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.tarunnahak.in/internal/models"
)

// Define a home handler function which writes a byte slice containing response body

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return

	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, "home.tmpl.html", templateData{Snippets: snippets})
}

// Add a snippetView handler function
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	app.render(w, r, http.StatusOK, "view.tmpl.html", templateData{Snippet: snippet})

}

// Add a snippetCreate handler function
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "o snail test"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
