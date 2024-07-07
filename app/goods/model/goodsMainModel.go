package model

var _ GoodsMainModel = (*customGoodsMainModel)(nil)
var softDeletableGoodsMain = true

type (
	// GoodsMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsMainModel.
	GoodsMainModel interface {
		goodsMainModel
	}

	customGoodsMainModel struct {
		*defaultGoodsMainModel
		softDeletable bool
	}
	// 自定义方法加在customGoodsMainModel上
)
