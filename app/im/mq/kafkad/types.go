package kafkad

import "go-zero-dandan/app/im/ws/websocketd"

type MsgChatTransfer struct {
	ConversationId      string `json:"conversationId"`
	websocketd.ChatType `json:"chatType"`
	SendId              int64 `json:"sendId,string"`
	RecvId              int64 `json:"recvId,string"`
	SendTime            int64 `json:"sendTime"`
	websocketd.MsgType  `json:"msgType"`
	Content             string `json:"content"`
}
