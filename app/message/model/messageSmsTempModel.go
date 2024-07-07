package model

var _ MessageSmsTempModel = (*customMessageSmsTempModel)(nil)
var softDeletableMessageSmsTemp = true

type (
	// MessageSmsTempModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageSmsTempModel.
	MessageSmsTempModel interface {
		messageSmsTempModel
	}

	customMessageSmsTempModel struct {
		*defaultMessageSmsTempModel
		softDeletable bool
	}
	// 自定义方法加在customMessageSmsTempModel上
)
