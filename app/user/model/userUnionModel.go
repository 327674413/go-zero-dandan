package model

var _ UserUnionModel = (*customUserUnionModel)(nil)
var softDeletableUserUnion = true

type (
	// UserUnionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserUnionModel.
	UserUnionModel interface {
		userUnionModel
	}

	customUserUnionModel struct {
		*defaultUserUnionModel
		softDeletable bool
	}
	// 自定义方法加在customUserUnionModel上
)
