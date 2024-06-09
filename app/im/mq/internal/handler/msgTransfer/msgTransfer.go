package msgTransfer

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/websocketd"
)

type MsgChatTransfer struct {
	logx.Logger
	svc *svc.ServiceContext
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		Logger: logx.WithContext(context.Background()),
		svc:    svc,
	}
}
func (t *MsgChatTransfer) Consume(key, value string) error {
	var (
		data kafkad.MsgChatTransfer
		ctx  = context.Background()
	)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		logx.Errorf("msgTransfer 消费失败,err:%v", err)
	}
	//记录数据
	if err := t.addChatLog(ctx, &data); err != nil {
		logx.Errorf("msgTransfer 写入消息,err:%v", err)
	}

	//写推送消息
	return t.svc.WsClient.Send(websocketd.Message{
		FrameType: websocketd.FrameData,
		Method:    "push",
		FormCode:  "chat_system_root",
		Data:      data,
	})
}
func (t *MsgChatTransfer) addChatLog(ctx context.Context, data *kafkad.MsgChatTransfer) error {
	chatLog := modelMongo.ChatLog{
		ChatType:       data.ChatType,
		MsgType:        data.MsgType,
		SendTime:       data.SendTime,
		MsgContent:     data.Content,
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
	}
	err := t.svc.ChatLogModel.Insert(ctx, &chatLog)
	if err != nil {
		return err
	}
	return t.svc.ConversationModel.UpdateMsg(ctx, &chatLog)
}
