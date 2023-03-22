package main

import (
	"adcash/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/api/loans", handlers.Repo.Loans)
	mux.Post("/api/apply", handlers.Repo.Apply)

	return mux
}
