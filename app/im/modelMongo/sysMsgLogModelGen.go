// Code generated by goctl. DO NOT EDIT.
package modelMongo

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type sysMsgLogModel interface {
	Insert(ctx context.Context, data *SysMsgLog) error
	FindOne(ctx context.Context, id string) (*SysMsgLog, error)
	Update(ctx context.Context, data *SysMsgLog) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
	SetSysMsgReadByMsgClas(ctx context.Context, userId string, msgClasEm []int64) (updateNum int64, danErr error)
	SetSysMsgReadByIds(ctx context.Context, userId string, msgIds []string, msgClasEm int64) (updateNum int64, danErr error)
}

type defaultSysMsgLogModel struct {
	conn *mon.Model
}

func newDefaultSysMsgLogModel(conn *mon.Model) *defaultSysMsgLogModel {
	return &defaultSysMsgLogModel{conn: conn}
}

func (m *defaultSysMsgLogModel) Insert(ctx context.Context, data *SysMsgLog) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultSysMsgLogModel) FindOne(ctx context.Context, id string) (*SysMsgLog, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data SysMsgLog

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysMsgLogModel) Update(ctx context.Context, data *SysMsgLog) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultSysMsgLogModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
