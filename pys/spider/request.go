package spider

import (
	"io"
	"net/http"
)

func Get(url string, headers, params S) (*Response, error) {
	// 构造请求
	req, err := http.NewRequest("GET", MakeURL(url, params), nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// 读取内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &Response{Request: req, Text: string(body), StatusCode: resp.StatusCode}, nil
}
