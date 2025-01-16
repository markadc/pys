package spider

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response 响应对象
type Response struct {
	StatusCode int
	Text       string
	Request    *http.Request
}

// JSON 获取响应的JSON格式数据
func (r Response) JSON() (map[string]any, error) {
	var jsonData map[string]any
	err := json.Unmarshal([]byte(r.Text), &jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// JsonStringify JSON字符串格式化，优化打印显示
func (r Response) JsonStringify() string {
	jsonData, err := r.JSON()
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("JSON序列化失败: %v", err)
	}
	return string(b)
}
