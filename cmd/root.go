package cmd

import (
	"fmt"
	"github.com/shn27/RestaurantManagementSystem/internal/database"
	"github.com/shn27/RestaurantManagementSystem/internal/routes"
	"github.com/shn27/RestaurantManagementSystem/seed"
	"github.com/spf13/cobra"
	"os"
)

var main = &cobra.Command{
	Use:   "restaurantmanagement",
	Short: "RestaurantManagementSystem",
	Long:  `RestaurantManagementSystem`,
	Run: func(cmd *cobra.Command, args []string) {
		database.ConnectDB()
		database.ConnectRedis()
		routes.AddRoute(database.DB)
	},
}

var InitializeDB = &cobra.Command{
	Use:   "db",
	Short: "DB",
	Long:  `DB`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initialize DB called ...")
		if err := database.Connection.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := seed.Seed.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		database.CloseDB()
	},
}

func Execute() {
	main.AddCommand(InitializeDB)
	err := main.Execute()
	if err != nil {
		return
	}
}
