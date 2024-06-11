package modelMongo

import (
	"go-zero-dandan/app/im/ws/websocketd"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ChatLog struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	ConversationId string              `bson:"conversationId"`
	SendId         int64               `bson:"sendId"`
	RecvId         int64               `bson:"recvId"`
	MsgFrom        int                 `bson:"msgFrom"`
	ChatType       websocketd.ChatType `bson:"chatType"`
	MsgType        websocketd.MsgType  `bson:"msgType"`
	MsgContent     string              `bson:"msgContent"`
	SendTime       int64               `bson:"sendTime"`
	State          int                 `bson:"state"`
	ReadRecords    []byte              `bson:"readRecords"`
	PlatId         int64               `bson:"platId"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
