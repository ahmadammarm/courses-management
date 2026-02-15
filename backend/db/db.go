package db

import (
	"fmt"
	"github.com/ahmadammarm/courses-management/backend/config"
	"github.com/ahmadammarm/courses-management/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Database *gorm.DB

func DatabaseConnect() {
	dbUser := config.GetEnv("DB_USER")
	dbPassword := config.GetEnv("DB_PASSWORD")
	dbHost := config.GetEnv("DB_HOST")
	dbPort := config.GetEnv("DB_PORT")
	dbName := config.GetEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the database schema in development environment
	if config.GetEnv("APP_ENV") == "development" {
		err = Database.AutoMigrate(&models.Instructor{}, &models.Course{})
		if err != nil {
			log.Fatalf("Failed to auto-migrate database: %v", err)
		}
	}

	log.Println("Database connection established")
}
