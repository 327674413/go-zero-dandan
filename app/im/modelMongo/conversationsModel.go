package modelMongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go-zero-dandan/common/resd"
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
		FindByUserId(ctx context.Context, userId string) (*Conversations, error)
		Save(ctx context.Context, data *Conversations) error
		UpdateMsg(ctx context.Context, userId string, data *ChatLog) error
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
	if err != nil {
		return resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return err
}

func (m *defaultConversationsModel) UpdateMsg(ctx context.Context, userId string, chatLog *ChatLog) error {
	filter := bson.M{"userId": userId, "conversationList." + chatLog.ConversationId: bson.M{"$exists": true}}

	update := bson.M{
		"$set": bson.M{
			"conversationList." + chatLog.ConversationId + ".lastMsg": chatLog,
			"conversationList." + chatLog.ConversationId + ".lastAt":  chatLog.SendAtMs,
		},
		"$inc": bson.M{
			"conversationList." + chatLog.ConversationId + ".total":   1,
			"conversationList." + chatLog.ConversationId + ".readSeq": 1,
		},
	}
	_, err := m.conn.UpdateOne(ctx, filter, update)
	if err != nil {
		return resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}

	return nil
}
func (m *defaultConversationsModel) FindByUserId(ctx context.Context, userId string) (*Conversations, error) {
	var data Conversations

	err := m.conn.FindOne(ctx, &data, bson.M{"userId": userId})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, nil
	default:
		return nil, resd.ErrorCtx(ctx, err, resd.ErrMongoSelect)
	}
}
