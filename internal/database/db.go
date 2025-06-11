package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dns := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}

	return DB, nil
}
