package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	DB = db
	return DB, err
}
