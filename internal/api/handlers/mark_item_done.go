package handlers

import (
	"net/http"
	"todo/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MarkItemDone(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectName := c.Param("project_name")
		itemName := c.Param("item_name")
		var req struct {
			Method string `json:"method"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		status := "done"
		if req.Method == "undone" {
			status = "undone"
		}
		var item models.Item
		if err := db.Where("project_name = ? AND name = ?", projectName, itemName).First(&item).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "项目项不存在"})
			return
		}
		if err := db.Model(&item).Update("status", status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新项目项状态"})
			return
		}
		c.JSON(http.StatusOK, item)
	}
}
