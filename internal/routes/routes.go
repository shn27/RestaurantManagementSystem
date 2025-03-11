package routes

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shn27/RestaurantManagementSystem/internal/handlers"
	"gorm.io/gorm"
	"net/http"
)

func AddRoute(db *gorm.DB, es *elasticsearch.Client) {
	fmt.Println("AddRoutes Called")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/restaurants", func(r chi.Router) {
		r.Get("/open", handlers.GetOpenRestaurants(db))
		r.Get("/top", handlers.ListTopRestaurants(db))
	})

	r.Get("/search", handlers.Search(es, "names"))
	r.Post("/purchase", handlers.ProcessPurchase(db))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
