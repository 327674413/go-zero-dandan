package modelMongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ChatLogModel = (*customChatLogModel)(nil)

type (
	// ChatLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatLogModel.
	ChatLogModel interface {
		chatLogModel
		ListBySendTime(ctx context.Context, conversationId string, startSendTime, endSendTime, limit int64) ([]*ChatLog, error)
		ListByMsgIds(ctx context.Context, msgIds []string) ([]*ChatLog, error)
		UpdateMakeRead(ctx context.Context, id primitive.ObjectID, readRecords []byte) error
	}

	customChatLogModel struct {
		*defaultChatLogModel
	}
)

// NewChatLogModel returns a model for the mongo.
func NewChatLogModel(url, db, collection string) ChatLogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customChatLogModel{
		defaultChatLogModel: newDefaultChatLogModel(conn),
	}
}

func MustChatLogModel(url, db string) ChatLogModel {
	return NewChatLogModel(url, db, "chat_log")
}

var DefaultChatLogLimit int64 = 100

func (m *defaultChatLogModel) ListByMsgIds(ctx context.Context, msgIds []string) ([]*ChatLog, error) {
	var data []*ChatLog
	ids := make([]primitive.ObjectID, 0, len(msgIds))
	for _, id := range msgIds {
		oid, _ := primitive.ObjectIDFromHex(id)
		ids = append(ids, oid)
	}
	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	err := m.conn.Find(ctx, &data, filter)
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultChatLogModel) ListBySendTime(ctx context.Context, conversationId string, startSendTime, endSendTime, limit int64) ([]*ChatLog, error) {
	var data []*ChatLog
	opt := options.FindOptions{Limit: &DefaultChatLogLimit, Sort: bson.M{"sendTime": -1}}
	if limit > 0 {
		opt.Limit = &limit
	}
	filter := bson.M{
		"conversationId": conversationId,
	}
	if endSendTime > 0 {
		filter["sendTime"] = bson.M{
			"$gt":  endSendTime,
			"$lte": startSendTime,
		}
	} else {
		filter["sendTime"] = bson.M{
			"$lt": startSendTime,
		}
	}
	err := m.conn.Find(ctx, &data, filter, &opt)
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultChatLogModel) UpdateMakeRead(ctx context.Context, id primitive.ObjectID, readRecords []byte) error {
	_, err := m.conn.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"readRecords": readRecords}})
	return err
}