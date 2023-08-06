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
	assetNetdiskFileFieldNames          = builder.RawFieldNames(&AssetNetdiskFile{})
	assetNetdiskFileRows                = strings.Join(assetNetdiskFileFieldNames, ",")
	defaultAssetNetdiskFileFields       = strings.Join(assetNetdiskFileFieldNames, ",")
	assetNetdiskFileRowsExpectAutoSet   = strings.Join(stringx.Remove(assetNetdiskFileFieldNames, "`delete_at`"), ",")
	assetNetdiskFileRowsWithPlaceHolder = strings.Join(stringx.Remove(assetNetdiskFileFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

type (
	assetNetdiskFileModel interface {
		Insert(data map[string]string) (int64, error)
		TxInsert(tx *sql.Tx, data map[string]string) (int64, error)
		Update(data map[string]string) (int64, error)
		TxUpdate(tx *sql.Tx, data map[string]string) (int64, error)
		Save(data map[string]string) (int64, error)
		TxSave(tx *sql.Tx, data map[string]string) (int64, error)
		Delete(ctx context.Context, id int64) error
		Field(field string) *defaultAssetNetdiskFileModel
		Alias(alias string) *defaultAssetNetdiskFileModel
		Where(whereStr string, whereData ...any) *defaultAssetNetdiskFileModel
		WhereId(id int64) *defaultAssetNetdiskFileModel
		Order(order string) *defaultAssetNetdiskFileModel
		Plat(id int64) *defaultAssetNetdiskFileModel
		Find() (*AssetNetdiskFile, error)
		FindById(id int64) (*AssetNetdiskFile, error)
		CacheFind(redis *redisd.Redisd) (*AssetNetdiskFile, error)
		CacheFindById(redis *redisd.Redisd, id int64) (*AssetNetdiskFile, error)
		Page(page int64, rows int64) *defaultAssetNetdiskFileModel
		Select() ([]*AssetNetdiskFile, error)
		CacheSelect(redis *redisd.Redisd) ([]*AssetNetdiskFile, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultAssetNetdiskFileModel
		Reinit() *defaultAssetNetdiskFileModel
	}

	defaultAssetNetdiskFileModel struct {
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

	AssetNetdiskFile struct {
		Id            int64  `db:"id"`
		AssetId       int64  `db:"asset_id"`        // 关联资源id
		Name          string `db:"name"`            // 自定义文件别名
		StateEm       int64  `db:"state_em"`        // 文件状态
		ClasId        int64  `db:"clas_id"`         // 所属目录id
		AuthGroupId   int64  `db:"auth_group_id"`   // 权限组别id
		AuthGroupName string `db:"auth_group_name"` // 权限组别名称
		ClasName      string `db:"clas_name"`       // 所属目录名称
		Sha1          string `db:"sha1"`            // 文件哈希
		OriginalName  string `db:"original_name"`   // 上传时的原始名称
		ModeEm        int64  `db:"mode_em"`         // 存储模式枚举
		Mime          string `db:"mime"`            // 文件类型
		SizeNum       int64  `db:"size_num"`        // 文件字节
		SizeText      string `db:"size_text"`       // 文件大小
		Ext           string `db:"ext"`             // 文件后缀
		Path          string `db:"path"`            // 文件路径
		Url           string `db:"url"`             // 文件链接
		UserId        int64  `db:"user_id"`         // 上传用户标识
		FinishAt      int64  `db:"finish_at"`       // 完成上传时间
		PlatId        int64  `db:"plat_id"`         // 应用id
		CreateAt      int64  `db:"create_at"`       // 创建时间戳
		UpdateAt      int64  `db:"update_at"`       // 更新时间戳
		DeleteAt      int64  `db:"delete_at"`       // 删除时间戳
	}
)

func newAssetNetdiskFileModel(conn sqlx.SqlConn, platId int64) *defaultAssetNetdiskFileModel {
	dao := dao.NewSqlxDao(conn, "`asset_netdisk_file`", defaultAssetNetdiskFileFields, true, "delete_at")
	dao.Plat(platId)
	return &defaultAssetNetdiskFileModel{
		conn:            conn,
		dao:             dao,
		table:           "`asset_netdisk_file`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultAssetNetdiskFileModel) Ctx(ctx context.Context) *defaultAssetNetdiskFileModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultAssetNetdiskFileModel) WhereId(id int64) *defaultAssetNetdiskFileModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultAssetNetdiskFileModel) Where(whereStr string, whereData ...any) *defaultAssetNetdiskFileModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultAssetNetdiskFileModel) Alias(alias string) *defaultAssetNetdiskFileModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultAssetNetdiskFileModel) Field(field string) *defaultAssetNetdiskFileModel {
	m.dao.Field(field)
	return m
}
func (m *defaultAssetNetdiskFileModel) Order(order string) *defaultAssetNetdiskFileModel {
	m.dao.Order(order)
	return m
}
func (m *defaultAssetNetdiskFileModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultAssetNetdiskFileModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultAssetNetdiskFileModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultAssetNetdiskFileModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultAssetNetdiskFileModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultAssetNetdiskFileModel) Plat(id int64) *defaultAssetNetdiskFileModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultAssetNetdiskFileModel) Reinit() *defaultAssetNetdiskFileModel {
	m.dao.Reinit()
	return m
}

func (m *defaultAssetNetdiskFileModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAssetNetdiskFileModel) Find() (*AssetNetdiskFile, error) {
	resp := &AssetNetdiskFile{}
	err := m.dao.Find(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetNetdiskFileModel) FindById(id int64) (*AssetNetdiskFile, error) {
	resp := &AssetNetdiskFile{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetNetdiskFileModel) CacheFind(redis *redisd.Redisd) (*AssetNetdiskFile, error) {
	resp := &AssetNetdiskFile{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetNetdiskFileModel) CacheFindById(redis *redisd.Redisd, id int64) (*AssetNetdiskFile, error) {
	resp := &AssetNetdiskFile{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAssetNetdiskFileModel) Select() ([]*AssetNetdiskFile, error) {
	var resp []*AssetNetdiskFile
	err := m.dao.Select(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetNetdiskFileModel) CacheSelect(redis *redisd.Redisd) ([]*AssetNetdiskFile, error) {
	var resp []*AssetNetdiskFile
	err := m.dao.CacheSelect(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetNetdiskFileModel) Page(page int64, rows int64) *defaultAssetNetdiskFileModel {
	m.dao.Page(page, rows)
	return m
}

func (m *defaultAssetNetdiskFileModel) Insert(data map[string]string) (int64, error) {
	return m.dao.Insert(data)
}
func (m *defaultAssetNetdiskFileModel) TxInsert(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.TxInsert(tx, data)
}

func (m *defaultAssetNetdiskFileModel) Update(data map[string]string) (int64, error) {
	return m.dao.Update(data)
}
func (m *defaultAssetNetdiskFileModel) TxUpdate(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultAssetNetdiskFileModel) Save(data map[string]string) (int64, error) {
	return m.dao.Save(data)
}
func (m *defaultAssetNetdiskFileModel) TxSave(tx *sql.Tx, data map[string]string) (int64, error) {
	return m.dao.Save(data)
}

func (m *defaultAssetNetdiskFileModel) tableName() string {
	return m.table
}