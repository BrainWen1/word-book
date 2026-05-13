// internal/infra/database/mysql.go
// GORM初始化和数据库连接配置
package database

import (
	"log"
	"word-book/internal/config"
	"word-book/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局数据库连接变量
var DB *gorm.DB

// 连接数据库
func ConnectDB() {
	db, err := gorm.Open(mysql.Open(config.AppConfig.DB_dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用物理外键约束
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	DB = db
}

// 自动迁移
func MigrateDB() {
	DB.AutoMigrate(&model.User{}, &model.Word{})
}
