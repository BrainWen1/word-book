// model/word.go
// 存放单词相关的结构体定义
package model

import (
	"time"

	"gorm.io/gorm"
)

// MasteryLevel 表示单词掌握程度的枚举类型
type MasteryLevel int

const (
	Unknowned MasteryLevel = iota // 0: 未知
	Fuzzy                         // 1: 模糊
	Mastered                      // 2: 掌握
)

type Word struct {
	ID     int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int `gorm:"not null" json:"user_id"` // 外键关联用户

	Word       string `gorm:"not null" json:"word"` // 不能设置为unique，因为一个单词可能是多个用户的生词
	Phonetic   string `json:"phonetic,omitempty"`   // 音标
	Definition string `json:"definition,omitempty"` // 释义
	Example    string `json:"example,omitempty"`    // 例句

	Mastery        MasteryLevel `gorm:"not null;default:0" json:"mastery"` // 掌握程度
	LastReviewedAt int64        `gorm:"default:0" json:"last_reviewed_at"` // 上次复习时间
	ReviewCount    int          `gorm:"default:0" json:"review_count"`     // 复习次数

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除
}

func (Word) TableName() string {
	return "words"
}
