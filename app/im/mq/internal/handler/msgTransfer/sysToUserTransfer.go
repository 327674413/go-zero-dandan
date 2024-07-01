package msgTransfer

import (
	"context"
	"encoding/json"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/websocketd"
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
		data *kafkad.SysToUserMsg
		ctx  = context.Background()
	)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		t.Errorf("msgTransfer 消费失败,err:%v", err)
	}
	//发送给ws进行push
	return t.Transfer(ctx, &websocketd.Push{
		MsgClas:  data.MsgClas,
		RecvId:   data.UserId,
		MsgType:  data.MsgType,
		Content:  data.MsgContent,
		SendTime: data.SendTime,
		ChatType: websocketd.ChatTypeSingle,
	})
}
