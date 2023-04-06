package utils

import (
	"log"
	"os"

	"github.com/caiovillardo/crawler/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.Ticket{})

	return db

}
