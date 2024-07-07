package model

var _ SocialGroupMemberApplyModel = (*customSocialGroupMemberApplyModel)(nil)
var softDeletableSocialGroupMemberApply = true

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
	// 自定义方法加在customSocialGroupMemberApplyModel上
)
