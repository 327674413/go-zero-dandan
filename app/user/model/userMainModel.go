package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserMainModel = (*customUserMainModel)(nil)

type (
	// UserMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMainModel.
	UserMainModel interface {
		userMainModel
		GetOne[T string | int | int64](ctx context.Context, rowBuilder squirrel.SelectBuilder, where T) (*UserMain, error)
	}

	customUserMainModel struct {
		*defaultUserMainModel
	}
)

// NewUserMainModel returns a model for the database table.
func NewUserMainModel(conn sqlx.SqlConn) UserMainModel {
	return &customUserMainModel{
		defaultUserMainModel: newUserMainModel(conn),
	}
}

func (m *defaultUserMainModel) GetOne[T int | int64 | string](ctx context.Context, rowBuilder squirrel.SelectBuilder, where T) (*UserMain, error) {
	resp := &UserMain{}
	query := fmt.Sprintf("select %s from %s where ? = ? limit 1", userMainRows, m.table)
	if err := conn.QueryRowCtx(ctx, &resp, query, where); err != nil {
		return nil, err
	}
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
