package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shn27/RestaurantManagementSystem/internal/database"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

		// Create Redis cache key
		cacheKey := fmt.Sprintf("open_restaurants:%s:%s", dayName, currentTime)

		// Check Redis for cached data
		ctx := context.Background()
		cachedData, err := database.RedisClient.Get(ctx, cacheKey).Result()
		if err == nil {
			// Cache hit: return cached data
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cachedData))
			w.WriteHeader(http.StatusOK)
			fmt.Println("cache hit")
			return
		}
		fmt.Println("cache miss ", cacheKey)

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
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		database.RedisClient.Set(context.Background(), cacheKey, responseJSON, 0)

		w.Write(responseJSON)
		w.WriteHeader(http.StatusOK)
	}
}

func ListTopRestaurants(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		minPrice := r.URL.Query().Get("minPrice")
		if minPrice == "" {
			http.Error(w, "minPrice query parameter is required", http.StatusBadRequest)
			return
		}
		minPriceInt, err := strconv.Atoi(minPrice)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid minPrice format: %v", err), http.StatusBadRequest)
			return
		}

		maxPrice := r.URL.Query().Get("maxPrice")
		if maxPrice == "" {
			http.Error(w, "maxPrice query parameter is required", http.StatusBadRequest)
			return
		}
		maxPriceInt, err := strconv.Atoi(maxPrice)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid maxPrice format: %v", err), http.StatusBadRequest)
			return
		}

		numDish := r.URL.Query().Get("numDish")
		if numDish == "" {
			http.Error(w, "numDish query parameter is required", http.StatusBadRequest)
			return
		}
		numDishInt, err := strconv.Atoi(numDish)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid numDish format: %v", err), http.StatusBadRequest)
			return
		}

		isMore := r.URL.Query().Get("isMore")
		if isMore == "" {
			http.Error(w, "isMore query parameter is required", http.StatusBadRequest)
			return
		}
		isMoreBool, err := strconv.ParseBool(isMore)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid isMore format: %v", err), http.StatusBadRequest)
			return
		}

		limit := r.URL.Query().Get("limit")
		if minPrice == "" {
			http.Error(w, "limit query parameter is required", http.StatusBadRequest)
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid limit format: %v", err), http.StatusBadRequest)
			return
		}

		// Create Redis cache key
		cacheKey := fmt.Sprintf("top_restaurants:numDish=%d:minPrice=%.2f:maxPrice=%.2f:isMore=%v:limit=%d",
			numDish, minPrice, maxPrice, isMore, limit)

		// Check Redis for cached data
		ctx := context.Background()
		cachedData, err := database.RedisClient.Get(ctx, cacheKey).Result()
		if err == nil {
			// Cache hit: return cached data
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cachedData))
			w.WriteHeader(http.StatusOK)
			fmt.Println("cache hit")
			return
		}
		fmt.Println("cache miss")

		var restaurantNames []string
		query := `
SELECT r.restaurant_name
FROM restaurants r
JOIN (
    SELECT restaurant_id, COUNT(*) as dish_count
    FROM menus
    WHERE price BETWEEN ? AND ?
    GROUP BY restaurant_id
) m ON r.id = m.restaurant_id
WHERE m.dish_count > ? 
ORDER BY r.restaurant_name ASC
LIMIT ?;
`

		query1 := `
SELECT r.restaurant_name
FROM restaurants r
JOIN (
    SELECT restaurant_id, COUNT(*) as dish_count
    FROM menus
    WHERE price BETWEEN ? AND ?
    GROUP BY restaurant_id
) m ON r.id = m.restaurant_id
WHERE m.dish_count < ?
ORDER BY r.restaurant_name ASC
LIMIT ?;
`
		if isMoreBool {
			err = db.Raw(query, float64(minPriceInt), float64(maxPriceInt), numDishInt, limitInt).Scan(&restaurantNames).Error
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			err = db.Raw(query1, float64(minPriceInt), float64(maxPriceInt), numDishInt, limitInt).Scan(&restaurantNames).Error
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		response := map[string]interface{}{
			"open_restaurants": restaurantNames,
		}
		responseJSON, err := json.Marshal(response)

		// add cache
		database.RedisClient.Set(context.Background(), cacheKey, responseJSON, 0)

		w.Write(responseJSON)
		w.WriteHeader(http.StatusOK)
	}
}
