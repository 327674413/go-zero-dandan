package msgTransfer

import (
	"context"
	"encoding/json"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/websocketd"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SysToUserTransfer struct {
	*baseMsgTransfer
}

func NewSysToUserTransfer(svcCtx *svc.ServiceContext) *SysToUserTransfer {
	return &SysToUserTransfer{
		NewBaseMsgTransfer(svcCtx),
	}
}
func (t *SysToUserTransfer) Consume(key, value string) error {
	var (
		data  *kafkad.SysToUserMsg
		ctx   = context.Background()
		msgId = primitive.NewObjectID()
	)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		t.Errorf("msgTransfer 消费失败,err:%v", err)
	}
	//写入消息数据到mongo
	if err := t.addSysMsgLog(ctx, msgId, data); err != nil {
		t.Errorf("msgTransfer 写入消息,err:%v", err)
	}

	//发送给ws进行push
	return t.Transfer(ctx, &websocketd.Push{
		MsgClas:  data.MsgClas,
		RecvId:   data.RecvId,
		MsgType:  data.MsgType,
		Content:  data.MsgContent,
		SendTime: data.SendTime,
		ChatType: websocketd.ChatTypeSingle,
	})
}
func (t *SysToUserTransfer) addSysMsgLog(ctx context.Context, msgId primitive.ObjectID, data *kafkad.SysToUserMsg) error {
	//todo::如果用事务需要mongo配置集群，没弄成功，以后研究
	/*
		sess, err := t.svc.SysMsgStatModel.StartSession()
		if err != nil {
			return err
		}
		defer sess.EndSession(ctx)
		// 开始事务，目前拿不到原生连接，不能主动提交或回滚，得用下面封装的方式自动提交回滚
		res, err := sess.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			//存储消息数据
			if err := t.svc.SysMsgLogModel.Insert(ctx, &modelMongo.SysMsgLog{
				ID:         msgId,
				MsgContent: data.MsgContent,
				RecvId:     data.RecvId,
				IsRead:     0,
				MsgClas:    data.MsgClas,
				CreateAt:   time.Now(),
				UpdateAt:   time.Now(),
			}); err != nil {
				return nil, err
			}
			if _, err := t.svc.SysMsgStatModel.IncSysMsgUnreadNum(ctx, data.RecvId, data.MsgClas, 1); err != nil {
				return nil, err
			}
			if err != nil {
				return nil, err
			}
			return nil, nil
		})
	*/
	if err := t.svc.SysMsgLogModel.Insert(ctx, &modelMongo.SysMsgLog{
		ID:         msgId,
		MsgContent: data.MsgContent,
		RecvId:     data.RecvId,
		IsRead:     0,
		MsgClasEm:  data.MsgClas,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}); err != nil {
		return err
	}
	if _, err := t.svc.SysMsgStatModel.IncSysMsgUnreadNum(ctx, data.RecvId, int64(data.MsgClas), 1); err != nil {
		return err
	}
	return nil
}
