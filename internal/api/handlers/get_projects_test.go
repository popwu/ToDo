package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/internal/middleware"
	"todo/internal/models"

	"todo/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 模拟登录并获取 token 的辅助函数
func loginAndGetToken(t *testing.T, r *gin.Engine, db *gorm.DB) string {
	// 创建测试用户
	password := "testpassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	testUser := models.User{
		Username: "testuser",
		Password: string(hashedPassword),
	}
	result := db.Create(&testUser)
	assert.NoError(t, result.Error)

	// 构造登录请求
	loginData := map[string]string{
		"username": "testuser",
		"password": password,
	}
	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// 执行登录请求
	r.ServeHTTP(w, req)

	// 打印返回状态和内容
	// t.Logf("返回状态: %d", w.Code)
	// t.Logf("返回内容: %s", w.Body.String())

	// 检查响应
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	token, exists := response["token"]
	assert.True(t, exists, "响应中应包含 token")

	return token
}

func TestGetProjects(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.User{}, &models.Project{})
	assert.NoError(t, err)

	// 创建一个配置对象
	cfg := &config.Config{
		JWTSecret: "your-test-jwt-secret",
	}

	r := gin.Default()
	r.POST("/login", Login(db, cfg.JWTSecret))
	r.GET("/projects", middleware.AuthMiddleware(cfg.JWTSecret), GetProjects(db))

	// 获取登录 token
	token := loginAndGetToken(t, r, db)

	// 创建测试数据
	testUser := models.User{}
	db.Where("username = ?", "testuser").First(&testUser)
	testProjects := []models.Project{
		{Name: "Project 1", UserID: testUser.ID},
		{Name: "Project 2", UserID: testUser.ID},
		{Name: "Project 3", UserID: testUser.ID + 1},
	}
	for _, project := range testProjects {
		result := db.Create(&project)
		assert.NoError(t, result.Error)
	}

	// 使用获取到的 token 创建请求
	req, _ := http.NewRequest("GET", "/projects", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var projects []models.Project
	err = json.Unmarshal(w.Body.Bytes(), &projects)
	assert.NoError(t, err)
	assert.Len(t, projects, 2)

	// 添加长度检查
	if len(projects) > 0 {
		assert.Equal(t, "Project 1", projects[0].Name)
		assert.Equal(t, "Project 2", projects[1].Name)
	} else {
		t.Error("没有返回项目")
	}
}
