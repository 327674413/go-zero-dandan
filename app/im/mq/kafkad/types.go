package kafkad

import "go-zero-dandan/app/im/ws/websocketd"

// MsgChatTransfer 消息发送
type MsgChatTransfer struct {
	ConversationId      string `json:"conversationId"`
	websocketd.ChatType `json:"chatType"`
	SendId              int64   `json:"sendId,string"`
	RecvId              int64   `json:"recvId,string"`
	RecvIds             []int64 `json:"recvIds,string"`
	SendTime            int64   `json:"sendTime"`
	websocketd.MsgType  `json:"msgType"`
	Content             string `json:"content"`
	PlatId              int64  `json:"platId,string"`
}

// MsgMarkRead 消息读取状态
type MsgMarkRead struct {
	ConversationId      string `json:"conversationId"`
	websocketd.ChatType `json:"chatType"`
	SendId              int64    `json:"sendId,string"`
	RecvId              int64    `json:"recvId,string"`
	MsgIds              []string `json:"msgIds"`
}
