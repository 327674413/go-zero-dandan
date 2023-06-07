package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/plat/api/global"
)

var _ PlatMainModel = (*customPlatMainModel)(nil)

type (
	// PlatMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatMainModel.
	PlatMainModel interface {
		platMainModel
		Field(field string) *defaultPlatMainModel
		Alias(alias string) *defaultPlatMainModel
		WhereStr(whereStr string) *defaultPlatMainModel
		WhereId(id int) *defaultPlatMainModel
		WhereMap(whereMap map[string]any) *defaultPlatMainModel
		WhereRaw(whereStr string, whereData []any) *defaultPlatMainModel
		Order(order string) *defaultPlatMainModel
		Plat(id int) *defaultPlatMainModel
		Find(ctx context.Context, id ...any) (*PlatMain, error)
		Page(ctx context.Context, page int, rows int) ([]*PlatMain, error)
		List(ctx context.Context) ([]*PlatMain, error)
		Count(ctx context.Context) int
		Inc(ctx context.Context, field string, num int) error
		Dec(ctx context.Context, field string, num int) error
	}

	customPlatMainModel struct {
		*defaultPlatMainModel
	}
)

// NewPlatMainModel returns a model for the database table.
func NewPlatMainModel(conn ...sqlx.SqlConn) PlatMainModel {
	if len(conn) > 0 {
		return &customPlatMainModel{
			defaultPlatMainModel: newPlatMainModel(conn[0]),
		}
	} else {
		fmt.Println(global.Config.DB)
		return &customPlatMainModel{
			defaultPlatMainModel: newPlatMainModel(sqlx.NewMysql(global.Config.DB.DataSource)),
		}

	}
}
func (m *defaultPlatMainModel) Find(ctx context.Context, id ...any) (*PlatMain, error) {
	var err error
	if err = m.err; err != nil {
		m.err = nil
		return nil, err
	}
	var resp PlatMain
	var sql string
	field := platMainRows
	if m.fieldSql != "" {
		field = m.fieldSql
	}
	if len(id) > 0 {
		sql = fmt.Sprintf("select %s from %s where id=? limit 1", field, m.table)
		err = m.conn.QueryRowPartialCtx(ctx, &resp, sql, id[0]) //QueryRowCtx 必须字段都覆盖
	} else {
		sql = fmt.Sprintf("select %s from %s %s where "+m.whereSql+" limit 1", field, m.table, m.aliasSql)
		err = m.conn.QueryRowPartialCtx(ctx, &resp, sql, m.whereData...)
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
func (m *defaultPlatMainModel) WhereId(id int) *defaultPlatMainModel {
	m.whereSql = "id=?"
	m.whereData = append(m.whereData, id)
	return m
}
func (m *defaultPlatMainModel) Page(ctx context.Context, page int, rows int) ([]*PlatMain, error) {

	return nil, nil
}
func (m *defaultPlatMainModel) List(ctx context.Context) ([]*PlatMain, error) {

	return nil, nil
}
func (m *defaultPlatMainModel) WhereStr(whereStr string) *defaultPlatMainModel {
	return m
}

func (m *defaultPlatMainModel) WhereMap(whereMap map[string]any) *defaultPlatMainModel {
	return m
}
func (m *defaultPlatMainModel) WhereRaw(whereStr string, whereData []any) *defaultPlatMainModel {
	if m.whereSql != "" {
		m.whereSql += " AND (" + whereStr + ")"
	} else {
		m.whereSql = "(" + whereStr + ")"
	}
	m.whereData = append(m.whereData, whereData...)
	return m
}
func (m *defaultPlatMainModel) Alias(field string) *defaultPlatMainModel {
	m.aliasSql = field
	return m
}
func (m *defaultPlatMainModel) Field(field string) *defaultPlatMainModel {
	m.fieldSql = field
	return m
}
func (m *defaultPlatMainModel) Order(order string) *defaultPlatMainModel {

	return m
}
func (m *defaultPlatMainModel) Count(ctx context.Context) int {

	return 0
}
func (m *defaultPlatMainModel) Inc(ctx context.Context, field string, num int) error {

	return nil
}
func (m *defaultPlatMainModel) Dec(ctx context.Context, field string, num int) error {

	return nil
}
func (m *defaultPlatMainModel) Plat(id int) *defaultPlatMainModel {

	return nil
}
