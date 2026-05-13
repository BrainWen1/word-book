// internal/utils/response.go
// 统一响应结构体和函数
package utils

import (
	"github.com/gin-gonic/gin"
)

// 统一返回状态码常量
const (
	CodeSuccess = 200 // 成功
	CodeFail    = 400 // 业务失败
	CodeAuth    = 401 // 未登录/令牌失效
	CodeServer  = 500 // 服务器内部错误
)

// Response 定义统一的响应结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 消息
	Data    interface{} `json:"data"`    // 数据
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(CodeSuccess, Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

func SuccessMessageResponse(c *gin.Context, message string) {
	c.JSON(CodeSuccess, Response{
		Code:    CodeSuccess,
		Message: message,
		Data:    nil,
	})
}

func FailResponse(c *gin.Context, data interface{}) {
	c.JSON(CodeFail, Response{
		Code:    CodeFail,
		Message: "fail",
		Data:    data,
	})
}

func FailAuthResponse(c *gin.Context, message string) {
	c.JSON(CodeAuth, Response{
		Code:    CodeAuth,
		Message: message,
		Data:    nil,
	})
}

func FailServerResponse(c *gin.Context, message string) {
	c.JSON(CodeServer, Response{
		Code:    CodeServer,
		Message: message,
		Data:    nil,
	})
}
