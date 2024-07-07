package model

var _ MessageSmsSendModel = (*customMessageSmsSendModel)(nil)
var softDeletableMessageSmsSend = true

type (
	// MessageSmsSendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageSmsSendModel.
	MessageSmsSendModel interface {
		messageSmsSendModel
	}

	customMessageSmsSendModel struct {
		*defaultMessageSmsSendModel
		softDeletable bool
	}
	// 自定义方法加在customMessageSmsSendModel上
)
