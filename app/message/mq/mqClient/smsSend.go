package mqClient

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
)

// SmsSendMsg 短信发送mq消息体
type SmsSendMsg struct {
	Id      string `json:"id"`
	Phone   string `json:"phone"`
	TempId  string `json:"tempId"`
	PlatId  string `json:"platId"`
	Content string `json:"content"`
}

// SmsSendCli 短信发送mq接口
type SmsSendCli interface {
	Push(msg *SmsSendMsg) error
}

// smsSendCli 短信发送的mq接口的kq实现
type smsSendCli struct {
	pusher *kq.Pusher
}

func (t *smsSendCli) Push(msg *SmsSendMsg) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return t.pusher.Push(string(body))
}

func NewSmsSendCli(addr []string, topic string, opts ...kq.PushOption) SmsSendCli {
	return &smsSendCli{
		pusher: kq.NewPusher(addr, topic, opts...),
	}
}
