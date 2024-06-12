package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SocialGroupModel = (*customSocialGroupModel)(nil)

type (
	// SocialGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSocialGroupModel.
	SocialGroupModel interface {
		socialGroupModel
	}

	customSocialGroupModel struct {
		*defaultSocialGroupModel
		softDeletable bool
	}
)

// NewSocialGroupModel returns a model for the database table.
func NewSocialGroupModel(conn sqlx.SqlConn, platId ...string) SocialGroupModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customSocialGroupModel{
		defaultSocialGroupModel: newSocialGroupModel(conn, platid),
		softDeletable:           true, //是否启用软删除
	}
}
