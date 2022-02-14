package db

import (
	"github.com/teten-nugraha/go-admin/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=postgres password=postgres dbname=go_admin port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = db

	// migrations from model/*
	db.AutoMigrate(&model.User{})
}
