package api

import (
	"todo/internal/api/handlers"
	"todo/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, jwtSecret string) {
	// 添加 CORS 中间件，允许所有域
	r.Use(cors.Default())

	api := r.Group("/api")

	// 认证路由
	auth := api.Group("/user")
	{
		auth.POST("/register", handlers.Register(db, jwtSecret))
		auth.POST("/login", handlers.Login(db, jwtSecret))
	}

	// 受保护的路由
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// 项目路由
		projects := protected.Group("/projects")
		{
			projects.GET("", handlers.GetProjects(db))
			projects.GET("/:project_name/items", handlers.GetProjectItems(db))
			projects.POST("/:project_name/item", handlers.AddProjectItem(db))
			projects.PATCH("/:project_name/item/:item_name", handlers.UpdateProjectItem(db))
			projects.PATCH("/:project_name/item/:item_name/done", handlers.MarkItemDone(db))
			projects.DELETE("/:project_name/item/:item_name", handlers.DeleteProjectItem(db))
		}
	}
}
