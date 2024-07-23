package kafkad

import (
	"go-zero-dandan/app/im/ws/websocketd"
)

// MsgChatTransfer 消息发送 (聊天消息、提醒消息都用这个了)
type MsgChatTransfer struct {
	ConversationId      string `json:"conversationId"`
	websocketd.ChatType `json:"chatType"`
	SendId              string   `json:"sendId"`
	RecvId              string   `json:"recvId"`
	RecvIds             []string `json:"recvIds"`
	SendTime            string   `json:"sendTime"`
	websocketd.MsgType  `json:"msgType"`
	Content             string `json:"content"`
	PlatId              string `json:"platId"`
	MsgId               string `json:"msgId"`
	websocketd.MsgClas  `json:"msgClas"`
}

// MsgMarkRead 消息读取状态
type MsgMarkRead struct {
	ConversationId      string `json:"conversationId"`
	websocketd.ChatType `json:"chatType"`
	SendId              string   `json:"sendId"`
	RecvId              string   `json:"recvId"`
	MsgIds              []string `json:"msgIds"`
	PlatId              string   `json:"platId"`
}

// SysToUserMsg 系统消息
type SysToUserMsg struct {
	websocketd.MsgClas `json:"msgClas"`
	RecvId             string `json:"recvId"`
	websocketd.MsgType `json:"msgType"`
	MsgContent         string `json:"msgContent"`
	SendTime           string `json:"sendTime"`
	PlatId             string `json:"platId"`
}
