package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("dev.ddbb"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
