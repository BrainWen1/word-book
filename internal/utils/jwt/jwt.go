// internal/utils/jwt/jwt.go
// 该文件负责JWT相关的功能，如生成和验证JWT令牌
package jwt

import (
	"errors"
	"time"
	"word-book/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

// 配置项
var TokenExpireDuration = 24 * 7 * time.Hour // 默认7天过期

func jwtSecretKey() []byte {
	return config.AppConfig.JWTSecretKey
}

type CustomClaims struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID int, username string) (string, error) {
	// 创建自定义声明
	claims := CustomClaims{
		UserID:   userID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			Issuer:    "word-book",                                             // 发行者
		},
	}

	// 创建JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用HS256算法签名

	// 返回签名后的令牌字符串，使用全局配置的JWT密钥进行签名
	return token.SignedString(jwtSecretKey())
}

// ParseToken 解析和验证JWT令牌
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey(), nil // 使用配置中的JWT密钥进行验证
	})
	if err != nil {
		return nil, err
	}

	// 验证令牌有效性
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
