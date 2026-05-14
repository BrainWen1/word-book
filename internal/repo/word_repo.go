// internal/repo/word_repo.go
package repo

import (
	"word-book/internal/model"

	"gorm.io/gorm"
)

type WordRepo struct {
	DB *gorm.DB
}

func NewWordRepo(db *gorm.DB) *WordRepo {
	return &WordRepo{DB: db}
}

// Create 创建一个新的单词记录
func (r *WordRepo) Create(word *model.Word) error {
	return r.DB.Create(word).Error
}

// FindByUserAndWord 根据用户ID和单词查询单词记录
func (r *WordRepo) FindByUserAndWord(UserID int, word string) (*model.Word, error) {
	var result model.Word
	if err := r.DB.Where("user_id = ? AND word = ?", UserID, word).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// FindByID 根据ID查询单词记录
func (r *WordRepo) FindByID(userID, wordID int) (*model.Word, error) {
	var result model.Word
	if err := r.DB.Where("id = ? AND user_id = ?", wordID, userID).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserWords 根据用户ID查询所有单词记录
func (r *WordRepo) GetUserWords(userID int) ([]model.Word, error) {
	var result []model.Word
	if err := r.DB.Where("user_id = ?", userID).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// ListByUser 分页查询，支持按掌握度筛选
func (r *WordRepo) ListByUser(userID int, mastery *model.MasteryLevel, page, pageSize int) ([]model.Word, int64, error) {
	var words []model.Word
	var total int64

	query := r.DB.Model(&model.Word{}).Where("user_id = ?", userID) // 先过滤用户ID
	if mastery != nil {
		query = query.Where("mastery = ?", *mastery) // 如果提供了掌握度参数，则继续过滤掌握度
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize // 跳过(page - 1)页
	if err := query.Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&words).Error; err != nil {
		return nil, 0, err
	}
	return words, total, nil
}

// Update 更新单词
func (r *WordRepo) Update(word *model.Word) error {
	return r.DB.Save(word).Error
}

// Delete 删除
func (r *WordRepo) Delete(userID, wordID int) error {
	return r.DB.Where("id = ? AND user_id = ?", wordID, userID).Delete(&model.Word{}).Error
}
