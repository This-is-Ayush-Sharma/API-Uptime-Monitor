package database

import (
	"fmt"
	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		//		use default
		dsn = "host=localhost user=postgres password=password dbname=uptime_monitor port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	DB = db
	fmt.Println("Connected to database & Migrated.")
	return db
}
