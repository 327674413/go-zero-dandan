package mqclient

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"go-zero-dandan/app/chat/mq/kafkad"
)

type MsgChatTransferClient interface {
	Push(msg *kafkad.MsgChatTransfer) error
}
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
