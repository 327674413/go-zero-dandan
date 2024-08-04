package modelMongo

import (
	"go-zero-dandan/app/im/ws/websocketd"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationId string              `bson:"conversationId,omitempty"` //会话id
	ChatType       websocketd.ChatType `bson:"chatType,omitempty"`       //聊天类型
	TargetId       string              `bson:"targetId,omitempty"`       //好友/群组id
	IsShow         bool                `bson:"isShow,omitempty"`         //是否显示
	Total          int64               `bson:"total,omitempty"`          //消息总数
	Seq            int64               `bson:"seq"`                      //消息读取节点序号
	LastMsg        *ChatLog            `bson:"msg"`                      //最后一条消息
	LastAt         int64               `bson:"lastAt"`                   //最后一条时间
	PlatId         string              `bson:"platId"`
	UpdateAt       time.Time           `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt       time.Time           `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
