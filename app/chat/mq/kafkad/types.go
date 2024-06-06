package kafkad

import "go-zero-dandan/app/chat/ws/websocketd"

type MsgChatTransfer struct {
	ConversationCode    string `json:"conversationCode"`
	websocketd.ChatType `json:"chatType"`
	SendId              int64 `json:"sendId,string"`
	RecvId              int64 `json:"recvId,string"`
	SendTime            int64 `json:"sendTime"`
	websocketd.MsgType  `json:"msgType"`
	Content             string `json:"content"`
}
