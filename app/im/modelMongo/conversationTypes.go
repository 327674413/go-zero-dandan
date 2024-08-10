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
	UserId         string              `bson:"userId,omitempty"`         //归属人
	TargetId       string              `bson:"targetId,omitempty"`       //对象好友id/群主id
	IsShow         bool                `bson:"isShow,omitempty"`         //列表专属，是否显示
	Total          int64               `bson:"total,omitempty"`          //列表专属，消息总数
	ReadSeq        int64               `bson:"readSeq"`                  //列表专属，消息读取节点序号，读到50条就是50，后面就可以从51条开始读
	DeleteSeq      int64               `bson:"deleteSeq"`                //列表专属，删除的节点序号，比如50，那就是50以前的消息的会话都被删除了
	LastMsg        *ChatLog            `bson:"lastMsg"`                  //最后一条消息
	LastAt         int64               `bson:"lastAt"`                   //最后一条时间
	PlatId         string              `bson:"platId"`
	UpdateAt       time.Time           `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt       time.Time           `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
