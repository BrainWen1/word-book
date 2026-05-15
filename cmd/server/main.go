package main

// @title           单词本 API
// @version         1.0
// @description     基于 Go + Gin + JWT + GORM + Redis 的在线单词本项目
// @termsOfService  http://swagger.io/terms/

// @contact.name   BrainWen
// @contact.url    https://github.com/BrainWen1
// @contact.email  brianclarkxx@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization

import (
	"log"
	"word-book/internal/config"
	"word-book/internal/infra/cache"
	"word-book/internal/infra/database"
	"word-book/internal/router"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 连接数据库
	database.ConnectDB()

	// 自动迁移
	database.MigrateDB()

	// 初始化Redis缓存
	cache := cache.NewRedisCache(
		config.AppConfig.Redis_addr,
		config.AppConfig.Redis_password,
		config.AppConfig.Redis_db,
	)

	// 设置路由
	r := router.SetupRouter(cache)

	log.Println("Server has been started successfully.")

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
