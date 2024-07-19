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
		Content     string            `mapstructure:"content" json:"content"`
		MsgId       string            `mapstructure:"msgId" json:"msgId"`
		ReadRecords map[string]string `mapstructure:"readRecords" json:"readRecords"`
	}
	//Chat 聊天会话
	Chat struct {
		ConversationId string `mapstructure:"conversationId" json:"conversationId"`
		SendId         string `mapstructure:"sendId" json:"sendId"`
		RecvId         string `mapstructure:"recvId" json:"recvId"`
		Msg            `mapstructure:"msg" json:"msg"`
		ChatType       `mapstructure:"chatType" json:"chatType"`
		SendTime       string `mapstructure:"sendTime" json:"sendTime"`
	}
	// Push 解析kafka的消息
	Push struct {
		ConversationId string `mapstructure:"conversationId" json:"conversationId"`
		ChatType       `mapstructure:"chatType" json:"chatType"`
		MsgId          string                                  `mapstructure:"msgId" json:"msgId"`
		SendId         string                                  `mapstructure:"sendId" json:"sendId"`
		RecvId         string                                  `mapstructure:"recvId" json:"recvId"`
		RecvIds        []string                                `mapstructure:"recvIds" json:"recvIds"`
		SendTime       string                                  `mapstructure:"sendTime" json:"sendTime"`
		ReadRecords    map[string]string                       `mapstructure:"readRecords" json:"readRecords"`
		MsgClas        MsgClas                                 `mapstructure:"msgClas" json:"msgClas"` //业务类型：0聊天消息 1消息已读等
		MsgType        `mapstructure:"msgType" json:"msgType"` //消息数据类型：文本消息、图片消息等
		Content        string                                  `mapstructure:"content" json:"content"`
	}
	// MarkRead 已读消息
	MarkRead struct {
		ChatType       `mapstructure:"chatType" json:"chatType"`
		RecvId         string   `mapstructure:"recvId" json:"recvId"`
		ConversationId string   `mapstructure:"conversationId" json:"conversationId"`
		MsgIds         []string `mapstructure:"msgIds" json:"msgIds"`
	}
	// SysMsg 系统消息
	SysMsg struct {
		MsgClas    MsgClas `mapstructure:"msgClas" json:"msgClas"`
		MsgType    MsgType `mapstructure:"msgTtpe" json:"msgType"`
		MsgContent string  `mapstructure:"msgContent" json:"msgContent"`
		SendTime   string  `mapstructure:"sendTime" json:"sendTime"`
	}
)
