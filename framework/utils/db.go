package utils

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return nil
}
