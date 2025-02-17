package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func addRoute() {
	r := chi.NewRouter()
	r.Get("/restaurants/open", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/restaurants/top", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/purchase", func(w http.ResponseWriter, r *http.Request) {})
}
