package main

import (
	"UrlShortener/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Define PostgreSQL connection string
	dsn := "host=localhost user=postgres password= dbname=postgres port=5432 sslmode=disable"

	// Open connection using GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL using GORM")

	// Run Migrations
	err = DB.AutoMigrate(&models.URL{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate: %v", err)
	}
	fmt.Println("✅ Migration completed")
}

func CloseDB() {
	dbInstance, err := DB.DB()
	if err != nil {
		log.Fatalf("❌ Error getting DB instance: %v", err)
	}
	dbInstance.Close()
}
