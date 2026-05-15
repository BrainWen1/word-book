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
	DB_dsn   string // 数据库源
	Dict_api string // 词典API地址

	JWTSecretKey []byte // JWT密钥

	Redis_addr     string // Redis地址
	Redis_password string // Redis密码
	Redis_db       int    // Redis数据库索引
}

// AppConfig 全局配置变量
var AppConfig Config

// LoadConfig 从配置文件或环境变量加载配置
func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found, falling back to system environment variables")
	}

	AppConfig = Config{
		DB_dsn:         os.Getenv("DB_DSN"),
		Dict_api:       os.Getenv("DICT_API"),
		JWTSecretKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		Redis_addr:     os.Getenv("REDIS_ADDR"),
		Redis_password: os.Getenv("REDIS_PASSWORD"),
		Redis_db:       0, // 默认使用0号数据库
	}
}
