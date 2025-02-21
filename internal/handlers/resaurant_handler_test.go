package handlers

import (
	"fmt"
	"gorm.io/gorm"
)

func showTablesAndColumns(db *gorm.DB) {
	var tables []string

	// Query to get all table names
	db.Raw("SHOW TABLES").Scan(&tables)

	fmt.Println("Tables in the database:")
	for _, table := range tables {
		fmt.Println("\nTable:", table)

		// Query to get column names for each table
		var columns []struct {
			Field string `gorm:"column:Field"`
		}
		db.Raw("SHOW COLUMNS FROM " + table).Scan(&columns)

		fmt.Println("Columns:")
		for _, col := range columns {
			fmt.Println(" -", col.Field)
		}
	}
}

func getOpeningHours(db *gorm.DB) {
	var openingHours []struct {
		Day         string `json:"day"`
		OpeningTime string `json:"opening_time"`
		ClosingTime string `json:"closing_time"`
	}
	var tot int
	query := "SELECT COUNT(*) FROM restaurants;"
	db.Raw(query).Scan(&tot)
	fmt.Println("restaurants TOTAL:", tot)

	query = "SELECT COUNT(*) FROM users;"
	db.Raw(query).Scan(&tot)
	fmt.Println("users TOTAL:", tot)

	query = "SELECT COUNT(*) FROM purchase_histories;"
	db.Raw(query).Scan(&tot)
	fmt.Println("purchase_histories TOTAL:", tot)

	query = "SELECT COUNT(*) FROM menus;"
	db.Raw(query).Scan(&tot)
	fmt.Println("menus TOTAL:", tot)

	query = "SELECT COUNT(*) FROM opening_hours;"
	db.Raw(query).Scan(&tot)
	fmt.Println("opening_hours TOTAL:", tot)

	fmt.Println("Opening Hours:")
	for _, oh := range openingHours {
		fmt.Printf("Day: %s, Open: %s, Close: %s\n", oh.Day, oh.OpeningTime, oh.ClosingTime)
	}
}
