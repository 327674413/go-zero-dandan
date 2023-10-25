package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserAuthRuleModel = (*customUserAuthRuleModel)(nil)

type (
	// UserAuthRuleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthRuleModel.
	UserAuthRuleModel interface {
		userAuthRuleModel
	}

	customUserAuthRuleModel struct {
		*defaultUserAuthRuleModel
		softDeletable bool
	}
)

// NewUserAuthRuleModel returns a model for the database table.
func NewUserAuthRuleModel(conn sqlx.SqlConn, platId ...int64) UserAuthRuleModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customUserAuthRuleModel{
		defaultUserAuthRuleModel: newUserAuthRuleModel(conn, platid),
		softDeletable:            true, //是否启用软删除
	}
}
