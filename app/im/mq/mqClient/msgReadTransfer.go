package mqClient

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"go-zero-dandan/app/im/mq/kafkad"
)

// MsgReadTransferClient 定义消息已读发送的队列工具的接口
type MsgReadTransferClient interface {
	Push(msg *kafkad.MsgMarkRead) error
}

// msgReadTransferClient 消息已读发送的kafka实现类
type msgReadTransferClient struct {
	pusher *kq.Pusher
}

func NewMsgReadTransferClient(addr []string, topic string, opts ...kq.PushOption) MsgReadTransferClient {
	return &msgReadTransferClient{
		pusher: kq.NewPusher(addr, topic, opts...),
	}
}
func (t *msgReadTransferClient) Push(msg *kafkad.MsgMarkRead) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return t.pusher.Push(string(body))
}
