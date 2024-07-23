package modelMongo

import (
	"go-zero-dandan/app/im/ws/websocketd"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SysMsgStat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId    string             `bson:"userId"`
	MsgClasEm websocketd.MsgClas `bson:"msgClasEm"`
	UnreadNum int64              `bson:"unreadNum"`
	UpdateAt  time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
