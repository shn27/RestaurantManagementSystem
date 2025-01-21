package main

import "github.com/shn27/RestaurantManagementSystem/internal/database"

func main() {
	err := database.Connection.Execute()
	if err != nil {
		return
	}
}
