package database

import (
	"fmt"
)

func deleteTables() {
	err := DB.Migrator().DropTable(
		&User{},
		&Restaurant{},
		&PurchaseHistory{},
		&OpeningHours{},
		&Menu{},
	)
	if err != nil {
		fmt.Println("Failed to delete databases: %v", err)
	}
}

func migrateDatabase() {
	err := DB.AutoMigrate(
		&User{},
		&Restaurant{},
		&PurchaseHistory{},
		&OpeningHours{},
		&Menu{},
	)
	if err != nil {
		fmt.Println("Failed to migrate database: %v", err)
	}
}
