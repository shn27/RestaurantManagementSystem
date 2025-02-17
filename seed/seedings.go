package seed

import (
	"encoding/json"
	"fmt"
	"github.com/shn27/RestaurantManagementSystem/internal/database"
	"github.com/shn27/RestaurantManagementSystem/seed/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const FILEPATH_USERS_WITH_PURCHASE_HISTORY = "data/users_with_purchase_history.json"
const FILEPATH_RESTAURANT_WITH_MENU = "data/restaurant_with_menu.json"

var Seed = &cobra.Command{
	Use:   "seed",
	Short: "Seed data",
	Long:  "Seed data",
	Run: func(cmd *cobra.Command, args []string) {
		processMenu()
		processUsersWithPurChaseHistory()
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

type Users struct {
	CashBalance     float64           `json:"cashBalance"`
	ID              uint64            `json:"id"`
	Name            string            `json:"name"`
	PurchaseHistory []PurchaseHistory `json:"purchaseHistory"`
}

type PurchaseHistory struct {
	DishName          string  `json:"dishName"`
	RestaurantName    string  `json:"restaurantName"`
	TransactionAmount float64 `json:"transactionAmount"`
	TransactionDate   string  `json:"transactionDate"`
}

func processOpeningHours(str string, restaurantID int) []database.OpeningHours {
	var data []database.OpeningHours
	str = strings.ToLower(str)
	words := strings.Fields(str)
	length := len(words)
	for i := 3; i < length-2; i++ {
		word := words[i]
		if word == "-" {
			closingTime := words[i+1]
			if closingTime[0] >= '0' && closingTime[0] <= '9' {
				closingTimeAmPM := words[i+2]
				openingTimeAmPM := words[i-1]
				openingTime := words[i-2]
				day := words[i-3]
				flag, day, openingHour, openingMinite, closingHour, closingMinite := utils.CheckValidity(day, closingTime, openingTime, openingTimeAmPM, closingTimeAmPM)
				if !flag {
					continue
				}
				openTime := strconv.Itoa(openingHour) + ":" + strconv.Itoa(openingMinite) + ":00"
				closeTime := strconv.Itoa(closingHour) + ":" + strconv.Itoa(closingMinite) + ":00"

				data = append(data, database.OpeningHours{
					Day:          day,
					OpeningTime:  openTime,
					ClosingTime:  closeTime,
					RestaurantID: uint(restaurantID),
				})
			}
		}
	}
	return data
}

func processMenu() {
	fmt.Println("process menu...")

	filePath := FILEPATH_RESTAURANT_WITH_MENU
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
	var OpeningHoursData []database.OpeningHours

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
		OpeningHoursData = append(OpeningHoursData, processOpeningHours(restaurant.OpeningHours, restaurantID)...)
	}
	database.DB.Create(&restaurantsData)
	database.DB.Create(&menuData)
	database.DB.Create(&OpeningHoursData)
	fmt.Println("Successfully processed menu!")
}

func processUsersWithPurChaseHistory() {
	fmt.Println("process Users with purChase history...")

	filePath := FILEPATH_USERS_WITH_PURCHASE_HISTORY
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
	var users []Users
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	var purchaseHistoryData []database.PurchaseHistory
	var usersData []database.User

	for _, user := range users {
		usersData = append(usersData, database.User{
			ID:          uint(user.ID),
			UserName:    user.Name,
			CashBalance: user.CashBalance,
		})
		for _, purchaseHistory := range user.PurchaseHistory {
			layout := "01/02/2006 03:04 PM" //todo
			parsedTime, err := time.Parse(layout, purchaseHistory.TransactionDate)
			if err != nil {
				fmt.Println("Error parsing date:", err)
				return
			}
			purchaseHistoryData = append(purchaseHistoryData, database.PurchaseHistory{
				UserID:            uint(user.ID),
				DishName:          purchaseHistory.DishName,
				RestaurantName:    purchaseHistory.RestaurantName,
				TransactionAmount: purchaseHistory.TransactionAmount,
				Time:              parsedTime,
			})
		}
	}
	database.DB.Create(&usersData)
	database.DB.Create(&purchaseHistoryData)
	fmt.Println("Successfully processed purchase history!")
}
