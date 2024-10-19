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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	return db, err
}

func TestRegister(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	r := gin.Default()
	r.POST("/register", Register(db, "test_secret"))

	user := models.User{Username: "testuser", Password: "testpassword"}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "token")
}

func TestLogin(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	r := gin.Default()
	r.POST("/login", Login(db, "test_secret"))

	// 首先注册一个用户
	password := "testpassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	user := models.User{Username: "testuser", Password: string(hashedPassword)}
	db.Create(&user)

	loginInfo := map[string]string{"username": "testuser", "password": "testpassword"}
	jsonValue, _ := json.Marshal(loginInfo)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "token")
}
