package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint    `gorm:"primary_key:autoIncrement"`
	UserName    string  `json:"user_name"`
	CashBalance float64 `json:"cash_balance"`
	gorm.Model
}

type PurchaseHistory struct {
	ID                uint      `json:"id" gorm:"primary_key:autoIncrement"`
	UserID            uint      `json:"user_id" gorm:"unique"`
	DishName          string    `json:"dish_name"`
	RestaurantName    string    `json:"restaurant_name"`
	TransactionAmount float64   `json:"transaction_amount"`
	Time              time.Time `json:"time"`
}

type Restaurant struct {
	ID             uint    `json:"id" gorm:"primary_key:autoIncrement"`
	RestaurantName string  `json:"restaurant_name" gorm:"unique"`
	CashBalance    float64 `json:"cash_balance"`
}

type Menu struct {
	ID       uint    `json:"id" gorm:"primary_key:autoIncrement"`
	DishID   string  `json:"dish-id" gorm:"unique"`
	DishName string  `json:"dish_name"`
	Price    float64 `json:"price"`
}

type OpeningHours struct {
	ID           uint      `json:"id" gorm:"primary_key:autoIncrement"`
	RestaurantID string    `json:"restaurant_id" gorm:"unique"`
	Day          int       `json:"day"`
	OpeningTime  time.Time `json:"opening_time"`
	ClosingTime  time.Time `json:"closing_time"`
}
