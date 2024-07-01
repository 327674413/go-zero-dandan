package msgTransfer

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/app/social/rpc/social"
)

// baseMsgTransfer 定义mq消息处理器的基类，封装各个消费业务的通用功能
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

// Transfer 消息发送的工厂入口，根据消息类型执行不同的发送方法，目前来看分为单用户推送、群组推送（往群找人），todo::后面可能要加组织推送
func (t *baseMsgTransfer) Transfer(ctx context.Context, data *websocketd.Push) error {
	var err error
	switch data.ChatType {
	case websocketd.ChatTypeSingle:
		err = t.single(ctx, data)
	case websocketd.ChatTypeGroup:
		err = t.group(ctx, data)
	default:
		return fmt.Errorf("push暂不支持的chatType：%d", data.ChatType)
	}
	return err
}

// channel 跟私聊应该一样，但可能需要检测不存在会话就先创建会话
func (t *baseMsgTransfer) channel(ctx context.Context, data *websocketd.Push) error {
	return t.svc.WsClient.Send(websocketd.Message{
		FrameType: websocketd.FrameData,
		Method:    "push",
		FormCode:  "chat_system_root", //目前这个formcode的作用不清楚
		Data:      data,
	})
}

// single 私聊消息发送，借助ws客户端，走ws的push类型消息的方式发送
func (t *baseMsgTransfer) single(ctx context.Context, data *websocketd.Push) error {
	return t.svc.WsClient.Send(websocketd.Message{
		FrameType: websocketd.FrameData,
		Method:    "push",
		FormCode:  "chat_system_root", //目前这个formcode的作用不清楚
		Data:      data,
	})
}

// group 群聊消息发送，借助ws客户端，走ws的push类型消息的方式发送
func (t *baseMsgTransfer) group(ctx context.Context, data *websocketd.Push) error {
	//群聊时，根据消息数据中的接受者id，即群id，查询出该群的用户列表
	users, err := t.svc.SocialRpc.GetGroupMemberList(ctx, &social.GetGroupMemberListReq{
		GroupId: data.RecvId,
		PlatId:  "1",
	})
	if err != nil {
		return err
	}
	//组装要推送的id列表，并过滤发送者自己
	data.RecvIds = make([]string, 0, len(users.List))
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
