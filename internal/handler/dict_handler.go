// internal/handler/dict_handler.go
// 调用词典api的处理器的http处理器
package handler

import (
	"word-book/internal/service"
	"word-book/internal/utils/response"

	"github.com/gin-gonic/gin"
)

type DictHandler struct {
	DictService *service.DictService
}

func NewDictHandler(dictService *service.DictService) *DictHandler {
	return &DictHandler{DictService: dictService}
}

// SearchWord 查词
// @Summary 查询单词
// @Description 根据单词查询释义、音标、例句
// @Tags 词典模块
// @Accept json
// @Produce json
// @Param word query string true "要查询的单词"
// @Success 200 {object} response.Response{data=[]external.DictResponse}
// @Failure 404 {object} response.Response "单词不存在"
// @Router /search [get]
func (h *DictHandler) SearchWord(c *gin.Context) {
	// 从查询参数中获取单词
	word := c.Query("word") // 使用查询参数: http://localhost:8080/api/v1/search?word=hello
	if word == "" {
		response.FailResponse(c, "请输入要查询的单词")
		return
	}

	// 调用服务层的SearchWord方法
	result, err := h.DictService.SearchWord(word)
	if err != nil {
		response.FailResponse(c, gin.H{
			"message": "查询单词失败",
			"error":   err.Error(),
		})
		return
	}
	response.SuccessResponse(c, result)
}
