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
	userUnionFieldNames          = builder.RawFieldNames(&UserUnion{})
	userUnionRows                = strings.Join(userUnionFieldNames, ",")
	defaultUserUnionFields       = strings.Join(userUnionFieldNames, ",")
	userUnionRowsExpectAutoSet   = strings.Join(stringx.Remove(userUnionFieldNames, "`delete_at`"), ",")
	userUnionRowsWithPlaceHolder = strings.Join(stringx.Remove(userUnionFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	UserUnion_Id       dao.TableField = "id"
	UserUnion_CreateAt dao.TableField = "create_at"
	UserUnion_UpdateAt dao.TableField = "update_at"
	UserUnion_DeleteAt dao.TableField = "delete_at"
)

type (
	userUnionModel interface {
		Delete(id ...string) (effectRow int64, danErr error)
		TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error)
		Insert(data *UserUnion) (effectRow int64, danErr error)
		TxInsert(tx *sql.Tx, data *UserUnion) (effectRow int64, danErr error)
		Update(data map[dao.TableField]any) (effectRow int64, danErr error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, danErr error)
		Save(data *UserUnion) (effectRow int64, danErr error)
		TxSave(tx *sql.Tx, data *UserUnion) (effectRow int64, danErr error)
		Field(field string) *defaultUserUnionModel
		Except(fields ...string) *defaultUserUnionModel
		Alias(alias string) *defaultUserUnionModel
		Where(whereStr string, whereData ...any) *defaultUserUnionModel
		WhereId(id string) *defaultUserUnionModel
		Order(order string) *defaultUserUnionModel
		Limit(num int64) *defaultUserUnionModel
		Plat(id string) *defaultUserUnionModel
		Find() (*UserUnion, error)
		FindById(id string) (data *UserUnion, danErr error)
		CacheFind(redis *redisd.Redisd) (data *UserUnion, danErr error)
		CacheFindById(redis *redisd.Redisd, id string) (data *UserUnion, danErr error)
		Page(page int64, rows int64) *defaultUserUnionModel
		Total() (total int64, danErr error)
		Select() (dataList []*UserUnion, danErr error)
		SelectWithTotal() (dataList []*UserUnion, total int64, danErr error)
		CacheSelect(redis *redisd.Redisd) (dataList []*UserUnion, danErr error)
		Count() (total int64, danErr error)
		Inc(field string, num int) (effectRow int64, danErr error)
		Dec(field string, num int) (effectRow int64, danErr error)
		StartTrans() (tx *sql.Tx, danErr error)
		Commit(tx *sql.Tx) (danErr error)
		Rollback(tx *sql.Tx) (danErr error)
		Ctx(ctx context.Context) *defaultUserUnionModel
		Reinit() *defaultUserUnionModel
		Dao() *dao.SqlxDao
	}

	defaultUserUnionModel struct {
		conn            sqlx.SqlConn
		table           string
		dao             *dao.SqlxDao
		softDeleteField string
		softDeletable   bool
		fieldSql        string
		whereSql        string
		aliasSql        string
		orderSql        string
		platId          string
		whereData       []any
		err             error
		ctx             context.Context
	}

	UserUnion struct {
		Id       string `db:"id" json:"id"`
		CreateAt int64  `db:"create_at" json:"createAt"` // 创建时间戳
		UpdateAt int64  `db:"update_at" json:"updateAt"` // 更新时间戳
		DeleteAt int64  `db:"delete_at" json:"deleteAt"` // 删除时间戳
	}
)

// NewUserUnionModel returns a model for the database table.
func NewUserUnionModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) UserUnionModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customUserUnionModel{
		defaultUserUnionModel: newUserUnionModel(ctxOrNil, conn, platid),
		softDeletable:         softDeletableUserUnion,
	}
}
func newUserUnionModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultUserUnionModel {
	dao := dao.NewSqlxDao(conn, "`user_union`", defaultUserUnionFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultUserUnionModel{
		ctx:             ctx,
		conn:            conn,
		dao:             dao,
		table:           "`user_union`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultUserUnionModel) Ctx(ctx context.Context) *defaultUserUnionModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultUserUnionModel) WhereId(id string) *defaultUserUnionModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultUserUnionModel) Where(whereStr string, whereData ...any) *defaultUserUnionModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultUserUnionModel) Alias(alias string) *defaultUserUnionModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultUserUnionModel) Field(field string) *defaultUserUnionModel {
	m.dao.Field(field)
	return m
}
func (m *defaultUserUnionModel) Except(fields ...string) *defaultUserUnionModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultUserUnionModel) Order(order string) *defaultUserUnionModel {
	m.dao.Order(order)
	return m
}
func (m *defaultUserUnionModel) Limit(num int64) *defaultUserUnionModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultUserUnionModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultUserUnionModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultUserUnionModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultUserUnionModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserUnionModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserUnionModel) Plat(id string) *defaultUserUnionModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultUserUnionModel) Reinit() *defaultUserUnionModel {
	m.dao.Reinit()
	return m
}
func (m *defaultUserUnionModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultUserUnionModel) Find() (*UserUnion, error) {
	resp := &UserUnion{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserUnionModel) FindById(id string) (*UserUnion, error) {
	resp := &UserUnion{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserUnionModel) CacheFind(redis *redisd.Redisd) (*UserUnion, error) {
	resp := &UserUnion{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserUnionModel) CacheFindById(redis *redisd.Redisd, id string) (*UserUnion, error) {
	resp := &UserUnion{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserUnionModel) Total() (total int64, danErr error) {
	return m.dao.Total()
}
func (m *defaultUserUnionModel) Select() ([]*UserUnion, error) {
	resp := make([]*UserUnion, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserUnionModel) SelectWithTotal() ([]*UserUnion, int64, error) {
	resp := make([]*UserUnion, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultUserUnionModel) CacheSelect(redis *redisd.Redisd) ([]*UserUnion, error) {
	resp := make([]*UserUnion, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserUnionModel) Page(page int64, size int64) *defaultUserUnionModel {
	m.dao.Page(page, size)
	return m
}
func (m *defaultUserUnionModel) Insert(data *UserUnion) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultUserUnionModel) TxInsert(tx *sql.Tx, data *UserUnion) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}
func (m *defaultUserUnionModel) Delete(id ...string) (effectRow int64, danErr error) {
	return m.dao.Delete(id...)
}
func (m *defaultUserUnionModel) TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error) {
	return m.dao.TxDelete(tx, id...)
}
func (m *defaultUserUnionModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultUserUnionModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultUserUnionModel) Save(data *UserUnion) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultUserUnionModel) TxSave(tx *sql.Tx, data *UserUnion) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultUserUnionModel) StartTrans() (tx *sql.Tx, danErr error) {
	return dao.StartTrans(m.conn, m.ctx)
}
func (m *defaultUserUnionModel) Commit(tx *sql.Tx) (danErr error) {
	return dao.Commit(tx)
}
func (m *defaultUserUnionModel) Rollback(tx *sql.Tx) (danErr error) {
	return dao.Rollback(tx)
}
func (m *defaultUserUnionModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultUserUnionModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
