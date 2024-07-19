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
	assetMainFieldNames          = builder.RawFieldNames(&AssetMain{})
	assetMainRows                = strings.Join(assetMainFieldNames, ",")
	defaultAssetMainFields       = strings.Join(assetMainFieldNames, ",")
	assetMainRowsExpectAutoSet   = strings.Join(stringx.Remove(assetMainFieldNames, "`delete_at`"), ",")
	assetMainRowsWithPlaceHolder = strings.Join(stringx.Remove(assetMainFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	AssetMain_Id       dao.TableField = "id"
	AssetMain_StateEm  dao.TableField = "state_em"
	AssetMain_Sha1     dao.TableField = "sha1"
	AssetMain_Name     dao.TableField = "name"
	AssetMain_ModeEm   dao.TableField = "mode_em"
	AssetMain_Mime     dao.TableField = "mime"
	AssetMain_SizeNum  dao.TableField = "size_num"
	AssetMain_SizeText dao.TableField = "size_text"
	AssetMain_Ext      dao.TableField = "ext"
	AssetMain_Url      dao.TableField = "url"
	AssetMain_Path     dao.TableField = "path"
	AssetMain_PlatId   dao.TableField = "plat_id"
	AssetMain_CreateAt dao.TableField = "create_at"
	AssetMain_UpdateAt dao.TableField = "update_at"
	AssetMain_DeleteAt dao.TableField = "delete_at"
)

type (
	assetMainModel interface {
		Delete(id ...string) (effectRow int64, danErr error)
		TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error)
		Insert(data *AssetMain) (effectRow int64, danErr error)
		TxInsert(tx *sql.Tx, data *AssetMain) (effectRow int64, danErr error)
		Update(data map[dao.TableField]any) (effectRow int64, danErr error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, danErr error)
		Save(data *AssetMain) (effectRow int64, danErr error)
		TxSave(tx *sql.Tx, data *AssetMain) (effectRow int64, danErr error)
		Field(field string) *defaultAssetMainModel
		Except(fields ...string) *defaultAssetMainModel
		Alias(alias string) *defaultAssetMainModel
		Where(whereStr string, whereData ...any) *defaultAssetMainModel
		WhereId(id string) *defaultAssetMainModel
		Order(order string) *defaultAssetMainModel
		Limit(num int64) *defaultAssetMainModel
		Plat(id string) *defaultAssetMainModel
		Find() (*AssetMain, error)
		FindById(id string) (data *AssetMain, danErr error)
		CacheFind(redis *redisd.Redisd) (data *AssetMain, danErr error)
		CacheFindById(redis *redisd.Redisd, id string) (data *AssetMain, danErr error)
		Page(page int64, rows int64) *defaultAssetMainModel
		Total() (total int64, danErr error)
		Select() (dataList []*AssetMain, danErr error)
		SelectWithTotal() (dataList []*AssetMain, total int64, danErr error)
		CacheSelect(redis *redisd.Redisd) (dataList []*AssetMain, danErr error)
		Count() (total int64, danErr error)
		Inc(field string, num int) (effectRow int64, danErr error)
		Dec(field string, num int) (effectRow int64, danErr error)
		StartTrans() (tx *sql.Tx, danErr error)
		Commit(tx *sql.Tx) (danErr error)
		Rollback(tx *sql.Tx) (danErr error)
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
		platId          string
		whereData       []any
		err             error
		ctx             context.Context
	}

	AssetMain struct {
		Id       string `db:"id" json:"id"`
		StateEm  int64  `db:"state_em" json:"stateEm"`   // 文件状态
		Sha1     string `db:"sha1" json:"sha1"`          // 哈希值
		Name     string `db:"name" json:"name"`          // 上传时的原始名称
		ModeEm   int64  `db:"mode_em" json:"modeEm"`     // 存储模式枚举
		Mime     string `db:"mime" json:"mime"`          // 文件类型
		SizeNum  int64  `db:"size_num" json:"sizeNum"`   // 文件字节
		SizeText string `db:"size_text" json:"sizeText"` // 文件大小
		Ext      string `db:"ext" json:"ext"`            // 文件后缀
		Url      string `db:"url" json:"url"`            // 文件链接
		Path     string `db:"path" json:"path"`          // 存储路径
		PlatId   string `db:"plat_id" json:"platId"`     // 应用id
		CreateAt int64  `db:"create_at" json:"createAt"` // 创建时间戳
		UpdateAt int64  `db:"update_at" json:"updateAt"` // 更新时间戳
		DeleteAt int64  `db:"delete_at" json:"deleteAt"` // 删除时间戳
	}
)

// NewAssetMainModel returns a model for the database table.
func NewAssetMainModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) AssetMainModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customAssetMainModel{
		defaultAssetMainModel: newAssetMainModel(ctxOrNil, conn, platid),
		softDeletable:         softDeletableAssetMain,
	}
}
func newAssetMainModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultAssetMainModel {
	dao := dao.NewSqlxDao(conn, "`asset_main`", defaultAssetMainFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultAssetMainModel{
		ctx:             ctx,
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
func (m *defaultAssetMainModel) WhereId(id string) *defaultAssetMainModel {
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
func (m *defaultAssetMainModel) Except(fields ...string) *defaultAssetMainModel {
	m.dao.Except(fields...)
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
func (m *defaultAssetMainModel) Plat(id string) *defaultAssetMainModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultAssetMainModel) Reinit() *defaultAssetMainModel {
	m.dao.Reinit()
	return m
}
func (m *defaultAssetMainModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultAssetMainModel) Find() (*AssetMain, error) {
	resp := &AssetMain{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetMainModel) FindById(id string) (*AssetMain, error) {
	resp := &AssetMain{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
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
func (m *defaultAssetMainModel) CacheFindById(redis *redisd.Redisd, id string) (*AssetMain, error) {
	resp := &AssetMain{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetMainModel) Total() (total int64, danErr error) {
	return m.dao.Total()
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

func (m *defaultAssetMainModel) Page(page int64, size int64) *defaultAssetMainModel {
	m.dao.Page(page, size)
	return m
}
func (m *defaultAssetMainModel) Insert(data *AssetMain) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultAssetMainModel) TxInsert(tx *sql.Tx, data *AssetMain) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}
func (m *defaultAssetMainModel) Delete(id ...string) (effectRow int64, danErr error) {
	return m.dao.Delete(id...)
}
func (m *defaultAssetMainModel) TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error) {
	return m.dao.TxDelete(tx, id...)
}
func (m *defaultAssetMainModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultAssetMainModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultAssetMainModel) Save(data *AssetMain) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultAssetMainModel) TxSave(tx *sql.Tx, data *AssetMain) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultAssetMainModel) StartTrans() (tx *sql.Tx, danErr error) {
	return dao.StartTrans(m.conn, m.ctx)
}
func (m *defaultAssetMainModel) Commit(tx *sql.Tx) (danErr error) {
	return dao.Commit(tx)
}
func (m *defaultAssetMainModel) Rollback(tx *sql.Tx) (danErr error) {
	return dao.Rollback(tx)
}
func (m *defaultAssetMainModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultAssetMainModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
