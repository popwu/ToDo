package database

import (
	"todo/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 数据库迁移
	err = db.AutoMigrate(&models.Project{}, &models.Item{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
