package main

import (
	"log"

	"todo/internal/api"
	"todo/internal/config"
	"todo/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("无法加载配置: %v", err)
	}

	// 初始化数据库
	db, err := database.Init(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}

	// 创建Gin引擎
	r := gin.Default()

	// 设置路由
	api.SetupRoutes(r, db, cfg.JWTSecret)

	// 启动服务器
	log.Printf("服务器正在运行于 %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}
