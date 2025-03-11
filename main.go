package main

import "github.com/shn27/RestaurantManagementSystem/cmd"

func main() {
	err := cmd.InitializeDB.Execute()
	if err != nil {
		return
	}
	cmd.Execute()
}
