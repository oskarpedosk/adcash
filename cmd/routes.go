package main

import (
	"adcash/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Post("/loans", handlers.Repo.Loans)
	mux.Post("/apply", handlers.Repo.PostApply)

	return mux
}
