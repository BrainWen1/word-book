package main

import (
	"log"
	"word-book/internal/config"
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

	// 设置路由
	r := router.SetupRouter()

	log.Println("Server has been started successfully.")

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
