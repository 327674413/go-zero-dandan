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
	MsgId          string              `bson:"msgId"` //不清楚mongo，自己的雪花id放在这里
	SendAtMs       int64               `bson:"sendAtMs"`
	ChatType       websocketd.ChatType `bson:"chatType"`
	MsgType        websocketd.MsgType  `bson:"msgType"`
	MsgContent     string              `bson:"msgContent"`
	SendTime       string              `bson:"sendTime"`
	MsgState       int64               `bson:"state"`
	//MsgReads        []byte              `bson:"msgReads"` 位图方式没弄好，暂时先不用
	ReadUsers   map[string]int32 `bson:"readUsers"` //在mongo里只有32 和 64
	ReadTotalNo int64            `bson:"readTotalNo"`
	TempId      string           `bson:"tempId"` //前端生成的临时id
	PlatId      string           `bson:"platId"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
