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
func NewUserInfoModel(conn sqlx.SqlConn, platId ...int64) UserInfoModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customUserInfoModel{
		defaultUserInfoModel: newUserInfoModel(conn, platid),
		softDeletable:        true, //是否启用软删除
	}
}
