package model

var _ SocialFriendModel = (*customSocialFriendModel)(nil)
var softDeletableSocialFriend = true

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
	// 自定义方法加在customSocialFriendModel上
)
