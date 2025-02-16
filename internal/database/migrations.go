package database

import "log"

func DeleteTables() {
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
