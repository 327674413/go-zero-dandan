package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserInfoModel = (*customUserInfoModel)(nil)

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
)

// NewUserInfoModel returns a model for the database table.
func NewUserInfoModel(conn sqlx.SqlConn, platId ...string) UserInfoModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customUserInfoModel{
		defaultUserInfoModel: newUserInfoModel(conn, platid),
		softDeletable:        true, //是否启用软删除
	}
}
