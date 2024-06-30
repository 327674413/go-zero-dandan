package mqClient

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"go-zero-dandan/app/im/ws/websocketd"
)

// ImSendMsg 消息发送
type ImSendMsg struct {
	websocketd.ChatType    `json:"chatType"`
	ChannelId              string   `json:"channelId"`
	SendId                 string   `json:"sendId"`
	RecvId                 string   `json:"recvId"`
	RecvIds                []string `json:"recvIds"`
	SendAt                 int64    `json:"sendAt"`
	websocketd.ContentType `json:"contentType"`
	websocketd.MsgType     `json:"msgType"`
	Content                string `json:"content"`
	PlatId                 string `json:"platId"`
	MsgId                  string `json:"msgId"`
}

// ImSendCli 短信发送mq接口
type ImSendCli interface {
	Push(msg *ImSendMsg) error
}

// imSendCli im消息发送的kq实现
type imSendCli struct {
	pusher *kq.Pusher
}

func (t *imSendCli) Push(msg *ImSendMsg) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return t.pusher.Push(string(body))
}
func NewImSendCli(addr []string, topic string, opts ...kq.PushOption) ImSendCli {
	return &imSendCli{
		pusher: kq.NewPusher(addr, topic, opts...),
	}
}
