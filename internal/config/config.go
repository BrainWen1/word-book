// config/config.go
// 存放数据库dsn,词典API地址等配置信息
// 以及初始化数据库连接和自动迁移的函数
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 存储应用程序的配置信息
type Config struct {
	DB_dsn       string // 数据库源
	Dict_api     string // 词典API地址
	JWTSecretKey []byte // JWT密钥
}

// AppConfig 全局配置变量
var AppConfig Config

// LoadConfig 从环境变量加载配置
func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found, falling back to system environment variables")
	}

	AppConfig = Config{
		DB_dsn:       os.Getenv("DB_DSN"),
		Dict_api:     os.Getenv("DICT_API"),
		JWTSecretKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}
}
