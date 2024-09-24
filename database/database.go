package database

import (
	"os"

	"github.com/ananascharles/binify/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	if _, err := os.Stat("database/test.db"); os.IsNotExist(err) {
		file, err := os.Create("database/test.db")
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open("database/test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Paste{})
	if err != nil {
		return err
	}

	return nil
}
