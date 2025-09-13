package database

import (
	"fmt"
	//"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/models"
)

var DB *gorm.DB

func connect() {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		//		use default
		dsn = "host=localhost user=postgres password=postgres dbname=uptime_monitor port=5432 sslmode=disable"
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
}
