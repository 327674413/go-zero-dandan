package model

var _ SocialGroupModel = (*customSocialGroupModel)(nil)
var softDeletableSocialGroup = true

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
	// 自定义方法加在customSocialGroupModel上
)
