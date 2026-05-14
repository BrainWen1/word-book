// internal/router/router.go
// 该文件负责设置和管理路由
package router

import (
	"word-book/internal/handler"
	"word-book/internal/handler/middleware"
	"word-book/internal/infra/database"
	"word-book/internal/repo"
	"word-book/internal/service"
	"word-book/internal/utils/response"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建默认的Gin引擎
	r := gin.Default()

	// 注册全局中间件
	r.Use(middleware.Cors()) // 跨域中间件

	// 初始化各层组件
	userRepo := repo.NewUserRepo(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	dictService := service.NewDictService()
	dictHandler := handler.NewDictHandler(dictService)

	// 公共路由
	api := r.Group("/api/v1") // http://localhost:8080/api/v1
	{
		// 健康检查接口
		api.GET("/ping", func(c *gin.Context) {
			response.SuccessResponse(c, "pong")
		})
		// 用户注册
		api.POST("/register", userHandler.Register)
		// 用户登录
		api.POST("/login", userHandler.Login)
		// 查询单词
		api.GET("/search", dictHandler.SearchWord)
	}

	// 受保护的路由
	auth := api.Group("/")      // http://localhost:8080/api/v1/
	auth.Use(middleware.Auth()) // 认证中间件
	{
		auth.GET("/profile", func(c *gin.Context) {
			userID := c.GetInt("userID")
			username := c.GetString("username")
			response.SuccessResponse(c, gin.H{
				"user_id":  userID,
				"username": username,
			})
		})
	}

	return r
}
