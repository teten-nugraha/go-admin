package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {

	dsn := "host=localhost user=postgres password=postgres dbname=go_admin port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

}
