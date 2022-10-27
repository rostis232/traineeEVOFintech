package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"traineeEVOFintech/internal/handlers"
)

// routes creates new http.Handler with chi
func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Test)
	//mux.Post()
	//mux.Get()
	//mux.Get()

	return mux
}
