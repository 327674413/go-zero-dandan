package mqClient

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/mq/kafkad"
)

// MsgChatTransferClient 定义消息发送队列工具的接口
type MsgChatTransferClient interface {
	Push(msg *kafkad.MsgChatTransfer) error
}

// msgChatTransferClient 消息发送队列工具的kafka实现类
type msgChatTransferClient struct {
	pusher *kq.Pusher
}

func NewMsgChatTransferClient(addr []string, topic string, opts ...kq.PushOption) MsgChatTransferClient {
	logx.Debug("msgChatTransferClient 连接 topic：", topic)
	return &msgChatTransferClient{
		pusher: kq.NewPusher(addr, topic, opts...),
	}
}
func (t *msgChatTransferClient) Push(msg *kafkad.MsgChatTransfer) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return t.pusher.Push(string(body))
}
