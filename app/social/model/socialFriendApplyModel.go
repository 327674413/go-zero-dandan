package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SocialFriendApplyModel = (*customSocialFriendApplyModel)(nil)

type (
	// SocialFriendApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSocialFriendApplyModel.
	SocialFriendApplyModel interface {
		socialFriendApplyModel
	}

	customSocialFriendApplyModel struct {
		*defaultSocialFriendApplyModel
		softDeletable bool
	}
)

// NewSocialFriendApplyModel returns a model for the database table.
func NewSocialFriendApplyModel(conn sqlx.SqlConn, platId ...string) SocialFriendApplyModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customSocialFriendApplyModel{
		defaultSocialFriendApplyModel: newSocialFriendApplyModel(conn, platid),
		softDeletable:                 true, //是否启用软删除
	}
}
