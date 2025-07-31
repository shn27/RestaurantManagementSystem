package handlers

import (
	"encoding/json"
	"github.com/shn27/RestaurantManagementSystem/internal/database"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func ProcessPurchase(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tr := otel.Tracer("restaurant-service")
		_, span := tr.Start(ctx, "process-purchase-handler")
		defer span.End()

		type purchaseRequest struct {
			DishID int `json:"dish_id"`
			UserID int `json:"user_id"`
		}
		var req purchaseRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		_, dbSpan := tr.Start(ctx, "DB: startTransaction")
		defer dbSpan.End()
		// Use a GORM transaction to ensure atomicity
		err := db.Transaction(func(tx *gorm.DB) error {
			var user database.User
			if err := tx.First(&user, req.UserID).Error; err != nil {
				return err
			}

			var dish database.Menu
			if err := tx.First(&dish, req.DishID).Error; err != nil {
				return err
			}

			if user.CashBalance < dish.Price {
				return http.ErrHandlerTimeout // Insufficient funds
			}

			var restaurant database.Restaurant
			if err := tx.First(&restaurant, dish.RestaurantID).Error; err != nil {
				return err
			}

			// Deduct from user and add to restaurant
			if err := tx.Model(&user).Update("cash_balance", user.CashBalance-dish.Price).Error; err != nil {
				return err
			}

			if err := tx.Model(&restaurant).Update("cash_balance", restaurant.CashBalance+dish.Price).Error; err != nil {
				return err
			}

			purchaseHistory := database.PurchaseHistory{
				UserID:            user.ID,
				DishName:          dish.DishName,
				RestaurantName:    restaurant.RestaurantName,
				TransactionAmount: dish.Price,
				Time:              time.Now(),
			}
			if err := tx.Create(&purchaseHistory).Error; err != nil {
				return err
			}

			return nil // Commit transaction if everything is fine
		})

		if err != nil {
			dbSpan.RecordError(err)
			http.Error(w, "Transaction failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Purchase successful"))
	}
}
