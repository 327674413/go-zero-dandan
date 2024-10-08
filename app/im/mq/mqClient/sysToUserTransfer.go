package mqClient

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"go-zero-dandan/app/im/mq/kafkad"
)

// SysToUserTransferClient 系统消息
type SysToUserTransferClient interface {
	Push(msg *kafkad.SysToUserMsg) error
}

// msgReadTransferClient 消息已读发送的kafka实现类
type sysMsgTransferClient struct {
	pusher *kq.Pusher
}

func NewSysToUserTransferClient(addr []string, topic string, opts ...kq.PushOption) SysToUserTransferClient {
	return &sysMsgTransferClient{
		pusher: kq.NewPusher(addr, topic, opts...),
	}
}
func (t *sysMsgTransferClient) Push(msg *kafkad.SysToUserMsg) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return t.pusher.Push(string(body))
}
