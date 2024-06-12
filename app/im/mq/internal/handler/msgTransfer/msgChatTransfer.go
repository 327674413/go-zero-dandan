package msgTransfer

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/pkg/bitmapd"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MsgChatTransfer struct {
	*baseMsgTransfer
}

func NewMsgChatTransfer(svcCtx *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		NewBaseMsgTransfer(svcCtx),
	}
}
func (t *MsgChatTransfer) Consume(key, value string) error {
	var (
		data  *kafkad.MsgChatTransfer
		ctx   = context.Background()
		msgId = primitive.NewObjectID()
	)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		logx.Errorf("msgTransfer 消费失败,err:%v", err)
	}
	//记录数据
	if err := t.addChatLog(ctx, msgId, data); err != nil {
		logx.Errorf("msgTransfer 写入消息,err:%v", err)
	}

	return t.Transfer(ctx, &websocketd.Push{
		ChatType:       data.ChatType,
		MsgType:        data.MsgType,
		MsgId:          msgId.Hex(),
		SendTime:       data.SendTime,
		Content:        data.Content,
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		RecvIds:        data.RecvIds,
	})
}

func (t *MsgChatTransfer) addChatLog(ctx context.Context, msgId primitive.ObjectID, data *kafkad.MsgChatTransfer) error {
	chatLog := modelMongo.ChatLog{
		ID:             msgId,
		ChatType:       data.ChatType,
		MsgType:        data.MsgType,
		SendTime:       data.SendTime,
		MsgContent:     data.Content,
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
	}
	//将发送者设置为已读
	readRecords := bitmapd.NewBitmap()
	readRecords.SetId(chatLog.SendId)
	chatLog.ReadRecords = readRecords.Export()

	err := t.svc.ChatLogModel.Insert(ctx, &chatLog)
	if err != nil {
		return err
	}
	return t.svc.ConversationModel.UpdateMsg(ctx, &chatLog)
}
