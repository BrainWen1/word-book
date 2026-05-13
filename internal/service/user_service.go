// internal/service/user_service.go
// 存放用户相关的业务逻辑函数，如注册、登录、获取用户信息等
package service

import (
	"errors"
	"word-book/internal/model"
	"word-book/internal/repo"

	"golang.org/x/crypto/bcrypt"
)

// UserService 负责用户相关的业务逻辑
type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

// 用户注册的业务逻辑
func (s *UserService) Register(username, password, email string) (*model.User, error) {
	// 查找是否已经存在该用户
	exist, err := s.UserRepo.FindByUsername(username)
	if err == nil && exist != nil {
		return nil, errors.New("用户名已被占用")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建新用户
	user := &model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Email:        email,
	}

	// 将用户保存到数据库
	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, errors.New("用户创建失败")
	}

	return user, nil
}
