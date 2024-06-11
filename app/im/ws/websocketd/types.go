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

type ContentType int

const (
	ContentChatMsg ContentType = iota
	ContentMakeRead
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
		MsgType     `mapstructure:"msgType"`
		Content     string            `mapstructure:"content"`
		MsgId       string            `mapstructure:"msgId"`
		ReadRecords map[string]string `mapstructure:"readRecords"`
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
		MsgId          string            `mapstructure:"msgId"`
		SendId         int64             `mapstructure:"sendId,string"`
		RecvId         int64             `mapstructure:"recvId,string"`
		RecvIds        []int64           `mapstructure:"recvIds"`
		SendTime       int64             `mapstructure:"sendTime"`
		ReadRecords    map[string]string `mapstructure:"readRecords"`
		ContentType    ContentType       `mapstructure:"contentType"`
		MsgType        `mapstructure:"msgType"`
		Content        string `mapstructure:"content"`
	}
	MarkRead struct {
		ChatType       `mapstructure:"chatType"`
		RecvId         int64    `mapstructure:"recvId,string"`
		ConversationId string   `mapstructure:"conversationId"`
		MsgIds         []string `mapstructure:"msgIds"`
	}
)
