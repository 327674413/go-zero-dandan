package constd

const (
	MsgTypeEmEmpty = 0 //无内容消息，使用消息类型就行
	MsgTypeEmText  = 1 //纯文本类型的消息
)
const (
	ChatTypeGroup  = 1
	ChatTypeSingle = 2
)
const (
	MsgClasEmChat                = 1
	MsgClasEmMarkRead            = 2
	MsgClasEmSysAnnounceNew      = 3
	MsgClasEmFriendApplyNew      = 4
	MsgClasEmFriendApplyOperated = 5
	MsgClasEmMarkSend            = 6 //消息服务器已收到，推入kafka
)

const (
	MsgStateEmSent   = 1 //已送达
	MsgStateEmRead   = 2 //已读
	MsgStateEmCancel = 3 //撤回
)
