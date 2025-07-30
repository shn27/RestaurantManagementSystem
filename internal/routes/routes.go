package routes

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/riandyrn/otelchi"
	"github.com/shn27/RestaurantManagementSystem/internal/handlers"
	"github.com/shn27/RestaurantManagementSystem/internal/tracing"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func AddRoute(db *gorm.DB, es *elasticsearch.Client) {
	shutdown := tracing.InitTracer("chi-service")
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown tracer: %v", err)
		}
	}()

	r := chi.NewRouter()

	// Add the OpenTelemetry middleware
	r.Use(otelchi.Middleware("chi-service"))

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
