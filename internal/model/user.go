// model/user.go
// 存放用户相关的结构体定义
package model

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"` // 不返回密码字段
	Email    string `gorm:"unique;not null" json:"email"`

	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt int64 `gorm:"index" json:"-"` // 软删除字段，不返回
}

func (User) TableName() string {
	return "users"
}
