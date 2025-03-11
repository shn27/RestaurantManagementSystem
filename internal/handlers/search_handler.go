package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/shn27/RestaurantManagementSystem/internal/database"
	"io"
	"net/http"
)

func Search(esClient *elasticsearch.Client, indexName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("search")
		if query == "" {
			http.Error(w, "search query parameter is required", http.StatusBadRequest)
			return
		}

		// Create Redis cache key
		cacheKey := fmt.Sprintf("search:%s", query)

		// Check Redis for cached data
		ctx := context.Background()
		cachedData, err := database.RedisClient.Get(ctx, cacheKey).Result()
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cachedData))
			w.WriteHeader(http.StatusOK)
			return
		}

		searchBody := map[string]interface{}{
			"query": map[string]interface{}{
				"multi_match": map[string]interface{}{
					"query":  query,
					"fields": []string{"name"},
				},
			},
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(searchBody); err != nil {
			http.Error(w, "Error encoding search query", http.StatusInternalServerError)
			return
		}

		res, err := esClient.Search(
			esClient.Search.WithContext(ctx),
			esClient.Search.WithIndex(indexName),
			esClient.Search.WithBody(&buf),
			esClient.Search.WithTrackTotalHits(true),
			esClient.Search.WithPretty(),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		var response []byte
		// Extract hits
		if hits, found := result["hits"].(map[string]interface{}); found {
			if hitArray, ok := hits["hits"].([]interface{}); ok && len(hitArray) > 0 {
				for _, hit := range hitArray {
					if hitMap, valid := hit.(map[string]interface{}); valid {
						if source, exists := hitMap["_source"]; exists {
							sourceJSON, _ := json.MarshalIndent(source, "", "  ")
							w.Write(sourceJSON)
							response = append(response, sourceJSON...)
						}
					}
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}
		database.RedisClient.Set(context.Background(), cacheKey, response, 0)
		w.WriteHeader(http.StatusOK)
	}
}
