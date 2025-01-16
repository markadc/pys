package pys

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

func Get(url string, headers, params S) (*Response, error) {
	// 构造请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 查询字符串
	values := req.URL.Query()
	for k, v := range params {
		values.Add(k, v)
	}
	req.URL.RawQuery = values.Encode()

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

func GET(url string, headers, params S) (*Response, error) {
	// 构造请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 查询字符串
	values := req.URL.Query()
	for k, v := range params {
		values.Add(k, v)
	}
	req.URL.RawQuery = values.Encode()

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

func GetByConfigs(srcURL string, reqConf *RequestConfigs) (*Response, error) {
	req, err := http.NewRequest("GET", srcURL, nil)
	if err != nil {
		return nil, err
	}

	if params := reqConf.Params; params != nil {
		values := req.URL.Query()
		for k, v := range params {
			values.Set(k, v)
		}
		req.URL.RawQuery = values.Encode()
	}

	if headers := reqConf.Header; headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	var proxyURL *url.URL
	if proxy := reqConf.Proxy; proxy != "" {
		proxyURL, err = url.Parse(proxy)
		if err != nil {
			return nil, err
		}
	}

	timeout := 10 * time.Second
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Request: req, Text: string(body), StatusCode: resp.StatusCode}, nil
}
