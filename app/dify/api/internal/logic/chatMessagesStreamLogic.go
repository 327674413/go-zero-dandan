package logic

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-zero-dandan/app/dify/api/internal/svc"
	"go-zero-dandan/app/dify/api/internal/types"
	"go-zero-dandan/common/fmtd"
	"go-zero-dandan/common/resd"
	"io"
	"net/http"
	"strings"
	"time"
)

type ChatMessagesStreamLogic struct {
	*ChatMessagesStreamLogicGen
}

func NewChatMessagesStreamLogic(ctx context.Context, svc *svc.ServiceContext) *ChatMessagesStreamLogic {
	return &ChatMessagesStreamLogic{
		ChatMessagesStreamLogicGen: NewChatMessagesStreamLogicGen(ctx, svc),
	}
}
func (l *ChatMessagesStreamLogic) ChatMessagesStream(w http.ResponseWriter, in *types.ChatMessagesReq) error {
	var err error
	if err = l.initReq(in); err != nil {
		return l.resd.Error(err)
	}
	if !l.hasUserInfo {
		return resd.NewErr("未登录")
	}
	data := map[string]any{
		"query":         l.req.Query,
		"inputs":        map[string]any{},
		"user":          l.meta.UserId,
		"response_mode": "streaming",
	}
	err = l.difyChatMessagesSteam(data, w)
	if err != nil {
		return l.resd.Error(err)
	}
	return nil
}

func (l *ChatMessagesStreamLogic) difyChatMessagesSteam(reqData map[string]any, w http.ResponseWriter) (err error) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	// API 地址
	url := l.svc.Config.Dify.Url + "/chat-messages"
	// 将请求数据编码为 JSON 格式
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return l.resd.Error(err, resd.ErrJsonEncode)
	}
	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return l.resd.Error(err, resd.ErrCurlCreate)
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.svc.Config.Dify.AppKey)
	// 发送请求
	client := &http.Client{
		Timeout: 0,
	}
	resp, err := client.Do(req)
	if err != nil {
		return l.resd.Error(err, resd.ErrCurlSend)
	}
	defer resp.Body.Close()
	// 判断响应类型
	contentType := resp.Header.Get("Content-Type")
	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return l.resd.NewErr(resd.ErrTrdNotFound)
		} else {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return l.resd.Error(fmt.Errorf("%d", resp.StatusCode), resd.ErrIoRead)
			}
			if contentType == "application/json" {
				errResp := make(map[string]any)
				json.Unmarshal(body, &errResp)
				return l.resd.Error(errors.New(string(body)), resd.ErrTrdDifyChatStream)
			} else {
				// 如果是普通字符串
				return l.resd.Error(errors.New(string(body)), resd.ErrTrdDifyChatStream)
			}
		}
	}
	if strings.Contains(contentType, "text/event-stream") {
		// 使用 bufio 逐行读取 SSE 数据流
		scanner := bufio.NewScanner(resp.Body)
		flusher, ok := w.(http.Flusher)
		if !ok {
			return l.resd.NewErr(resd.ErrCurlSteamNotSupported)
		}
		for scanner.Scan() {
			llmText := scanner.Text()
			fmtd.Info(llmText)
			if llmText == "" {
				fmt.Fprint(w, "\n")
			} else {
				fmt.Fprint(w, llmText+"\n") // 写入数据格式：data: <数据内容>\n\n
			}
			//手动刷新，及时发送
			flusher.Flush()
			//目前测试好像会粘包，前端一次性收到多条，加个间隔
			time.Sleep(50 * time.Millisecond)

			//后面部分是自己后端请求时使用，直接返还给前端就不额外处理了
			// 忽略注释行
			if strings.HasPrefix(llmText, ":") {
				continue
			}
			// 解析 "event" 和 "data" 字段
			//if strings.HasPrefix(llmText, "event:") {
			//	event := strings.TrimSpace(strings.TrimPrefix(llmText, "event:"))
			//	fmt.Println("Event:", event)
			//} else if strings.HasPrefix(llmText, "data:") {
			//	data := strings.TrimSpace(strings.TrimPrefix(llmText, "data:"))
			//
			//	// 尝试将 data 解析为 JSON（如果格式为 JSON）
			//	var jsonData map[string]interface{}
			//	if err := json.Unmarshal([]byte(data), &jsonData); err == nil {
			//		fmt.Println("JSON Data:", jsonData)
			//	} else {
			//		fmt.Println("Data:", data)
			//	}
			//}
		}

		if err = scanner.Err(); err != nil {
			return l.resd.Error(err, resd.ErrCurlSteamScan)
		}
	} else {
		return l.resd.Error(errors.New("返回的非流式数据"), resd.ErrTrdDifyChatStream)
	}
	return nil
}
