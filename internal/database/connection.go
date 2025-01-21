package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var Connection = &cobra.Command{
	Use:   "db",
	Short: "Database connection",
	Long:  "Database connection",
	Run: func(cmd *cobra.Command, args []string) {
		connection()
	},
}

var DB *gorm.DB

func connection() {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	for {
		var err error
		DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			log.Println("Failed to connect to database: ", err)
			log.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			log.Println("Successfully connected to database")
			fmt.Println(DB)
			break
		}
	}
	// Auto-migrate the User table
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create a new user
	newUser := User{UserName: "John Doe", CashBalance: 15}
	result := DB.Create(&newUser)
	if result.Error != nil {
		log.Fatalf("Failed to create user: %v", result.Error)
	}
	log.Println("User created:", newUser)

	newUser = User{UserName: "Sohan", CashBalance: 1015}
	result = DB.Create(&newUser)
	if result.Error != nil {
		log.Fatalf("Failed to create user: %v", result.Error)
	}
	log.Println("User created:", newUser)

	// Retrieve users
	var users []User
	DB.Find(&users)
	log.Println("Users:", users)
}
