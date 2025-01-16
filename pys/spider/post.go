package spider

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Post(URL string, headers, params S, data any) (*Response, error) {
	// 请求体
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 构造请求
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(jsonBytes))

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
