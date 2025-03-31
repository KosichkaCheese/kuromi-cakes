package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"kuromi_cakes/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open("host=db user=kuromi password=kuromi dbname=kuromi-cake port=5432 sslmode=disable TimeZone=UTC"), &gorm.Config{})
	// DB, err = gorm.Open(postgres.Open("host=localhost user=kuromi password=kuromi dbname=kuromi-cake port=5438 sslmode=disable TimeZone=UTC"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database")

	DB.AutoMigrate(&models.Cake{})
}
