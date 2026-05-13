// internal/handler/user_handler.go
// 存放用户相关的HTTP处理函数，如注册、登录、获取用户信息等
package handler

import (
	"word-book/internal/service"
	"word-book/internal/utils"

	"github.com/gin-gonic/gin"
)

// UserHandler 负责处理用户相关的HTTP请求
type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// RegisterRequest 定义用户注册请求的结构体
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func (h *UserHandler) Register(c *gin.Context) {
	// 解析请求体
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailResponse(c, gin.H{
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 调用 UserService 进行注册
	user, err := h.UserService.Register(req.Username, req.Password, req.Email)
	if err != nil {
		utils.FailResponse(c, gin.H{
			"message": "注册失败",
			"error":   err.Error(),
		})
		return
	}

	utils.SuccessResponse(c, gin.H{
		"message": "注册成功",
		"user_id": user.ID,
	})
}
