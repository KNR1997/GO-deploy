package main

import (
	"fmt"
	"log"
	"net/http"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) run() error {
	mux := http.NewServeMux()

	// Register a simple handler at "/"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Choreo!")
	})

	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	log.Printf("server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
