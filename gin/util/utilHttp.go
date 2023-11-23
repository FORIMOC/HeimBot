package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// SendJSON 发送HTTP POST的JSON请求
// 请求url, 请求json数据 => 请求数据的 body 部分字节流
func SendJSON(url string, jsonData map[string]interface{}) []byte {
	// 编码json
	requestBody, _ := json.Marshal(jsonData)

	// 创建HTTP请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return body
}
