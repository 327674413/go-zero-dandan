package modelMongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ ConversationsModel = (*customConversationsModel)(nil)

type (
	// ConversationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConversationsModel.
	ConversationsModel interface {
		conversationsModel
		FindByUserId(ctx context.Context, userId int64) (*Conversations, error)
		Save(ctx context.Context, data *Conversations) error
	}

	customConversationsModel struct {
		*defaultConversationsModel
	}
)

// NewConversationsModel returns a model for the mongo.
func NewConversationsModel(url, db, collection string) ConversationsModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customConversationsModel{
		defaultConversationsModel: newDefaultConversationsModel(conn),
	}
}

func MustConversationsModel(url, db string) ConversationsModel {
	return NewConversationsModel(url, db, "conversations")
}
func (m *defaultConversationsModel) Save(ctx context.Context, data *Conversations) error {
	data.UpdateAt = time.Now()
	_, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data}, options.Update().SetUpsert(true))
	return err
}
func (m *defaultConversationsModel) FindByUserId(ctx context.Context, userId int64) (*Conversations, error) {
	var data Conversations

	err := m.conn.FindOne(ctx, &data, bson.M{"userId": userId})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
