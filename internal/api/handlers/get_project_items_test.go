package handlers

import (
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

func TestGetProjectItems(t *testing.T) {
	// 设置测试数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Item{})
	assert.NoError(t, err)

	// 创建测试数据
	testItems := []models.Item{
		{Name: "item1", ProjectName: "TestProject", UserID: 1, ParentTime: 10},
		{Name: "item1/subitem1", ProjectName: "TestProject", UserID: 1, ParentTime: 5},
		{Name: "item1/subitem1/subsubitem1", ProjectName: "TestProject", UserID: 1, ParentTime: 2},
		{Name: "item2", ProjectName: "TestProject", UserID: 1, ParentTime: 15},
		{Name: "item3", ProjectName: "TestProject", UserID: 1, ParentTime: 20},
		{Name: "item3/subitem3", ProjectName: "TestProject", UserID: 1, ParentTime: 8},
	}
	for _, item := range testItems {
		result := db.Create(&item)
		assert.NoError(t, result.Error)
	}

	// 设置 Gin 路由
	r := gin.Default()
	r.GET("/project/:project_name/items", GetProjectItems(db))

	// 创建测试请求
	req, _ := http.NewRequest("GET", "/project/TestProject/items", nil)
	w := httptest.NewRecorder()

	// 执行请求
	r.ServeHTTP(w, req)

	// 检查响应状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 解析响应体
	var response []*TreeNode
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 验证树状结构
	assert.Len(t, response, 3) // 顶层应该有3个项目

	// 验证第一个项目及其子项目
	assert.Equal(t, "item1", response[0].ID)
	assert.Equal(t, "item1", response[0].Name)
	assert.Equal(t, float64(10), response[0].ParentTime)
	assert.Len(t, response[0].Children, 1)

	subitem := response[0].Children[0]
	assert.Equal(t, "item1/subitem1", subitem.ID)
	assert.Equal(t, "subitem1", subitem.Name)
	assert.Equal(t, float64(5), subitem.ParentTime)
	assert.Len(t, subitem.Children, 1)

	subsubitem := subitem.Children[0]
	assert.Equal(t, "item1/subitem1/subsubitem1", subsubitem.ID)
	assert.Equal(t, "subsubitem1", subsubitem.Name)
	assert.Equal(t, float64(2), subsubitem.ParentTime)
	assert.Len(t, subsubitem.Children, 0)

	// 验证第二个项目
	assert.Equal(t, "item2", response[1].ID)
	assert.Equal(t, "item2", response[1].Name)
	assert.Equal(t, float64(15), response[1].ParentTime)
	assert.Len(t, response[1].Children, 0)

	// 验证第三个项目及其子项目
	assert.Equal(t, "item3", response[2].ID)
	assert.Equal(t, "item3", response[2].Name)
	assert.Equal(t, float64(20), response[2].ParentTime)
	assert.Len(t, response[2].Children, 1)

	subitem3 := response[2].Children[0]
	assert.Equal(t, "item3/subitem3", subitem3.ID)
	assert.Equal(t, "subitem3", subitem3.Name)
	assert.Equal(t, float64(8), subitem3.ParentTime)
	assert.Len(t, subitem3.Children, 0)
}
