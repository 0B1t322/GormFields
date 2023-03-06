package gorm

import (
	"github.com/0B1t322/GormFields/example/gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToSqlite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"))
	if err != nil {
		return nil, err
	}

	if err := createModels(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createModels(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Model{}); err != nil {
		return err
	}

	return nil
}
