package model

var _ MessageSysConfigModel = (*customMessageSysConfigModel)(nil)
var softDeletableMessageSysConfig = true

type (
	// MessageSysConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageSysConfigModel.
	MessageSysConfigModel interface {
		messageSysConfigModel
	}

	customMessageSysConfigModel struct {
		*defaultMessageSysConfigModel
		softDeletable bool
	}
	// 自定义方法加在customMessageSysConfigModel上
)
