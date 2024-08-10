// 不要用goctl生成覆盖
package modelMongo

import (
	"context"
	"go-zero-dandan/common/resd"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type conversationModel interface {
	Insert(ctx context.Context, data *Conversation) error
	FindOne(ctx context.Context, id string) (*Conversation, error)
	Update(ctx context.Context, data *Conversation) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultConversationModel struct {
	conn *mon.Model
}

func newDefaultConversationModel(conn *mon.Model) *defaultConversationModel {
	return &defaultConversationModel{conn: conn}
}

func (m *defaultConversationModel) Insert(ctx context.Context, data *Conversation) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultConversationModel) FindOne(ctx context.Context, id string) (*Conversation, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Conversation

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultConversationModel) Update(ctx context.Context, data *Conversation) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultConversationModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoIdHex)
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
