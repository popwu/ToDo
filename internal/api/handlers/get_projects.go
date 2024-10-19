package handlers

import (
	"net/http"
	"todo/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProjects(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		var projects []models.Project
		if err := db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取项目列表"})
			return
		}
		c.JSON(http.StatusOK, projects)
	}
}
