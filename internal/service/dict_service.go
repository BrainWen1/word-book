// internal/service/dict_service.go
// 存放调用词典API的业务逻辑，包括redis缓存
package service

import (
	"context"
	"encoding/json"
	"time"
	"word-book/internal/infra/cache"
	"word-book/internal/infra/external"
)

type DictService struct {
	cache cache.Cache
}

func NewDictService(cache cache.Cache) *DictService {
	return &DictService{cache: cache}
}

func (s *DictService) SearchWord(word string) (external.DictResponse, error) {
	// 尝试从缓存中获取
	ctx := context.Background()
	key := "dict:" + word
	cached, err := s.cache.Get(ctx, key)
	if err == nil && cached != "" {
		// 从缓存中解析并返回
		var resp external.DictResponse
		err = json.Unmarshal([]byte(cached), &resp)
		if err == nil {
			return resp, nil
		}
	}

	// 调用外部api
	resp, err := external.SearchWord(word)
	if err != nil {
		return nil, err
	}

	// 写入缓存
	data, _ := json.Marshal(resp)
	s.cache.Set(ctx, key, string(data), 24*time.Hour) // 过期时间为24小时

	return resp, nil
}
