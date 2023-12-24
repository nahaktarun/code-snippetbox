package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	// parse the flag
	flag.Parse()

	// structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// initialize a new instance of our application struct , containing the dependencies

	app := &application{logger: logger}

	logger.Info("Starting server on ", "addr", *addr)

	// Create mux and attach with the handler routes() method
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
