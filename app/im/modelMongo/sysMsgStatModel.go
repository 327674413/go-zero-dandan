package modelMongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go-zero-dandan/common/resd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ SysMsgStatModel = (*customSysMsgStatModel)(nil)

type (
	// SysMsgStatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMsgStatModel.
	SysMsgStatModel interface {
		sysMsgStatModel
	}

	customSysMsgStatModel struct {
		*defaultSysMsgStatModel
	}
)

// NewSysMsgStatModel returns a model for the mongo.
func NewSysMsgStatModel(url, db, collection string) SysMsgStatModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customSysMsgStatModel{
		defaultSysMsgStatModel: newDefaultSysMsgStatModel(conn),
	}
}
func MustSysMsgStatModel(url, db string) SysMsgStatModel {
	return NewSysMsgStatModel(url, db, "sys_msg_stat")
}
func (t *customSysMsgStatModel) StartSession() (sess mongo.Session, err error) {
	sess, err = t.conn.StartSession()
	return sess, err
}
func (t *customSysMsgStatModel) Conn() *mon.Model {
	return t.conn
}
func (m *defaultSysMsgStatModel) IncSysMsgUnreadNum(ctx context.Context, userId string, msgClasEm int64, incNum int64) (effectedNun int64, danErr error) {
	update := bson.M{
		"$inc": bson.M{"unreadNum": incNum},
	}
	where := bson.M{"userId": userId, "msgClasEm": msgClasEm}
	opts := options.Update().SetUpsert(true)
	res, err := m.conn.UpdateOne(ctx, where, update, opts)
	if err != nil {
		return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return res.ModifiedCount, err
}
func (m *defaultSysMsgStatModel) DecSysMsgUnreadNum(ctx context.Context, userId string, msgClasEm int64, decNum int64) (effectedNun int64, danErr error) {
	update := bson.M{
		"$dec": bson.M{"unreadNum": decNum},
	}
	where := bson.M{"userId": userId, "msgClasEm": msgClasEm}
	opts := options.Update()
	res, err := m.conn.UpdateOne(ctx, where, update, opts)
	if err != nil {
		return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return res.ModifiedCount, err
}
func (m *defaultSysMsgStatModel) SetZeroSysMsgUnreadNum(ctx context.Context, userId string, msgClasEm []int64) (effectedNun int64, danErr error) {
	update := bson.M{
		"$set": bson.M{"unreadNum": 0},
	}
	where := bson.M{"userId": userId, "msgClasEm": bson.M{"$in": msgClasEm}}
	opts := options.Update()
	res, err := m.conn.UpdateOne(ctx, where, update, opts)
	if err != nil {
		return 0, resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return res.ModifiedCount, err
}
func (m *defaultSysMsgStatModel) GetUserUnread(ctx context.Context, userId string, msgClasEm int64) ([]*SysMsgStat, error) {

	data := make([]*SysMsgStat, 0)
	where := bson.M{"userId": userId}
	if msgClasEm > 0 {
		where["msgClasEm"] = msgClasEm
	}
	err := m.conn.Find(ctx, &data, where)
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, nil
	default:
		return nil, resd.ErrorCtx(ctx, err, resd.ErrMongoSelect)
	}
}

func (m *defaultSysMsgStatModel) SetSysMsgRead(ctx context.Context, userId string, msgClasEm []int64) (*mongo.UpdateResult, error) {
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
		return nil, resd.ErrorCtx(ctx, err, resd.ErrMongoUpdate)
	}
	return res, err
}
