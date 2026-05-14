// internal/handler/middleware/auth.go
// 认证中间件
package middleware

import (
	"strings"
	"word-book/internal/utils/jwt"
	"word-book/internal/utils/response"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 拿到请求头中的Authorization字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.FailAuthResponse(c, "缺少Authorization头")
			c.Abort()
			return
		}

		// 拿到token字符串，通常格式是 "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.FailAuthResponse(c, "Authorization头格式错误")
			c.Abort()
			return
		}

		// 解析和验证token
		token := parts[1]
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.FailAuthResponse(c, "无效的token")
			c.Abort()
			return
		}

		// 将用户信息存储在上下文中，供后续处理函数使用
		c.Set("userID", claims.UserID)
		c.Set("username", claims.UserName)

		c.Next()
	}
}
