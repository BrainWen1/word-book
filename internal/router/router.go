// internal/router/router.go
// 该文件负责设置和管理路由
package router

import (
	"word-book/internal/handler"
	"word-book/internal/handler/middleware"
	"word-book/internal/infra/database"
	"word-book/internal/repo"
	"word-book/internal/service"
	"word-book/internal/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建默认的Gin引擎
	r := gin.Default()

	// 初始化各层组件
	userRepo := repo.NewUserRepo(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 注册中间件
	r.Use(middleware.Cors()) // 跨域中间件

	// 公共路由
	api := r.Group("/api/v1")
	{
		// 健康检查接口
		api.GET("/ping", func(c *gin.Context) {
			utils.SuccessResponse(c, "pong")
		})
		// 用户注册
		api.POST("/register", userHandler.Register)
	}

	return r
}
