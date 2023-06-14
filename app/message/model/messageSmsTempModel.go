package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/common/redisd"
	"strconv"
)

var _ MessageSmsTempModel = (*customMessageSmsTempModel)(nil)

type (
	// MessageSmsTempModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageSmsTempModel.
	MessageSmsTempModel interface {
		messageSmsTempModel
		Field(field string) *defaultMessageSmsTempModel
		Alias(alias string) *defaultMessageSmsTempModel
		WhereStr(whereStr string) *defaultMessageSmsTempModel
		WhereId(id int) *defaultMessageSmsTempModel
		WhereMap(whereMap map[string]any) *defaultMessageSmsTempModel
		WhereRaw(whereStr string, whereData []any) *defaultMessageSmsTempModel
		Order(order string) *defaultMessageSmsTempModel
		Plat(id int) *defaultMessageSmsTempModel
		Find(ctx context.Context, id ...any) (*MessageSmsTemp, error)
		CacheFind(ctx context.Context, redis *redisd.Redisd, id ...int64) (*MessageSmsTemp, error)
		Page(ctx context.Context, page int, rows int) ([]*MessageSmsTemp, error)
		List(ctx context.Context) ([]*MessageSmsTemp, error)
		Count(ctx context.Context) int
		Inc(ctx context.Context, field string, num int) error
		Dec(ctx context.Context, field string, num int) error
	}

	customMessageSmsTempModel struct {
		*defaultMessageSmsTempModel
	}
)

// NewMessageSmsTempModel returns a model for the database table.
func NewMessageSmsTempModel(conn sqlx.SqlConn) MessageSmsTempModel {
	return &customMessageSmsTempModel{
		defaultMessageSmsTempModel: newMessageSmsTempModel(conn),
	}
}
func (m *defaultMessageSmsTempModel) Find(ctx context.Context, id ...any) (*MessageSmsTemp, error) {
	var err error
	if err = m.err; err != nil {
		m.err = nil
		return nil, err
	}
	var resp MessageSmsTemp
	var sql string
	field := messageSmsTempRows
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
		return &resp, nil
	default:
		return nil, err
	}
}
func (m *defaultMessageSmsTempModel) WhereId(id int) *defaultMessageSmsTempModel {
	m.whereSql = "id=?"
	m.whereData = append(m.whereData, id)
	return m
}
func (m *defaultMessageSmsTempModel) Page(ctx context.Context, page int, rows int) ([]*MessageSmsTemp, error) {

	return nil, nil
}
func (m *defaultMessageSmsTempModel) List(ctx context.Context) ([]*MessageSmsTemp, error) {

	return nil, nil
}
func (m *defaultMessageSmsTempModel) WhereStr(whereStr string) *defaultMessageSmsTempModel {
	return m
}

func (m *defaultMessageSmsTempModel) WhereMap(whereMap map[string]any) *defaultMessageSmsTempModel {
	return m
}
func (m *defaultMessageSmsTempModel) WhereRaw(whereStr string, whereData []any) *defaultMessageSmsTempModel {
	if m.whereSql != "" {
		m.whereSql += " AND (" + whereStr + ")"
	} else {
		m.whereSql = "(" + whereStr + ")"
	}
	m.whereData = append(m.whereData, whereData...)
	return m
}
func (m *defaultMessageSmsTempModel) Alias(field string) *defaultMessageSmsTempModel {
	m.aliasSql = field
	return m
}
func (m *defaultMessageSmsTempModel) Field(field string) *defaultMessageSmsTempModel {
	m.fieldSql = field
	return m
}
func (m *defaultMessageSmsTempModel) Order(order string) *defaultMessageSmsTempModel {

	return m
}
func (m *defaultMessageSmsTempModel) Count(ctx context.Context) int {

	return 0
}
func (m *defaultMessageSmsTempModel) Inc(ctx context.Context, field string, num int) error {

	return nil
}
func (m *defaultMessageSmsTempModel) Dec(ctx context.Context, field string, num int) error {

	return nil
}
func (m *defaultMessageSmsTempModel) Plat(id int) *defaultMessageSmsTempModel {

	return nil
}
func (m *defaultMessageSmsTempModel) CacheFind(ctx context.Context, redis *redisd.Redisd, id ...int64) (*MessageSmsTemp, error) {
	resp := &MessageSmsTemp{}
	cacheField := "model_" + m.tableName()
	cacheKey := strconv.FormatInt(id[0], 10)
	// todo::需要把where条件一起放进去作为key，这样就能支持更多的自动缓存查询
	err := redis.GetData(cacheField, cacheKey, resp)
	if err == nil {
		return resp, nil
	}
	resp, err = m.Find(ctx, id[0])
	fmt.Println(resp, err)
	if err != nil {
		return resp, err
	}
	if resp.Id != 0 {
		redis.SetData(cacheField, cacheKey, resp)
	}
	return resp, nil
}
