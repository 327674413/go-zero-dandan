package modelMongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go-zero-dandan/common/resd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ SysMsgLogModel = (*customSysMsgLogModel)(nil)

type (
	// SysMsgLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMsgLogModel.
	SysMsgLogModel interface {
		sysMsgLogModel
	}

	customSysMsgLogModel struct {
		*defaultSysMsgLogModel
	}
)

// NewSysMsgLogModel returns a model for the mongo.
func NewSysMsgLogModel(url, db, collection string) SysMsgLogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customSysMsgLogModel{
		defaultSysMsgLogModel: newDefaultSysMsgLogModel(conn),
	}
}

func MustSysMsgLogModel(url, db string) SysMsgLogModel {
	return NewSysMsgLogModel(url, db, "sys_msg_log")
}

func (m *defaultSysMsgLogModel) SetSysMsgReadByMsgClas(ctx context.Context, userId string, msgClasEm []int64) (updateNum int64, danErr error) {
	update := bson.M{
		"$set": bson.M{"unreadNum": 0},
	}
	where := bson.M{"userId": userId, "unreadNum": bson.M{"$ne": 0}}
	if len(msgClasEm) > 0 {
		where["msgClasEm"] = bson.M{"$in": msgClasEm}
	}
	opts := options.Update()
	res, err := m.conn.UpdateMany(ctx, where, update, opts)
	if err != nil {
		return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return res.ModifiedCount, err
}
func (m *defaultSysMsgLogModel) SetSysMsgReadByIds(ctx context.Context, userId string, msgIds []string, msgClasEm int64) (updateNum int64, danErr error) {
	ids := make([]primitive.ObjectID, 0)
	for _, v := range msgIds {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoStrToId)
		}
		ids = append(ids, id)
	}
	update := bson.M{
		"$set": bson.M{"isRead": 1},
	}
	where := bson.M{"userId": userId, "msgClasEm": msgClasEm, "isRead": 0, "_id": bson.M{"$in": ids}}
	opts := options.Update()
	res, err := m.conn.UpdateOne(ctx, where, update, opts)
	if err != nil {
		return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return res.ModifiedCount, err
}
