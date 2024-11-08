package main

import (
	"github.com/go-suriya/go-fiber-api/database"
	"github.com/go-suriya/go-fiber-api/entities"
	"gorm.io/gorm"
)

func main() {
	db := database.NewPostgresDatabase()

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
