package handlers

import (
	"net/http"
	"todo/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddProjectItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectName := c.Param("project_name")
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item.ProjectID = getProjectID(db, projectName)
		if err := db.Create(&item).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法添加项目项"})
			return
		}
		c.JSON(http.StatusCreated, item)
	}
}
