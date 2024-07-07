package model

var _ SocialFriendApplyModel = (*customSocialFriendApplyModel)(nil)
var softDeletableSocialFriendApply = true

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
	// 自定义方法加在customSocialFriendApplyModel上
)
