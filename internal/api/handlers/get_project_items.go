package handlers

import (
	"net/http"
	"strings"
	"todo/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TreeNode struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	ParentTime float64     `json:"parent_time"`
	Children   []*TreeNode `json:"children,omitempty"`
}

func buildTree(items []models.Item) []*TreeNode {
	itemMap := make(map[string]*TreeNode)
	var roots []*TreeNode

	for _, item := range items {
		parts := strings.Split(item.Name, "/")
		currentPath := ""
		var currentNode *TreeNode

		for i, part := range parts {
			if i > 0 {
				currentPath += "/"
			}
			currentPath += part

			if node, exists := itemMap[currentPath]; exists {
				currentNode = node
			} else {
				newNode := &TreeNode{
					ID:         currentPath,
					Name:       part,
					ParentTime: item.ParentTime,
				}
				itemMap[currentPath] = newNode

				if currentNode == nil {
					roots = append(roots, newNode)
				} else {
					currentNode.Children = append(currentNode.Children, newNode)
				}
				currentNode = newNode
			}
		}
	}

	return roots
}

func GetProjectItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectName := c.Param("project_name")
		var items []models.Item
		if err := db.Where("project_name = ?", projectName).Find(&items).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取项目项目"})
			return
		}
		treeItems := buildTree(items)
		c.JSON(http.StatusOK, treeItems)
	}
}
