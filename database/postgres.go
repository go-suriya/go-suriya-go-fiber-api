package database

import (
	"log"
	"time"

	"github.com/avast/retry-go/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=168 dbname=go_fiber port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	err := retry.Do(
		func() error {
			var dbErr error
			DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if dbErr != nil {
				log.Println("Database connection failed. Retrying...")
			}
			return dbErr
		},
		retry.Attempts(5),
		retry.Delay(2*time.Second),
		retry.DelayType(retry.BackOffDelay),
	)

	if err != nil {
		log.Fatalf("Could not connect to the database after 5 attempts: %v", err)
	}

	log.Println("Database connection established")
}

func CloseDB() {
	DBConn, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Close the database connection
	if err := DBConn.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	} else {
		log.Println("Database connection closed")
	}
}
