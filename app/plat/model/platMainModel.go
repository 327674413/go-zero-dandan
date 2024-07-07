package model

var _ PlatMainModel = (*customPlatMainModel)(nil)
var softDeletablePlatMain = true

type (
	// PlatMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatMainModel.
	PlatMainModel interface {
		platMainModel
	}

	customPlatMainModel struct {
		*defaultPlatMainModel
		softDeletable bool
	}
	// 自定义方法加在customPlatMainModel上
)
