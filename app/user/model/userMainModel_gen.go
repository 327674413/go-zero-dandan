// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/redisd"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userMainFieldNames          = builder.RawFieldNames(&UserMain{})
	userMainRows                = strings.Join(userMainFieldNames, ",")
	defaultUserMainFields       = strings.Join(userMainFieldNames, ",")
	userMainRowsExpectAutoSet   = strings.Join(stringx.Remove(userMainFieldNames, "`delete_at`"), ",")
	userMainRowsWithPlaceHolder = strings.Join(stringx.Remove(userMainFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

type (
	userMainModel interface {
		Insert(data map[string]string) (int64, error)
		TxInsert(tx *sql.Tx, data map[string]string) (int64, error)
		Update(data map[string]string) (int64, error)
		TxUpdate(tx *sql.Tx, data map[string]string) (int64, error)
		Save(data map[string]string) (int64, error)
		TxSave(tx *sql.Tx, data map[string]string) (int64, error)
		Delete(ctx context.Context, id int64) error
		Field(field string) *defaultUserMainModel
		Alias(alias string) *defaultUserMainModel
		Where(whereStr string, whereData ...any) *defaultUserMainModel
		WhereId(id int64) *defaultUserMainModel
		Order(order string) *defaultUserMainModel
		Plat(id int64) *defaultUserMainModel
		Find() (*UserMain, error)
		FindById(id int64) (*UserMain, error)
		CacheFind(redis *redisd.Redisd) (*UserMain, error)
		CacheFindById(redis *redisd.Redisd, id int64) (*UserMain, error)
		Page(page int64, rows int64) *defaultUserMainModel
		Select() ([]*UserMain, error)
		CacheSelect(redis *redisd.Redisd) ([]*UserMain, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultUserMainModel
		Reinit() *defaultUserMainModel
	}

	defaultUserMainModel struct {
		conn            sqlx.SqlConn
		table           string
		dao             *dao.SqlxDao
		softDeleteField string
		softDeletable   bool
		fieldSql        string
		whereSql        string
		aliasSql        string
		orderSql        string
		platId          int64
		whereData       []any
		err             error
		ctx             context.Context
	}

	UserMain struct {
		Id        int64  `db:"id"`
		UnionId   int64  `db:"union_id"`   // 平台层用户唯一表示
		StateEm   int64  `db:"state_em"`   // 用户状态枚举
		Account   string `db:"account"`    // 登录账号
		Password  string `db:"password"`   // 登录密码
		Code      string `db:"code"`       // 用户编号
		Nickname  string `db:"nickname"`   // 昵称
		Phone     string `db:"phone"`      // 手机号
		PhoneArea string `db:"phone_area"` // 手机区号
		Email     string `db:"email"`      // 邮箱地址
		Avatar    string `db:"avatar"`     // 头像
		SexEm     int64  `db:"sex_em"`     // 性别枚举
		PlatId    int64  `db:"plat_id"`    // 应用id
		CreateAt  int64  `db:"create_at"`  // 创建时间戳
		UpdateAt  int64  `db:"update_at"`  // 更新时间戳
		DeleteAt  int64  `db:"delete_at"`  // 删除时间戳
	}
)

func newUserMainModel(conn sqlx.SqlConn, platId int64) *defaultUserMainModel {
	dao := dao.NewSqlxDao(conn, "`user_main`", defaultUserMainFields, true, "delete_at")
	dao.Plat(platId)
	return &defaultUserMainModel{
		conn:            conn,
		dao:             dao,
		table:           "`user_main`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultUserMainModel) Ctx(ctx context.Context) *defaultUserMainModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultUserMainModel) WhereId(id int64) *defaultUserMainModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultUserMainModel) Where(whereStr string, whereData ...any) *defaultUserMainModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultUserMainModel) Alias(alias string) *defaultUserMainModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultUserMainModel) Field(field string) *defaultUserMainModel {
	m.dao.Field(field)
	return m
}
func (m *defaultUserMainModel) Order(order string) *defaultUserMainModel {
	m.dao.Order(order)
	return m
}
func (m *defaultUserMainModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultUserMainModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultUserMainModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultUserMainModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserMainModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserMainModel) Plat(id int64) *defaultUserMainModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultUserMainModel) Reinit() *defaultUserMainModel {
	m.dao.Reinit()
	return m
}

func (m *defaultUserMainModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserMainModel) Find() (*UserMain, error) {
	resp := &UserMain{}
	err := m.dao.Find(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) FindById(id int64) (*UserMain, error) {
	resp := &UserMain{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) CacheFind(redis *redisd.Redisd) (*UserMain, error) {
	resp := &UserMain{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) CacheFindById(redis *redisd.Redisd, id int64) (*UserMain, error) {
	resp := &UserMain{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserMainModel) Select() ([]*UserMain, error) {
	var resp []*UserMain
	err := m.dao.Select(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) CacheSelect(redis *redisd.Redisd) ([]*UserMain, error) {
	var resp []*UserMain
	err := m.dao.CacheSelect(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) Page(page int64, rows int64) *defaultUserMainModel {
	m.dao.Page(page, rows)
	return m
}

func (m *defaultUserMainModel) Insert(data map[string]string) (int64, error) {
	return m.dao.Insert(data)
}
func (m *defaultUserMainModel) TxInsert(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.TxInsert(tx, data)
}

func (m *defaultUserMainModel) Update(data map[string]string) (int64, error) {
	return m.dao.Update(data)
}
func (m *defaultUserMainModel) TxUpdate(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultUserMainModel) Save(data map[string]string) (int64, error) {
	return m.dao.Save(data)
}
func (m *defaultUserMainModel) TxSave(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.Save(data)
}

func (m *defaultUserMainModel) tableName() string {
	return m.table
}
