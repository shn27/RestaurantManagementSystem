package database

import "log"

func deleteTables() {
	err := DB.Migrator().DropTable(
		&User{},
		&Restaurant{},
		&PurchaseHistory{},
		&OpeningHours{},
		&Menu{},
	)
	if err != nil {
		log.Fatalf("Failed to delete databases: %v", err)
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
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
