package main

import (
	"log"

	"github.com/go-suriya/go-fiber-api/config"
	"github.com/go-suriya/go-fiber-api/database"
	"github.com/go-suriya/go-fiber-api/entities"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Migration Failed to load configuration: %v", err)
	}

	db := database.NewPostgresDatabase(*config)

	tx := db.Connect().Begin()

	categoriesMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func categoriesMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Category{})
}
