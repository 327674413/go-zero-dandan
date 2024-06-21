package kafkad

import "go-zero-dandan/app/im/ws/websocketd"

// MsgChatTransfer 消息发送
type MsgChatTransfer struct {
	ConversationId      string `json:"conversationId"`
	websocketd.ChatType `json:"chatType"`
	SendId              string   `json:"sendId"`
	RecvId              string   `json:"recvId"`
	RecvIds             []string `json:"recvIds"`
	SendTime            int64    `json:"sendTime"`
	websocketd.MsgType  `json:"msgType"`
	Content             string `json:"content"`
	PlatId              string `json:"platId"`
	MsgId               string `json:"msgId"`
}

// MsgMarkRead 消息读取状态
type MsgMarkRead struct {
	ConversationId      string `json:"conversationId"`
	websocketd.ChatType `json:"chatType"`
	SendId              string   `json:"sendId"`
	RecvId              string   `json:"recvId"`
	MsgIds              []string `json:"msgIds"`
}
