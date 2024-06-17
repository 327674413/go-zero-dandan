package mqclient

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
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
