// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/redisd"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userInfoFieldNames          = builder.RawFieldNames(&UserInfo{})
	userInfoRows                = strings.Join(userInfoFieldNames, ",")
	defaultUserInfoFields       = strings.Join(userInfoFieldNames, ",")
	userInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(userInfoFieldNames, "`delete_at`"), ",")
	userInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

type (
	userInfoModel interface {
		Insert(data map[string]string) (int64, error)
		TxInsert(tx *sql.Tx, data map[string]string) (int64, error)
		Update(data map[string]string) (int64, error)
		TxUpdate(tx *sql.Tx, data map[string]string) (int64, error)
		Save(data map[string]string) (int64, error)
		TxSave(tx *sql.Tx, data map[string]string) (int64, error)
		Delete(ctx context.Context, id int64) error
		Field(field string) *defaultUserInfoModel
		Alias(alias string) *defaultUserInfoModel
		Where(whereStr string, whereData ...any) *defaultUserInfoModel
		WhereId(id int64) *defaultUserInfoModel
		Order(order string) *defaultUserInfoModel
		Plat(id int64) *defaultUserInfoModel
		Find() (*UserInfo, error)
		FindById(id int64) (*UserInfo, error)
		CacheFind(redis *redisd.Redisd) (*UserInfo, error)
		CacheFindById(redis *redisd.Redisd, id int64) (*UserInfo, error)
		Page(page int64, rows int64) *defaultUserInfoModel
		Select() ([]*UserInfo, error)
		SelectWithTotal() ([]*UserInfo, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*UserInfo, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultUserInfoModel
		Reinit() *defaultUserInfoModel
	}

	defaultUserInfoModel struct {
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

	UserInfo struct {
		Id           int64     `db:"id"`
		BirthDate    time.Time `db:"birth_date"`    // 出生日期
		GraduateFrom string    `db:"graduate_from"` // 毕业学校
		PlatId       int64     `db:"plat_id"`       // 应用id
		CreateAt     int64     `db:"create_at"`     // 创建时间戳
		UpdateAt     int64     `db:"update_at"`     // 更新时间戳
		DeleteAt     int64     `db:"delete_at"`     // 删除时间戳
	}
)

func newUserInfoModel(conn sqlx.SqlConn, platId int64) *defaultUserInfoModel {
	dao := dao.NewSqlxDao(conn, "`user_info`", defaultUserInfoFields, true, "delete_at")
	dao.Plat(platId)
	return &defaultUserInfoModel{
		conn:            conn,
		dao:             dao,
		table:           "`user_info`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultUserInfoModel) Ctx(ctx context.Context) *defaultUserInfoModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultUserInfoModel) WhereId(id int64) *defaultUserInfoModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultUserInfoModel) Where(whereStr string, whereData ...any) *defaultUserInfoModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultUserInfoModel) Alias(alias string) *defaultUserInfoModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultUserInfoModel) Field(field string) *defaultUserInfoModel {
	m.dao.Field(field)
	return m
}
func (m *defaultUserInfoModel) Order(order string) *defaultUserInfoModel {
	m.dao.Order(order)
	return m
}
func (m *defaultUserInfoModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultUserInfoModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultUserInfoModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultUserInfoModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserInfoModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserInfoModel) Plat(id int64) *defaultUserInfoModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultUserInfoModel) Reinit() *defaultUserInfoModel {
	m.dao.Reinit()
	return m
}

func (m *defaultUserInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserInfoModel) Find() (*UserInfo, error) {
	resp := &UserInfo{}
	err := m.dao.Find(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserInfoModel) FindById(id int64) (*UserInfo, error) {
	resp := &UserInfo{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserInfoModel) CacheFind(redis *redisd.Redisd) (*UserInfo, error) {
	resp := &UserInfo{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserInfoModel) CacheFindById(redis *redisd.Redisd, id int64) (*UserInfo, error) {
	resp := &UserInfo{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserInfoModel) Select() ([]*UserInfo, error) {
	resp := make([]*UserInfo, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserInfoModel) SelectWithTotal() ([]*UserInfo, int64, error) {
	resp := make([]*UserInfo, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultUserInfoModel) CacheSelect(redis *redisd.Redisd) ([]*UserInfo, error) {
	resp := make([]*UserInfo, 0)
	err := m.dao.CacheSelect(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserInfoModel) Page(page int64, rows int64) *defaultUserInfoModel {
	m.dao.Page(page, rows)
	return m
}

func (m *defaultUserInfoModel) Insert(data map[string]string) (int64, error) {
	return m.dao.Insert(data)
}
func (m *defaultUserInfoModel) TxInsert(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.TxInsert(tx, data)
}

func (m *defaultUserInfoModel) Update(data map[string]string) (int64, error) {
	return m.dao.Update(data)
}
func (m *defaultUserInfoModel) TxUpdate(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultUserInfoModel) Save(data map[string]string) (int64, error) {
	return m.dao.Save(data)
}
func (m *defaultUserInfoModel) TxSave(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.Save(data)
}

func (m *defaultUserInfoModel) tableName() string {
	return m.table
}
