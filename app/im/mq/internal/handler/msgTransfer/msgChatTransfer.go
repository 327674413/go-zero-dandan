package msgTransfer

import (
	"context"
	"encoding/json"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/fmtd"
	"go-zero-dandan/common/resd"
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
	fmtd.Info(data.MsgContent)
	//发送给ws进行push
	return t.Transfer(ctx, &websocketd.Push{
		Id:             msgId.Hex(),
		ChatType:       data.ChatType,
		MsgType:        data.MsgType,
		SendTime:       data.SendTime,
		SendAtMs:       data.SendAtMs,
		MsgContent:     data.MsgContent,
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		RecvIds:        data.RecvIds,
		MsgClas:        constd.MsgClasEmChat,
		TempId:         data.TempId,
	})
}

func (t *MsgChatTransfer) addChatLog(ctx context.Context, msgId primitive.ObjectID, data *kafkad.MsgChatTransfer) error {
	chatLog := modelMongo.ChatLog{
		ID:             msgId,
		ChatType:       data.ChatType,
		MsgType:        data.MsgType,
		SendTime:       data.SendTime,
		SendAtMs:       data.SendAtMs,
		MsgContent:     data.MsgContent,
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		TempId:         data.TempId,
	}
	fmtd.Info(chatLog.SendAtMs)
	//将此条消息的发送者设置为已读(每条消息都记录是否已读方式)
	readRecords := bitmapd.NewBitmap()
	readRecords.SetId(chatLog.SendId)
	chatLog.MsgReads = readRecords.Export()
	//存储消息数据
	err := t.svc.ChatLogModel.Insert(ctx, &chatLog)
	if err != nil {
		return resd.ErrorCtx(ctx, err)
	}
	//更新该消息所在会话中的消息总数加1
	err = t.svc.ConversationModel.UpdateMsg(ctx, &chatLog)
	if err != nil {
		return resd.ErrorCtx(ctx, err)
	}
	//更新我的消息列表内容
	err = t.svc.ConversationsModel.UpdateMsg(ctx, chatLog.SendId, &chatLog)
	if err != nil {
		return resd.ErrorCtx(ctx, err)
	}
	return nil
}
