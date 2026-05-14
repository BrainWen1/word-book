// internal/service/word_service.go
// 存放生词相关的业务逻辑函数，如添加生词、更新掌握程度、获取生词列表等
package service

import (
	"errors"
	"fmt"
	"word-book/internal/model"
	"word-book/internal/repo"

	"gorm.io/gorm"
)

type WordService struct {
	WordRepo *repo.WordRepo
}

func NewWordService(wordRepo *repo.WordRepo) *WordService {
	return &WordService{WordRepo: wordRepo}
}

// AddWord 添加一个新的生词
func (s *WordService) AddWord(userID int, word, phonetic, definition, example string) (*model.Word, error) {
	// 先检查这个单词是否已经存在
	_, err := s.WordRepo.FindByUserAndWord(userID, word)
	if err != nil {
		// 当错误不是“记录未找到”时，返回错误
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("查询单词失败: %w", err)
		}
		// 如果是“记录未找到”，说明单词不存在，继续执行
	} else {
		// 没有错误，说明查到了数据，单词已存在
		return nil, errors.New("你已经添加过这个单词了")
	}

	// 创建新的单词记录并插入
	newWord := &model.Word{
		UserID:     userID,
		Word:       word,
		Phonetic:   phonetic,
		Definition: definition,
		Example:    example,
		Mastery:    model.Unknowned, // 默认掌握度为未知
	}
	if err := s.WordRepo.Create(newWord); err != nil {
		return nil, err
	}

	return newWord, nil
}

// ListWords 分页列表
func (s *WordService) ListWords(userID int, mastery *model.MasteryLevel, page, pageSize int) ([]model.Word, int64, error) {
	// 将来可以从缓存读取，目前直接查库
	return s.WordRepo.ListByUser(userID, mastery, page, pageSize)
}

// UpdateMastery 修改掌握度
func (s *WordService) UpdateMastery(userID, wordID int, mastery model.MasteryLevel) error {
	newWord, err := s.WordRepo.FindByID(userID, wordID)
	if err != nil {
		return fmt.Errorf("单词不存在")
	}
	newWord.Mastery = mastery
	return s.WordRepo.Update(newWord)
}

// DeleteWord 删除单词
func (s *WordService) DeleteWord(userID int, wordID int) error {
	existing, err := s.WordRepo.FindByID(userID, wordID)
	if err != nil {
		return fmt.Errorf("单词不存在")
	}
	return s.WordRepo.Delete(userID, existing.ID)
}
