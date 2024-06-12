package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SocialGroupMemberModel = (*customSocialGroupMemberModel)(nil)

type (
	// SocialGroupMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSocialGroupMemberModel.
	SocialGroupMemberModel interface {
		socialGroupMemberModel
	}

	customSocialGroupMemberModel struct {
		*defaultSocialGroupMemberModel
		softDeletable bool
	}
)

// NewSocialGroupMemberModel returns a model for the database table.
func NewSocialGroupMemberModel(conn sqlx.SqlConn, platId ...string) SocialGroupMemberModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customSocialGroupMemberModel{
		defaultSocialGroupMemberModel: newSocialGroupMemberModel(conn, platid),
		softDeletable:                 true, //是否启用软删除
	}
}
