package seed

import (
	"encoding/json"
	"fmt"
	"github.com/shn27/RestaurantManagementSystem/internal/database"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var Seed = &cobra.Command{
	Use:   "seed",
	Short: "Seed data",
	Long:  "Seed data",
	Run: func(cmd *cobra.Command, args []string) {
		database.ConnectDB()
		database.MigrateDatabase()
		processMenu()
	},
}

type Dish struct {
	DishName string  `json:"dishName"`
	Price    float64 `json:"price"`
}

type Restaurant struct {
	CashBalance    float64 `json:"cashBalance"`
	Menu           []Dish  `json:"menu"`
	OpeningHours   string  `json:"openingHours"`
	RestaurantName string  `json:"restaurantName"`
}

func processMenu() {
	fmt.Println("\nprocessMenu==============")
	filePath := "data/restaurant_with_menu.json"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}
	// Unmarshal JSON data into a slice of Restaurant structs
	var restaurants []Restaurant
	err = json.Unmarshal(bytes, &restaurants)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	var menuData []database.Menu
	var restaurantsData []database.Restaurant

	restaurantID := 0
	for _, restaurant := range restaurants {
		restaurantsData = append(restaurantsData, database.Restaurant{
			RestaurantName: restaurant.RestaurantName,
			CashBalance:    restaurant.CashBalance,
		})
		restaurantID++
		for _, menu := range restaurant.Menu {
			menuData = append(menuData, database.Menu{
				RestaurantID: uint(restaurantID),
				DishName:     menu.DishName,
				Price:        menu.Price,
			})
		}
	}
	database.DB.Create(&restaurantsData)
	database.DB.Create(&menuData)
}
