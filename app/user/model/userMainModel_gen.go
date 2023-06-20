// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"go-zero-dandan/common/redisd"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userMainFieldNames          = builder.RawFieldNames(&UserMain{})
	userMainRows                = strings.Join(userMainFieldNames, ",")
	userMainRowsExpectAutoSet   = strings.Join(stringx.Remove(userMainFieldNames, "`delete_at`"), ",")
	userMainRowsWithPlaceHolder = strings.Join(stringx.Remove(userMainFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

type (
	userMainModel interface {
		Insert(ctx context.Context, data *UserMain) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserMain, error)
		Update(ctx context.Context, data *UserMain) error
		Delete(ctx context.Context, id int64) error
		Field(field string) *defaultUserMainModel
		Alias(alias string) *defaultUserMainModel
		WhereStr(whereStr string) *defaultUserMainModel
		WhereId(id int) *defaultUserMainModel
		WhereMap(whereMap map[string]any) *defaultUserMainModel
		WhereRaw(whereStr string, whereData []any) *defaultUserMainModel
		Order(order string) *defaultUserMainModel
		Plat(id int) *defaultUserMainModel
		Find(ctx context.Context, id ...any) (*UserMain, error)
		CacheFind(ctx context.Context, redis *redisd.Redisd, id ...int64) (*UserMain, error)
		Page(ctx context.Context, page int, rows int) ([]*UserMain, error)
		List(ctx context.Context) ([]*UserMain, error)
		Count(ctx context.Context) int
		Inc(ctx context.Context, field string, num int) error
		Dec(ctx context.Context, field string, num int) error
	}

	defaultUserMainModel struct {
		conn            sqlx.SqlConn
		table           string
		softDeleteField string
		softDeletable   bool
		fieldSql        string
		whereSql        string
		aliasSql        string
		orderSql        string
		platId          int64
		whereData       []any
		err             error
	}

	UserMain struct {
		Id          int64  `db:"id"`
		UserUnionId int64  `db:"user_union_id"` // 平台层用户唯一表示
		StateEm     int64  `db:"state_em"`      // 用户状态枚举
		Account     string `db:"account"`       // 登录账号
		Password    string `db:"password"`      // 登录密码
		Uid         string `db:"uid"`           // 用户编号
		Nickname    string `db:"nickname"`      // 昵称
		Phone       string `db:"phone"`         // 手机号
		PhoneArea   string `db:"phone_area"`    // 手机区号
		Email       string `db:"email"`         // 邮箱地址
		Avatar      string `db:"avatar"`        // 头像
		SexEm       int64  `db:"sex_em"`        // 性别枚举
		PlatId      int64  `db:"plat_id"`       // 应用id
		CreateAt    int64  `db:"create_at"`     // 创建时间戳
		UpdateAt    int64  `db:"update_at"`     // 更新时间戳
		DeleteAt    int64  `db:"delete_at"`     // 删除时间戳
	}
)

func newUserMainModel(conn sqlx.SqlConn, platId int64) *defaultUserMainModel {
	return &defaultUserMainModel{
		conn:            conn,
		table:           "`user_main`",
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
		platId:          platId,
	}
}

func (m *defaultUserMainModel) WhereId(id int) *defaultUserMainModel {
	m.whereSql = "id=?"
	m.whereData = append(m.whereData, id)
	return m
}

func (m *defaultUserMainModel) WhereStr(whereStr string) *defaultUserMainModel {
	if m.whereSql != "" {
		m.whereSql += " AND (" + whereStr + ")"
	} else {
		m.whereSql = "(" + whereStr + ")"
	}
	return m
}

func (m *defaultUserMainModel) WhereMap(whereMap map[string]any) *defaultUserMainModel {
	return m
}
func (m *defaultUserMainModel) WhereRaw(whereStr string, whereData []any) *defaultUserMainModel {
	if m.whereSql != "" {
		m.whereSql += " AND (" + whereStr + ")"
	} else {
		m.whereSql = "(" + whereStr + ")"
	}
	m.whereData = append(m.whereData, whereData...)
	return m
}
func (m *defaultUserMainModel) Alias(field string) *defaultUserMainModel {
	m.aliasSql = field
	return m
}
func (m *defaultUserMainModel) Field(field string) *defaultUserMainModel {
	m.fieldSql = field
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
func (m *defaultUserMainModel) CacheFind(ctx context.Context, redis *redisd.Redisd, id ...int64) (*UserMain, error) {
	resp := &UserMain{}
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

func (m *defaultUserMainModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserMainModel) Find(ctx context.Context, id ...any) (*UserMain, error) {
	defer m.Reinit()
	var err error
	if err = m.err; err != nil {
		m.err = nil
		return nil, err
	}
	var resp UserMain
	var sql string
	field := userMainRows
	if m.fieldSql != "" {
		field = m.fieldSql
	}
	if len(id) > 0 {
		if m.whereSql == "" {
			m.whereSql = "1=1"
		}
		if m.platId != 0 {
			m.whereSql = m.whereSql + fmt.Sprintf(" AND id=%d AND plat_id=%d", id[0], m.platId)
		} else {
			m.whereSql = m.whereSql + fmt.Sprintf(" AND id=%d", id[0])
		}
		sql = fmt.Sprintf("select %s from %s where %s limit 1", field, m.table, m.whereSql)
		err = m.conn.QueryRowPartialCtx(ctx, &resp, sql) //QueryRowCtx 必须字段都覆盖
	} else {
		if m.whereSql == "" {
			m.whereSql = "1=1"
		}
		if m.platId != 0 {
			m.whereSql = m.whereSql + " AND plat_id=" + fmt.Sprintf("%d", m.platId)
		}
		sql = fmt.Sprintf("select %s from %s %s where "+m.whereSql+" limit 1", field, m.table, m.aliasSql)
		err = m.conn.QueryRowPartialCtx(ctx, &resp, sql, m.whereData...)
	}
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return &resp, nil
	default:
		return nil, err
	}
}
func (m *defaultUserMainModel) List(ctx context.Context) ([]*UserMain, error) {
	defer m.Reinit()
	var err error
	if err = m.err; err != nil {
		m.err = nil
		return nil, err
	}
	var resp []*UserMain
	var sql string
	field := userMainRows
	if m.fieldSql != "" {
		field = m.fieldSql
	}
	if m.whereSql == "" {
		m.whereSql = "1=1"
	}
	sql = fmt.Sprintf("select %s from %s %s where "+m.whereSql, field, m.table, m.aliasSql)
	err = m.conn.QueryRowsPartialCtx(ctx, &resp, sql, m.whereData...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return resp, nil
	default:
		return nil, err
	}
}
func (m *defaultUserMainModel) Page(ctx context.Context, page int, rows int) ([]*UserMain, error) {
	defer m.Reinit()
	return nil, nil
}
func (m *defaultUserMainModel) FindOne(ctx context.Context, id int64) (*UserMain, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userMainRows, m.table)
	var resp UserMain
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultUserMainModel) Reinit() {
	m.whereSql = ""
	m.fieldSql = ""
	m.aliasSql = ""
	m.orderSql = ""
	m.whereData = make([]any, 0)
	m.err = nil
}

func (m *defaultUserMainModel) Insert(ctx context.Context, data *UserMain) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userMainRowsExpectAutoSet)
	data.CreateAt = time.Now().Unix()
	data.UpdateAt = time.Now().Unix()
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.UserUnionId, data.StateEm, data.Account, data.Password, data.Uid, data.Nickname, data.Phone, data.PhoneArea, data.Email, data.Avatar, data.SexEm, data.PlatId, data.CreateAt, data.UpdateAt)
	return ret, err
}

func (m *defaultUserMainModel) Update(ctx context.Context, data *UserMain) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userMainRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserUnionId, data.StateEm, data.Account, data.Password, data.Uid, data.Nickname, data.Phone, data.PhoneArea, data.Email, data.Avatar, data.SexEm, data.PlatId, data.CreateAt, data.UpdateAt, data.Id)
	return err
}

func (m *defaultUserMainModel) tableName() string {
	return m.table
}
