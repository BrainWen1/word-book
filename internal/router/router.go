// internal/router/router.go
// 该文件负责设置和管理路由
package router

import (
	"net/http"
	"word-book/internal/handler"
	"word-book/internal/handler/middleware"
	"word-book/internal/infra/database"
	"word-book/internal/repo"
	"word-book/internal/service"
	"word-book/internal/utils/response"
	"word-book/internal/webapp"

	_ "word-book/docs" // Swagger文档包

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	wordRepo := repo.NewWordRepo(database.DB)
	wordService := service.NewWordService(wordRepo)
	wordHandler := handler.NewWordHandler(wordService)

	// 页面与公共路由
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", webapp.IndexHTML)
	})
	// Swagger 文档
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
	auth := api.Group("/user")  // http://localhost:8080/api/v1/user
	auth.Use(middleware.Auth()) // 认证中间件
	{
		// 健康检查接口（认证后）
		auth.GET("/ping", func(c *gin.Context) {
			userID := c.GetInt("userID")
			username := c.GetString("username")
			response.SuccessResponse(c, gin.H{
				"user_id":  userID,
				"username": username,
			})
		})
		// 获取单词列表
		auth.GET("/words", wordHandler.ListWords)
		// 添加单词
		auth.POST("/words", wordHandler.AddWord)
		// 更新掌握度
		auth.PUT("/words/:id", wordHandler.UpdateMastery)
		// 删除单词
		auth.DELETE("/words/:id", wordHandler.DeleteWord)
	}

	return r
}
