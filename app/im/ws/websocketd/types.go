package websocketd

import "go-zero-dandan/common/constd"

type MsgType int

const (
	TextMsgEmpty = constd.MsgTypeEmEmpty
	TextMsgType  = constd.MsgTypeEmText
)
const (
	RedisSystemRootToken = "system:root:token"
	RedisOnlineUser      = "online:user"
)

type ChatType int

const (
	ChatTypeGroup  ChatType = iota + 1 //发到群组
	ChatTypeSingle                     //发到个人
)

type AckType int

const (
	AckTypeNoAck AckType = iota
	AckTypeOnlyAck
	AckTypeRigorAck
)

type MsgClas int

const (
	MsgClasChatMsg             = constd.MsgClasEmChat
	MsgClasMakeRead            = constd.MsgClasEmMarkRead
	MsgClasSysAnnounceNew      = constd.MsgClasEmSysAnnounceNew
	MsgClasFriendApplyNew      = constd.MsgClasEmFriendApplyNew
	MsgClasFriendApplyOperated = constd.MsgClasEmFriendApplyOperated
)

func (t AckType) ToString() string {
	switch t {
	case AckTypeOnlyAck:
		return "OnlyAck"
	case AckTypeRigorAck:
		return "RigorAck"
	}
	return "NoAck"
}

type (
	//Msg 聊天消息内容
	Msg struct {
		MsgType     `mapstructure:"msgType"`
		Content     string            `mapstructure:"content"`
		MsgId       string            `mapstructure:"msgId"`
		ReadRecords map[string]string `mapstructure:"readRecords"`
	}
	//Chat 聊天会话
	Chat struct {
		ConversationId string `mapstructure:"conversationId"`
		SendId         string `mapstructure:"sendId"`
		RecvId         string `mapstructure:"recvId"`
		Msg            `mapstructure:"msg"`
		ChatType       `mapstructure:"chatType"`
		SendTime       string `mapstructure:"sendTime"`
	}
	// Push 解析kafka的消息
	Push struct {
		ConversationId string `mapstructure:"conversationId"`
		ChatType       `mapstructure:"chatType"`
		MsgId          string                   `mapstructure:"msgId"`
		SendId         string                   `mapstructure:"sendId"`
		RecvId         string                   `mapstructure:"recvId"`
		RecvIds        []string                 `mapstructure:"recvIds"`
		SendTime       string                   `mapstructure:"sendTime"`
		ReadRecords    map[string]string        `mapstructure:"readRecords"`
		MsgClas        MsgClas                  `mapstructure:"msgClas"` //业务类型：0聊天消息 1消息已读等
		MsgType        `mapstructure:"msgType"` //消息数据类型：文本消息、图片消息等
		Content        string                   `mapstructure:"content"`
	}
	// MarkRead 已读消息
	MarkRead struct {
		ChatType       `mapstructure:"chatType"`
		RecvId         string   `mapstructure:"recvId"`
		ConversationId string   `mapstructure:"conversationId"`
		MsgIds         []string `mapstructure:"msgIds"`
	}
	// SysMsg 系统消息
	SysMsg struct {
		MsgClas    MsgClas `mapstructure:"msgClas"`
		MsgType    MsgType `mapstructure:"msgTtpe"`
		MsgContent string  `mapstructure:"msgContent"`
		SendTime   string  `mapstructure:"sendTime"`
	}
)
