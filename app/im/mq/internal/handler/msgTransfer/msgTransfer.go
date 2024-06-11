package msgTransfer

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/app/social/rpc/social"
)

type baseMsgTransfer struct {
	svc *svc.ServiceContext
	logx.Logger
}

func NewBaseMsgTransfer(svcCtx *svc.ServiceContext) *baseMsgTransfer {
	return &baseMsgTransfer{
		svc:    svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (t *baseMsgTransfer) Transfer(ctx context.Context, data *websocketd.Push) error {
	var err error
	switch data.ChatType {
	case websocketd.SingleChatType:
		err = t.single(ctx, data)
	case websocketd.GroupChatType:
		err = t.group(ctx, data)
	}
	return err
}
func (t *baseMsgTransfer) single(ctx context.Context, data *websocketd.Push) error {
	return t.svc.WsClient.Send(websocketd.Message{
		FrameType: websocketd.FrameData,
		Method:    "push",
		FormCode:  "chat_system_root",
		Data:      data,
	})
}
func (t *baseMsgTransfer) group(ctx context.Context, data *websocketd.Push) error {
	//要查询群的用户
	users, err := t.svc.SocialRpc.GroupUsers(ctx, &social.GroupUsersReq{
		GroupId: data.RecvId,
		PlatId:  1,
	})
	if err != nil {
		return err
	}
	data.RecvIds = make([]int64, 0, len(users.List))
	for _, item := range users.List {
		if item.UserId == data.SendId {
			continue
		}
		data.RecvIds = append(data.RecvIds, item.UserId)
	}
	return t.svc.WsClient.Send(websocketd.Message{
		FrameType: websocketd.FrameData,
		Method:    "push",
		FormCode:  "group msg",
		Data:      data,
	})
}
