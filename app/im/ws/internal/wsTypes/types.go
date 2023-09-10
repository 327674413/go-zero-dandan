package wsTypes

const (
	TargetTypeEmCrony = "crony" //好友
)
const (
	ChatMsgStateEmSending = 1 //发送中，未确定对方在线或离线
	ChatMsgStateEmSent    = 2 //已发送，在线则对方已接收，离线则写入离线缓存
	ChatMsgStateEmRead    = 3 //已读
)
const (
	MsgTypeEmChat     = "CHAT"      //私聊
	MsgTypeEmChatResp = "CHAT_RESP" //私聊状态回复
)

type Message struct {
	Id           int64  `json:"id,string"`       // 消息id
	Code         string `json:"code"`            //请求编号
	FromId       int64  `json:"fromId,string"`   //发送者
	TargetId     int64  `json:"targetId,string"` //接收者
	TargetTypeEm string `json:"targetTypeEm"`    //目标类型
	TypeEm       string `json:"typeEm"`          //发送类型：私聊、群聊、广播、红包、视频、语音、抖动
	Media        int64  `json:"media"`           //消息类型：文字、图片、语音、视频、文件、导航地址、h5分享链接
	Content      string `json:"content"`         //消息内容
	Pic          string `json:"pic"`
	Url          string `json:"url"`
	Desc         string `json:"desc"`
	Amount       int64  `json:"amount"` //其他统计
}
type MessageResp struct {
	TypeEm  string `json:"typeEm"`  //消息类型
	Code    string `json:"code"`    //请求编号
	StateEm int64  `json:"stateEm"` //状态，1已发送未读，2已读
	ErrCode int64  `json:"errCode"` //错误代码
}
type ConnectionReq struct {
	Token      string `form:"token"`
	PlatformEm int64  `form:"platformEm"`
}

type ConnectionRes struct {
	UserId int64 `json:"userId,string"`
}
