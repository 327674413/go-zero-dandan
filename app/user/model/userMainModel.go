package model

var _ UserMainModel = (*customUserMainModel)(nil)
var softDeletableUserMain = true

type (
	// UserMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMainModel.
	UserMainModel interface {
		userMainModel
	}

	customUserMainModel struct {
		*defaultUserMainModel
		softDeletable bool
	}
	// 自定义方法加在customUserMainModel上
)
