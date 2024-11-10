package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-suriya/go-fiber-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstace *postgresDatabase
	once                    sync.Once
)

func NewPostgresDatabase(config config.Config) Database {
	once.Do(func() {
		dsn := config.GetDSN()

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("Connected to database %s", "go_fiber")

		postgresDatabaseInstace = &postgresDatabase{conn}
	})

	return postgresDatabaseInstace
}

func (db *postgresDatabase) Connect() *gorm.DB {
	return postgresDatabaseInstace.DB
}

func (db *postgresDatabase) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("error getting underlying sql.DB: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("error closing database connection: %v", err)
	}

	log.Printf("Closed connection to database %s", "go_fiber")
	return nil
}
