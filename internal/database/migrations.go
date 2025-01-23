package database

import "log"

func MigrateDatabase() {
	err := DB.AutoMigrate(
		&User{},
		&Restaurant{},
		&PurchaseHistory{},
		&OpeningHours{},
		&Menu{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
