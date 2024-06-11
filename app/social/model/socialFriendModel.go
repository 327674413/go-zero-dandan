package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SocialFriendModel = (*customSocialFriendModel)(nil)

type (
	// SocialFriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSocialFriendModel.
	SocialFriendModel interface {
		socialFriendModel
	}

	customSocialFriendModel struct {
		*defaultSocialFriendModel
		softDeletable bool
	}
)

// NewSocialFriendModel returns a model for the database table.
func NewSocialFriendModel(conn sqlx.SqlConn, platId ...int64) SocialFriendModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customSocialFriendModel{
		defaultSocialFriendModel: newSocialFriendModel(conn, platid),
		softDeletable:            true, //是否启用软删除
	}
}
