package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	curl()
}

func curl() {
	// API 地址
	url := "http://192.168.0.98:8070/v1/chat-messages"
	// 请求数据
	requestData := map[string]interface{}{
		"query":         "你好",
		"inputs":        map[string]interface{}{},
		"response_mode": "streaming",
		"user":          "user_id_1",
	}
	// 将请求数据编码为 JSON 格式
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("JSON 编码错误:", err)
		return
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("创建请求错误:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer app-NPbrmyq5ACrvCPEHJQiRVuM5")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求发送错误:", err)
		return
	}
	defer resp.Body.Close()
	// 判断响应类型
	contentType := resp.Header.Get("Content-Type")
	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		if contentType == "application/json" {
			var result map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				fmt.Println("Error decoding JSON response:", err)
				return
			}
			fmt.Println("JSON response:", result)
			return
		} else {
			// 如果是普通字符串
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}
			fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode, "数据：\n", string(body))
			return
		}
	}

	// 如果是 JSON 格式
	if contentType == "application/json" {
		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Println("Error decoding JSON response:", err)
			return
		}
		fmt.Println("JSON response:", result)
		return
	} else if contentType == "application/octet-stream" || contentType == "application/pdf" {
		// 如果是文件，假设响应头包含文件名
		fileName := resp.Header.Get("Content-Disposition")
		if fileName == "" {
			fileName = "downloaded_file"
		}
		out, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println("Error saving file:", err)
			return
		}
		fmt.Println("File saved:", fileName)
		return
	} else if strings.Contains(contentType, "text/event-stream") {
		// 使用 bufio 逐行读取 SSE 数据流
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			// 忽略注释行
			if strings.HasPrefix(line, ":") {
				continue
			}
			// 解析 "event" 和 "data" 字段
			if strings.HasPrefix(line, "event:") {
				event := strings.TrimSpace(strings.TrimPrefix(line, "event:"))
				fmt.Println("Event:", event)
			} else if strings.HasPrefix(line, "data:") {
				data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))

				// 尝试将 data 解析为 JSON（如果格式为 JSON）
				var jsonData map[string]interface{}
				if err := json.Unmarshal([]byte(data), &jsonData); err == nil {
					fmt.Println("JSON Data:", jsonData)
				} else {
					fmt.Println("Data:", data)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading SSE stream:", err)
		}
	} else {
		// 如果是普通字符串
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		fmt.Println("String response:", string(body))
		return
	}

}
