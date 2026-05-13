// internal/repo/user_repo.go
// 数据访问层：负责与数据库进行交互，执行 CRUD 操作
package repo

import (
	"word-book/internal/model"

	"gorm.io/gorm"
)

// UserRepo 负责用户相关的数据访问操作
type UserRepo struct {
	DB *gorm.DB
}

// 创建一个新的 UserRepo 实例
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// 创建用户
func (r *UserRepo) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

// 根据用户名查询用户
func (r *UserRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
