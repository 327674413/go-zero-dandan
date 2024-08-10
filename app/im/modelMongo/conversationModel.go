package modelMongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"go.mongodb.org/mongo-driver/bson"
)

var _ ConversationModel = (*customConversationModel)(nil)

type (
	// ConversationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConversationModel.
	ConversationModel interface {
		conversationModel
		UpdateMsg(ctx context.Context, chatLog *ChatLog) error
		ListByConversationIds(ctx context.Context, ids []string) ([]*Conversation, error)
		FindByConvId(ctx context.Context, id string) (*Conversation, error)
	}

	customConversationModel struct {
		*defaultConversationModel
	}
)

// NewConversationModel returns a model for the mongo.
func NewConversationModel(url, db, collection string) ConversationModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customConversationModel{
		defaultConversationModel: newDefaultConversationModel(conn),
	}
}
func (m *defaultConversationModel) FindByConvId(ctx context.Context, id string) (*Conversation, error) {
	var data Conversation

	err := m.conn.FindOne(ctx, &data, bson.M{"conversationId": id})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, nil
	default:
		return nil, resd.ErrorCtx(ctx, err, resd.ErrMongoSelect)
	}
}

func (m *defaultConversationModel) ListByConversationIds(ctx context.Context, ids []string) ([]*Conversation, error) {
	var data []*Conversation

	err := m.conn.Find(ctx, &data, bson.M{
		"conversationId": bson.M{
			"$in": ids,
		},
	})
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, nil
	default:
		return nil, resd.ErrorCtx(ctx, err, resd.ErrMongoSelect)
	}
}

func (m *defaultConversationModel) UpdateMsg(ctx context.Context, chatLog *ChatLog) error {
	_, err := m.conn.UpdateOne(ctx,
		bson.M{"conversationId": chatLog.ConversationId},
		bson.M{
			// 更新会话总消息数
			"$inc": bson.M{"total": 1},
			"$set": bson.M{"lastMsg": chatLog, "lastAt": utild.GetStamp()},
		},
	)
	if err != nil {
		return resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return nil
}
func MustConversationModel(url, db string) ConversationModel {
	return NewConversationModel(url, db, "conversation")
}
