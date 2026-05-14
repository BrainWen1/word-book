// internal/infra/external/search_word.go
// 存放调用外部词典API的结构体和函数
package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"word-book/internal/config"
)

// DictResponse 定义从外部词典API返回的数据结构
type DictResponse []struct { // API返回的是一个切片，包含一个单词的多个含义
	Word      string `json:"word"` // 单词
	Phonetics []struct {
		Text      string `json:"text"`      // 音标文本
		Audio     string `json:"audio"`     // 发音音频URL
		SourceUrl string `json:"sourceUrl"` // 来源URL
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string     `json:"partOfSpeech"` // 词性
		Definitions  []struct { // 定义
			Definition string   `json:"definition"`         // 定义文本
			Synonyms   []string `json:"synonyms,omitempty"` // 同义词（可选）
			Antonyms   []string `json:"antonyms,omitempty"` // 反义词（可选）
			Example    string   `json:"example,omitempty"`  // 例句（可选）
		} `json:"definitions"`
		Synonyms []string `json:"synonyms,omitempty"` // 词义的同义词（可选）
		Antonyms []string `json:"antonyms,omitempty"` // 词义的反义词（可选）
	} `json:"meanings"`
}

func SearchWord(word string) (DictResponse, error) {
	// 拼接url
	apiURL := fmt.Sprintf("%s/%s", config.AppConfig.Dict_api, url.PathEscape(word))
	client := &http.Client{Timeout: 10 * time.Second} // 设置HTTP客户端的超时时间，避免请求挂起过久
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 注册关闭响应体的延迟调用

	// 错误处理
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("单词 '%s' 未找到", word)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	// 解析响应
	var dictResponse DictResponse
	if err := json.NewDecoder(resp.Body).Decode(&dictResponse); err != nil {
		return nil, err
	}

	return dictResponse, nil
}
