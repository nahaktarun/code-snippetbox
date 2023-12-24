package main

import (
	"net/http"
)

// Server Error helper writes a log entry at error level
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// ClientError helper sents a specific status code and corresponding error description

func (app *application) clientError(w http.ResponseWriter, status int) {

	http.Error(w, http.StatusText(status), status)

}

// notFound helper
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
