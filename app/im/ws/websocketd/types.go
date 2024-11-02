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
	MsgClasMarkSend            = constd.MsgClasEmMarkSend
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
	//Chat 聊天会话
	Chat struct {
		Id             string `mapstructure:"id" json:"id"`
		ConversationId string `mapstructure:"conversationId" json:"conversationId"`
		SendId         string `mapstructure:"sendId" json:"sendId"`
		RecvId         string `mapstructure:"recvId" json:"recvId"`
		MsgType        `mapstructure:"msgType" json:"msgType"`
		MsgContent     string                      `mapstructure:"msgContent" json:"msgContent"`
		ReadRecords    map[string]map[string]int32 `mapstructure:"readRecords" json:"readRecords"`
		MsgState       int64                       `mapstructure:"msgState" json:"msgState"`
		ChatType       `mapstructure:"chatType" json:"chatType"`
		MsgClas        `mapstructure:"msgClas" json:"msgClas"`
		SendTime       string `mapstructure:"sendTime" json:"sendTime"`
		SendAtMs       int64  `mapstructure:"sendAtMs" json:"sendAtMs"`
		TempId         string `mapstructure:"tempId" json:"tempId"`
	}
	// Push 解析kafka的消息
	Push struct {
		ConversationId string `mapstructure:"conversationId" json:"conversationId"`
		ChatType       `mapstructure:"chatType" json:"chatType"`
		Id             string                                  `mapstructure:"id" json:"id"`
		TempId         string                                  `mapstructure:"tempId" json:"tempId"`
		SendId         string                                  `mapstructure:"sendId" json:"sendId"`
		RecvId         string                                  `mapstructure:"recvId" json:"recvId"`
		RecvIds        []string                                `mapstructure:"recvIds" json:"recvIds"`
		SendTime       string                                  `mapstructure:"sendTime" json:"sendTime"`
		SendAtMs       int64                                   `mapstructure:"sendAtMs" json:"sendAtMs"`
		ReadRecords    map[string]map[string]int32             `mapstructure:"readRecords" json:"readRecords"` //消息id为key，消息已读情况为value的暂存群聊已读
		MsgState       int64                                   `mapstructure:"msgState" json:"msgState"`
		MsgClas        MsgClas                                 `mapstructure:"msgClas" json:"msgClas"` //业务类型：0聊天消息 1消息已读等
		MsgType        `mapstructure:"msgType" json:"msgType"` //消息数据类型：文本消息、图片消息等
		MsgContent     string                                  `mapstructure:"msgContent" json:"msgContent"`
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
