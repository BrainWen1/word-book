// internal/router/router.go
// 该文件负责设置和管理路由
package router

import (
	"word-book/internal/handler/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建默认的Gin引擎
	r := gin.Default()

	// 注册中间件
	r.Use(middleware.Cors()) // 跨域中间件

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
