package handlers

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func GetOpenRestaurants(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		datetime := r.URL.Query().Get("datetime")
		if datetime == "" {
			http.Error(w, "datetime query parameter is required", http.StatusBadRequest)
			return
		}

		t, err := time.Parse("2006-01-02 15:04:05", datetime)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid datetime format: %v", err), http.StatusBadRequest)
			return
		}
		dayName := t.Weekday().String()
		dayName = strings.ToLower(dayName)
		currentTime := t.Format("15:04:05")

		var restaurantNames []string
		query := `
SELECT r.restaurant_name
FROM restaurants r
JOIN opening_hours oh ON r.id = oh.restaurant_id
WHERE oh.day = ?
AND ? BETWEEN oh.opening_time AND oh.closing_time;
`
		err = db.Raw(query, dayName, currentTime).Scan(&restaurantNames).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]interface{}{
			"open_restaurants": restaurantNames,
		}
		responseJSON, err := json.Marshal(response)
		w.Write(responseJSON)
		w.WriteHeader(http.StatusOK)
	}
}
