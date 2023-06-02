package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserMainModel = (*customUserMainModel)(nil)

type (
	// UserMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMainModel.
	UserMainModel interface {
		userMainModel
		Field(field string) *defaultUserMainModel
		Alias(alias string) *defaultUserMainModel
		WhereStr(whereStr string) *defaultUserMainModel
		WhereId(id int) *defaultUserMainModel
		WhereMap(whereMap map[string]any) *defaultUserMainModel
		WhereRaw(whereStr string, whereData []any) *defaultUserMainModel
		Order(order string) *defaultUserMainModel
		Plat(id int) *defaultUserMainModel
		Find(ctx context.Context, id ...any) (*UserMain, error)
		Page(ctx context.Context, page int, rows int) ([]*UserMain, error)
		List(ctx context.Context) ([]*UserMain, error)
		Count(ctx context.Context) int
		Inc(ctx context.Context, field string, num int) error
		Dec(ctx context.Context, field string, num int) error
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

func (m *defaultUserMainModel) Find(ctx context.Context, id ...any) (*UserMain, error) {
	var err error
	if err = m.err; err != nil {
		m.err = nil
		return nil, err
	}
	var resp UserMain
	var sql string
	if len(id) > 0 {
		sql = fmt.Sprintf("select %s from %s where id=? limit 1", userMainRows, m.table)
		err = m.conn.QueryRowCtx(ctx, &resp, sql, id[0])
	} else {
		sql = fmt.Sprintf("select %s from %s where "+m.whereSql+" limit 1", userMainRows, m.table)
		fmt.Println(m.whereData)
		err = m.conn.QueryRowCtx(ctx, &resp, sql, m.whereData...)
	}
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		fmt.Println(err)
		return nil, err
	}
}
func (m *defaultUserMainModel) Page(ctx context.Context, page int, rows int) ([]*UserMain, error) {

	return nil, nil
}
func (m *defaultUserMainModel) List(ctx context.Context) ([]*UserMain, error) {

	return nil, nil
}
func (m *defaultUserMainModel) WhereStr(whereStr string) *defaultUserMainModel {
	return m
}
func (m *defaultUserMainModel) WhereId(id int) *defaultUserMainModel {
	m.whereSql = "id=?"
	m.whereData = append(m.whereData, id)
	return m
}
func (m *defaultUserMainModel) WhereMap(whereMap map[string]any) *defaultUserMainModel {
	return m
}
func (m *defaultUserMainModel) WhereRaw(whereStr string, whereData []any) *defaultUserMainModel {

	return m
}
func (m *defaultUserMainModel) Alias(field string) *defaultUserMainModel {

	return m
}
func (m *defaultUserMainModel) Field(field string) *defaultUserMainModel {

	return m
}
func (m *defaultUserMainModel) Order(order string) *defaultUserMainModel {

	return m
}
func (m *defaultUserMainModel) Count(ctx context.Context) int {

	return 0
}
func (m *defaultUserMainModel) Inc(ctx context.Context, field string, num int) error {

	return nil
}
func (m *defaultUserMainModel) Dec(ctx context.Context, field string, num int) error {

	return nil
}
func (m *defaultUserMainModel) Plat(id int) *defaultUserMainModel {

	return nil
}
