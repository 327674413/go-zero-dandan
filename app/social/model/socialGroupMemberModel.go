package model

var _ SocialGroupMemberModel = (*customSocialGroupMemberModel)(nil)
var softDeletableSocialGroupMember = true

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
	// 自定义方法加在customSocialGroupMemberModel上
)
