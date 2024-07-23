package modelMongo

import (
	"go-zero-dandan/app/im/ws/websocketd"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ChatLog struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	ConversationId string              `bson:"conversationId"`
	SendId         string              `bson:"sendId"`
	RecvId         string              `bson:"recvId"`
	MsgFrom        int                 `bson:"msgFrom"` //不知道啥用还
	ChatType       websocketd.ChatType `bson:"chatType"`
	MsgType        websocketd.MsgType  `bson:"msgType"`
	MsgContent     string              `bson:"msgContent"`
	SendTime       string              `bson:"sendTime"`
	State          int                 `bson:"state"`
	ReadRecords    []byte              `bson:"readRecords"`
	PlatId         string              `bson:"platId"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
