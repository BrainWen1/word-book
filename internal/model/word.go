// model/word.go
// 存放单词相关的结构体定义
package model

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

	Word           string       `gorm:"unique;not null" json:"word"`
	Mastery        MasteryLevel `gorm:"not null" json:"mastery"`                // 掌握程度
	LastReviewedAt int64        `gorm:"autoUpdateTime" json:"last_reviewed_at"` // 上次复习时间
	ReviewCount    int          `gorm:"default:0" json:"review_count"`          // 复习次数

	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt int64 `gorm:"index" json:"-"` // 软删除
}

func (Word) TableName() string {
	return "words"
}
