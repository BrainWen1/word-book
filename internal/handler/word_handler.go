// internal/service/word_handler.go
// 存放生词相关的业务逻辑函数，如添加生词、更新掌握程度、获取生词列表等
package handler

import (
	"strconv"
	"word-book/internal/model"
	"word-book/internal/service"
	"word-book/internal/utils/response"

	"github.com/gin-gonic/gin"
)

type WordHandler struct {
	WordService *service.WordService
}

func NewWordHandler(wordService *service.WordService) *WordHandler {
	return &WordHandler{WordService: wordService}
}

// AddWordRequest 定义添加单词的请求体结构
type AddWordRequest struct {
	Word       string `json:"word" binding:"required"`
	Phonetic   string `json:"phonetic,omitempty"`
	Definition string `json:"definition,omitempty"`
	Example    string `json:"example,omitempty"`
}

// AddWord 添加一个新的生词
func (h *WordHandler) AddWord(c *gin.Context) {
	// 解析请求体
	var req AddWordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailMsgAndDataResponse(c, "请求参数错误", err.Error())
		return
	}

	userID := c.GetInt("userID") // 从上下文获取用户ID

	newWord, err := h.WordService.AddWord(userID, req.Word, req.Phonetic, req.Definition, req.Example)
	if err != nil {
		response.FailMsgAndDataResponse(c, "单词添加失败", err.Error())
		return
	}

	response.SuccessMsgAndDataResponse(c, "单词添加成功", newWord)
}

// ListWords 获取单词列表，支持分页和按掌握度筛选
func (h *WordHandler) ListWords(c *gin.Context) {
	// 从查询参数获取分页和筛选信息
	userID := c.GetInt("userID")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	masteryStr := c.Query("mastery")
	var mastery *model.MasteryLevel
	if masteryStr != "" {
		m, err := strconv.Atoi(masteryStr)
		if err == nil && m >= 0 && m <= 2 {
			level := model.MasteryLevel(m)
			mastery = &level
		}
	}

	// 调用WordService获取单词列表
	words, total, err := h.WordService.ListWords(userID, mastery, page, pageSize)
	if err != nil {
		response.FailMsgAndDataResponse(c, "获取单词列表失败", err.Error())
		return
	}
	// 回显
	response.SuccessMsgAndDataResponse(c, "获取单词列表成功", gin.H{
		"items":     words,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// UpdateMasteryRequest 定义更新掌握度的请求体结构
type UpdateMasteryRequest struct {
	Mastery model.MasteryLevel `json:"mastery" binding:"required,oneof=0 1 2"`
}

func (h *WordHandler) UpdateMastery(c *gin.Context) {
	userID := c.GetInt("userID")
	wordIDStr := c.Param("id")
	wordID, err := strconv.ParseInt(wordIDStr, 10, 64)
	if err != nil {
		response.FailMsgAndDataResponse(c, "无效的单词ID", err.Error())
		return
	}

	var req UpdateMasteryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailMsgAndDataResponse(c, "请提供正确的掌握度 (0,1,2)", err.Error())
		return
	}

	if err := h.WordService.UpdateMastery(userID, (int)(wordID), req.Mastery); err != nil {
		response.FailMsgAndDataResponse(c, "更新掌握度失败", err.Error())
		return
	}
	response.SuccessMsgAndDataResponse(c, "更新成功", gin.H{
		"user_id":     userID,
		"word_id":     wordID,
		"new_mastery": req.Mastery,
	})
}

func (h *WordHandler) DeleteWord(c *gin.Context) {
	userID := c.GetInt("userID")
	wordIDStr := c.Param("id")
	wordID, err := strconv.ParseInt(wordIDStr, 10, 64)
	if err != nil {
		response.FailMsgAndDataResponse(c, "无效的单词ID", err.Error())
		return
	}

	if err := h.WordService.DeleteWord(userID, (int)(wordID)); err != nil {
		response.FailMsgAndDataResponse(c, "删除失败", err.Error())
		return
	}
	response.SuccessMsgAndDataResponse(c, "删除成功", gin.H{
		"user_id": userID,
		"word_id": wordID,
	})
}
