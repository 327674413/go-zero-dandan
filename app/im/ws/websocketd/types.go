package websocketd

type MsgType int

const (
	TextMsgType MsgType = iota //iota常量自动从0递增的赋值
)

type ChatType int

const (
	GroupChatType ChatType = iota + 1
	SingleChatType
)

type AckType int

const (
	NoAck AckType = iota
	OnlyAck
	RigorAck
)

func (t AckType) ToString() string {
	switch t {
	case OnlyAck:
		return "OnlyAck"
	case RigorAck:
		return "RigorAck"
	}
	return "NoAck"
}

type (
	Msg struct {
		MsgType `mapstructure:"msgType"`
		Content string `mapstructure:"content"`
	}
	Chat struct {
		ConversationId string `mapstructure:"conversationId"`
		SendId         int64  `mapstructure:"sendId"`
		RecvId         int64  `mapstructure:"recvId"`
		Msg            `mapstructure:"msg"`
		ChatType       `mapstructure:"chatType"`
		SendTime       int64 `mapstructure:"sendTime"`
	}
	Push struct {
		ConversationId string `mapstructure:"conversationId"`
		ChatType       `mapstructure:"chatType"`
		SendId         int64 `mapstructure:"sendId,string"`
		RecvId         int64 `mapstructure:"recvId,string"`
		SendTime       int64 `mapstructure:"sendTime"`
		MsgType        `mapstructure:"msgType"`
		Content        string `mapstructure:"content"`
	}
)
