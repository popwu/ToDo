package handlers

import (
	"net/http"
	"todo/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteProjectItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectName := c.Param("project_name")
		itemName := c.Param("item_name")
		var item models.Item
		if err := db.Where("project_name = ? AND name = ?", projectName, itemName).First(&item).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "项目项不存在"})
			return
		}
		if err := db.Delete(&item).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法删除项目项"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "项目项已成功删除"})
	}
}
