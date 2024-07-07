package model

var _ UserInfoModel = (*customUserInfoModel)(nil)
var softDeletableUserInfo = true

type (
	// UserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfoModel.
	UserInfoModel interface {
		userInfoModel
	}

	customUserInfoModel struct {
		*defaultUserInfoModel
		softDeletable bool
	}
	// 自定义方法加在customUserInfoModel上
)
