package cmd

import (
	"fmt"
	"github.com/shn27/RestaurantManagementSystem/internal/database"
	"github.com/shn27/RestaurantManagementSystem/seed"
)

func ExecuteConnection() {
	if err := database.Connection; err != nil {
		fmt.Println(err)
	}
}

func ExecuteSeeds() {
	if err := seed.Seed.Execute(); err != nil {
		fmt.Println(err)
	}
}
