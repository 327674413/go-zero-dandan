package msgTransfer

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
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
	logx.Info("消费：", value)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		t.Errorf("msgTransfer 消费失败,err:%v", err)
	}
	//发送给ws进行push
	return t.Transfer(ctx, &websocketd.Push{
		MsgClas:  websocketd.MsgClasFriendApplyOperated,
		RecvId:   data.UserId,
		MsgType:  websocketd.MsgType(data.MsgType),
		Content:  data.MsgContent,
		SendTime: data.SendTime,
	})
}
