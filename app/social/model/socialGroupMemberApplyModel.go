package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SocialGroupMemberApplyModel = (*customSocialGroupMemberApplyModel)(nil)

type (
	// SocialGroupMemberApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSocialGroupMemberApplyModel.
	SocialGroupMemberApplyModel interface {
		socialGroupMemberApplyModel
	}

	customSocialGroupMemberApplyModel struct {
		*defaultSocialGroupMemberApplyModel
		softDeletable bool
	}
)

// NewSocialGroupMemberApplyModel returns a model for the database table.
func NewSocialGroupMemberApplyModel(conn sqlx.SqlConn, platId ...string) SocialGroupMemberApplyModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customSocialGroupMemberApplyModel{
		defaultSocialGroupMemberApplyModel: newSocialGroupMemberApplyModel(conn, platid),
		softDeletable:                      true, //是否启用软删除
	}
}
