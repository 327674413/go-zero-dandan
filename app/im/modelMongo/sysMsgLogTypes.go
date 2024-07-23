package modelMongo

import (
	"go-zero-dandan/app/im/ws/websocketd"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SysMsgLog struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MsgClasEm  websocketd.MsgClas `bson:"msgClasEm"`
	RecvId     string             `bson:"recvId"`
	MsgContent string             `bson:"msgContent"`
	IsRead     int64              `bson:"isRead"`
	UpdateAt   time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt   time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
