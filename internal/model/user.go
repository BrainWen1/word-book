// model/user.go
// 存放用户相关的结构体定义
package model

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string `gorm:"type:varchar(100);unique;not null" json:"username"`
	PasswordHash string `gorm:"type:varchar(255);not null" json:"password_hash"` // 存储密码哈希值，不存储明文密码
	Email        string `gorm:"type:varchar(100)" json:"email"`

	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt int64 `gorm:"index" json:"-"` // 软删除字段，不返回
}

func (User) TableName() string {
	return "users"
}
