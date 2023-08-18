package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/app/user/rpc/user"
)

var _ UserCronyModel = (*customUserCronyModel)(nil)

type (
	// UserCronyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCronyModel.
	UserCronyModel interface {
		userCronyModel
		PbSelect2() (*pb.UserCronyList, error)
	}

	customUserCronyModel struct {
		*defaultUserCronyModel
		softDeletable bool
	}
)

// NewUserCronyModel returns a model for the database table.
func NewUserCronyModel(conn sqlx.SqlConn, platId ...int64) UserCronyModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customUserCronyModel{
		defaultUserCronyModel: newUserCronyModel(conn, platid),
		softDeletable:         true, //是否启用软删除
	}
}
func (m *defaultUserCronyModel) PbSelect() (*pb.UserCronyList, error) {
	list := make([]*pb.UserCronyInfo, 0)
	err := m.dao.Select(&list)
	if err != nil {
		return nil, err
	}
	return &pb.UserCronyList{
		List: list,
	}, nil
}
func (m *defaultUserCronyModel) PbSelect2() (*pb.UserCronyList, error) {
	//list := make([]*pb.UserCronyInfo, 0)
	list := make([]*user.UserCronyInfo, 0)
	err := m.conn.QueryRowPartialCtx(m.ctx, &list, "SELECT * FROM user_crony")
	if err != nil {
		logx.Error(err)
	}
	return &pb.UserCronyList{
		List: list,
	}, nil
}
