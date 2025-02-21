package routes

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shn27/RestaurantManagementSystem/internal/handlers"
	"gorm.io/gorm"
	"net/http"
)

func AddRoute(db *gorm.DB) {
	fmt.Println("AddRoutes Called")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/restaurants", func(r chi.Router) {
		r.Get("/open", handlers.GetOpenRestaurants(db))
		r.Get("/top", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("List of posts"))
		})
	})

	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/purchase", func(w http.ResponseWriter, r *http.Request) {})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
