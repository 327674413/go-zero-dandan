package msgTransfer

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/websocketd"
)

type MessageImSend struct {
	*baseMsgTransfer
}

func NewMessageImSend(svcCtx *svc.ServiceContext) *MessageImSend {
	return &MessageImSend{
		NewBaseMsgTransfer(svcCtx),
	}
}
func (t *MessageImSend) Consume(key, value string) error {
	var (
		data *kafkad.MsgChatTransfer
		ctx  = context.Background()
	)
	logx.Info("消费：", value)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		t.Errorf("msgTransfer 消费失败,err:%v", err)
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
		ContentType:    data.ContentType,
	})
}
