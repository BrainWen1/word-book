// internal/handler/user_handler.go
// 存放用户相关的HTTP处理函数，如注册、登录、获取用户信息等
package handler

import (
	"word-book/internal/service"
	"word-book/internal/utils/response"

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

// Register 用户注册
// @Summary 用户注册
// @Description 注册新用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param register body RegisterRequest true "注册参数"
// @Success 200 {object} response.Response "注册成功"
// @Failure 400 {object} response.Response "参数错误"
// @Router /register [post]
func (h *UserHandler) Register(c *gin.Context) {
	// 解析请求体
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailResponse(c, gin.H{
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 调用 UserService 进行注册
	user, err := h.UserService.Register(req.Username, req.Password, req.Email)
	if err != nil {
		response.FailResponse(c, gin.H{
			"message": "注册失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回成功响应
	response.SuccessResponse(c, gin.H{
		"message": "注册成功",
		"user_id": user.ID,
	})
}

// LoginRequest 定义用户登录请求的结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 用户登录
// @Summary 用户登录
// @Description 登录获取 JWT Token
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param login body LoginRequest true "登录参数"
// @Success 200 {object} response.Response{data=response.Response} "登录成功"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 401 {object} response.Response "用户名或密码错误"
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	// 解析请求体
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailResponse(c, gin.H{
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 调用 UserService 进行登录验证
	token, err := h.UserService.Login(req.Username, req.Password)
	if err != nil {
		response.FailResponse(c, gin.H{
			"message": "登录失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回成功响应
	response.SuccessResponse(c, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}
