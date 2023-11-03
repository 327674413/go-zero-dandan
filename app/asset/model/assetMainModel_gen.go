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
	assetMainFieldNames          = builder.RawFieldNames(&AssetMain{})
	assetMainRows                = strings.Join(assetMainFieldNames, ",")
	defaultAssetMainFields       = strings.Join(assetMainFieldNames, ",")
	assetMainRowsExpectAutoSet   = strings.Join(stringx.Remove(assetMainFieldNames, "`delete_at`"), ",")
	assetMainRowsWithPlaceHolder = strings.Join(stringx.Remove(assetMainFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

type (
	assetMainModel interface {
		Insert(data map[string]string) (int64, error)
		TxInsert(tx *sql.Tx, data map[string]string) (int64, error)
		Update(data map[string]string) (int64, error)
		TxUpdate(tx *sql.Tx, data map[string]string) (int64, error)
		Save(data map[string]string) (int64, error)
		TxSave(tx *sql.Tx, data map[string]string) (int64, error)
		Delete(ctx context.Context, id int64) error
		Field(field string) *defaultAssetMainModel
		Alias(alias string) *defaultAssetMainModel
		Where(whereStr string, whereData ...any) *defaultAssetMainModel
		WhereId(id int64) *defaultAssetMainModel
		Order(order string) *defaultAssetMainModel
		Limit(num int64) *defaultAssetMainModel
		Plat(id int64) *defaultAssetMainModel
		Find() (*AssetMain, error)
		FindById(id int64) (*AssetMain, error)
		CacheFind(redis *redisd.Redisd) (*AssetMain, error)
		CacheFindById(redis *redisd.Redisd, id int64) (*AssetMain, error)
		Page(page int64, rows int64) *defaultAssetMainModel
		Select() ([]*AssetMain, error)
		SelectWithTotal() ([]*AssetMain, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*AssetMain, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultAssetMainModel
		Reinit() *defaultAssetMainModel
		Dao() *dao.SqlxDao
	}

	defaultAssetMainModel struct {
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

	AssetMain struct {
		Id       int64  `db:"id"`
		StateEm  int64  `db:"state_em"`  // 文件状态
		Sha1     string `db:"sha1"`      // 哈希值
		Name     string `db:"name"`      // 上传时的原始名称
		ModeEm   int64  `db:"mode_em"`   // 存储模式枚举
		Mime     string `db:"mime"`      // 文件类型
		SizeNum  int64  `db:"size_num"`  // 文件字节
		SizeText string `db:"size_text"` // 文件大小
		Ext      string `db:"ext"`       // 文件后缀
		Url      string `db:"url"`       // 文件链接
		Path     string `db:"path"`      // 存储路径
		CreateAt int64  `db:"create_at"` // 创建时间戳
		UpdateAt int64  `db:"update_at"` // 更新时间戳
		DeleteAt int64  `db:"delete_at"` // 删除时间戳
	}
)

func newAssetMainModel(conn sqlx.SqlConn, platId int64) *defaultAssetMainModel {
	dao := dao.NewSqlxDao(conn, "`asset_main`", defaultAssetMainFields, true, "delete_at")
	dao.Plat(platId)
	return &defaultAssetMainModel{
		conn:            conn,
		dao:             dao,
		table:           "`asset_main`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultAssetMainModel) Ctx(ctx context.Context) *defaultAssetMainModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultAssetMainModel) WhereId(id int64) *defaultAssetMainModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultAssetMainModel) Where(whereStr string, whereData ...any) *defaultAssetMainModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultAssetMainModel) Alias(alias string) *defaultAssetMainModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultAssetMainModel) Field(field string) *defaultAssetMainModel {
	m.dao.Field(field)
	return m
}
func (m *defaultAssetMainModel) Order(order string) *defaultAssetMainModel {
	m.dao.Order(order)
	return m
}
func (m *defaultAssetMainModel) Limit(num int64) *defaultAssetMainModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultAssetMainModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultAssetMainModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultAssetMainModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultAssetMainModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultAssetMainModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultAssetMainModel) Plat(id int64) *defaultAssetMainModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultAssetMainModel) Reinit() *defaultAssetMainModel {
	m.dao.Reinit()
	return m
}
func (m *defaultAssetMainModel) Dao() *dao.SqlDao {
	return m.dao
}
func (m *defaultAssetMainModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAssetMainModel) Find() (*AssetMain, error) {
	resp := &AssetMain{}
	err := m.dao.Find(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetMainModel) FindById(id int64) (*AssetMain, error) {
	resp := &AssetMain{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetMainModel) CacheFind(redis *redisd.Redisd) (*AssetMain, error) {
	resp := &AssetMain{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetMainModel) CacheFindById(redis *redisd.Redisd, id int64) (*AssetMain, error) {
	resp := &AssetMain{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAssetMainModel) Select() ([]*AssetMain, error) {
	resp := make([]*AssetMain, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetMainModel) SelectWithTotal() ([]*AssetMain, int64, error) {
	resp := make([]*AssetMain, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultAssetMainModel) CacheSelect(redis *redisd.Redisd) ([]*AssetMain, error) {
	resp := make([]*AssetMain, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAssetMainModel) Page(page int64, rows int64) *defaultAssetMainModel {
	m.dao.Page(page, rows)
	return m
}

func (m *defaultAssetMainModel) Insert(data map[string]string) (int64, error) {
	return m.dao.Insert(data)
}
func (m *defaultAssetMainModel) TxInsert(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.TxInsert(tx, data)
}

func (m *defaultAssetMainModel) Update(data map[string]string) (int64, error) {
	return m.dao.Update(data)
}
func (m *defaultAssetMainModel) TxUpdate(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultAssetMainModel) Save(data map[string]string) (int64, error) {
	return m.dao.Save(data)
}
func (m *defaultAssetMainModel) TxSave(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.Save(data)
}

func (m *defaultAssetMainModel) tableName() string {
	return m.table
}
