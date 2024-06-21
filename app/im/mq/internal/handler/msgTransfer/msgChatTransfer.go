package msgTransfer

import (
	"context"
	"encoding/json"
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
		t.Errorf("msgTransfer 消费失败,err:%v", err)
	}
	//写入消息数据到mongo
	if err := t.addChatLog(ctx, msgId, data); err != nil {
		t.Errorf("msgTransfer 写入消息,err:%v", err)
	}
	//发送给ws进行push
	return t.Transfer(ctx, &websocketd.Push{
		ChatType:       data.ChatType,
		MsgType:        data.MsgType,
		MsgId:          data.MsgId,
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
	//将此条消息的发送者设置为已读
	readRecords := bitmapd.NewBitmap()
	readRecords.SetId(chatLog.SendId)
	chatLog.ReadRecords = readRecords.Export()
	//存储消息数据
	err := t.svc.ChatLogModel.Insert(ctx, &chatLog)
	if err != nil {
		return err
	}
	//更新该消息所在会话中的消息总数加1
	return t.svc.ConversationModel.UpdateMsg(ctx, &chatLog)
}
