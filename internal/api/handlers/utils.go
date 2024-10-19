package handlers

import (
	"todo/internal/models"

	"gorm.io/gorm"
)

func getProjectID(db *gorm.DB, projectName string) uint {
	var project models.Project
	db.Where("name = ?", projectName).First(&project)
	return project.ID
}
