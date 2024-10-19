package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAddProjectItem(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Project{}, &models.Item{})
	assert.NoError(t, err)

	// 创建测试项目
	testProject := models.Project{Name: "Test Project", UserID: 1}
	db.Create(&testProject)

	r := gin.Default()
	r.POST("/project/:project_name/item", AddProjectItem(db))

	newItem := models.Item{
		Name:       "New Item",
		ParentTime: 10,
	}
	jsonValue, _ := json.Marshal(newItem)

	req, _ := http.NewRequest("POST", "/project/Test Project/item", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	// 模拟认证中间件设置用户ID
	r.Use(func(c *gin.Context) {
		c.Set("user_id", uint(1))
		c.Next()
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Item
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "New Item", response.Name)
	assert.Equal(t, float64(10), response.ParentTime)
	assert.Equal(t, testProject.ID, response.ProjectID)
}
