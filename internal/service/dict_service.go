// internal/service/dict_service.go
// 存放调用词典API的业务逻辑，包括redis缓存
package service

import "word-book/internal/infra/external"

type DictService struct {
}

func NewDictService() *DictService {
	return &DictService{}
}

func (s *DictService) SearchWord(word string) (external.DictResponse, error) {
	// TODO: 添加redis缓存逻辑，先从缓存中获取，如果没有再调用外部API，并将结果存入缓存
	return external.SearchWord(word)
}
