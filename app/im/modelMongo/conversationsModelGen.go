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

type conversationsModel interface {
	Insert(ctx context.Context, data *Conversations) error
	FindOne(ctx context.Context, id string) (*Conversations, error)
	Update(ctx context.Context, data *Conversations) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultConversationsModel struct {
	conn *mon.Model
}

func newDefaultConversationsModel(conn *mon.Model) *defaultConversationsModel {
	return &defaultConversationsModel{conn: conn}
}

func (m *defaultConversationsModel) Insert(ctx context.Context, data *Conversations) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	if err != nil {
		return resd.ErrorCtx(ctx, err, resd.ErrMongoInsert)
	}
	return nil
}

func (m *defaultConversationsModel) FindOne(ctx context.Context, id string) (*Conversations, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, resd.ErrorCtx(ctx, err, resd.ErrMongoIdHex)
	}
	var data Conversations

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, resd.ErrorCtx(ctx, err, resd.ErrMongoSelect)
	}
}

func (m *defaultConversationsModel) Update(ctx context.Context, data *Conversations) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	if err != nil {
		return nil, resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return res, err
}

func (m *defaultConversationsModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoIdHex)
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoDelete)
	}
	return res, err
}
