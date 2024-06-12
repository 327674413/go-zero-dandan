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
	assetNetdiskFileFieldNames          = builder.RawFieldNames(&AssetNetdiskFile{})
	assetNetdiskFileRows                = strings.Join(assetNetdiskFileFieldNames, ",")
	defaultAssetNetdiskFileFields       = strings.Join(assetNetdiskFileFieldNames, ",")
	assetNetdiskFileRowsExpectAutoSet   = strings.Join(stringx.Remove(assetNetdiskFileFieldNames, "`delete_at`"), ",")
	assetNetdiskFileRowsWithPlaceHolder = strings.Join(stringx.Remove(assetNetdiskFileFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

type (
	assetNetdiskFileModel interface {
		Insert(data *AssetNetdiskFile) (int64, error)
		TxInsert(tx *sql.Tx, data *AssetNetdiskFile) (int64, error)
		Update(data map[string]any) (int64, error)
		TxUpdate(tx *sql.Tx, data map[string]any) (int64, error)
		Save(data *AssetNetdiskFile) (int64, error)
		TxSave(tx *sql.Tx, data *AssetNetdiskFile) (int64, error)
		Delete(ctx context.Context, id string) error
		Field(field string) *defaultAssetNetdiskFileModel
		Alias(alias string) *defaultAssetNetdiskFileModel
		Where(whereStr string, whereData ...any) *defaultAssetNetdiskFileModel
		WhereId(id string) *defaultAssetNetdiskFileModel
		Order(order string) *defaultAssetNetdiskFileModel
		Limit(num int64) *defaultAssetNetdiskFileModel
		Plat(id string) *defaultAssetNetdiskFileModel
		Find() (*AssetNetdiskFile, error)
		FindById(id string) (*AssetNetdiskFile, error)
		CacheFind(redis *redisd.Redisd) (*AssetNetdiskFile, error)
		CacheFindById(redis *redisd.Redisd, id string) (*AssetNetdiskFile, error)
		Page(page int64, rows int64) *defaultAssetNetdiskFileModel
		Select() ([]*AssetNetdiskFile, error)
		SelectWithTotal() ([]*AssetNetdiskFile, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*AssetNetdiskFile, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultAssetNetdiskFileModel
		Reinit() *defaultAssetNetdiskFileModel
		Dao() *dao.SqlxDao
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
		platId          string
		whereData       []any
		err             error
		ctx             context.Context
	}

	AssetNetdiskFile struct {
		Id            string `db:"id"`
		AssetId       string `db:"asset_id"`        // 关联资源id
		Name          string `db:"name"`            // 自定义文件别名
		StateEm       int64  `db:"state_em"`        // 文件状态
		ClasId        string `db:"clas_id"`         // 所属目录id
		AuthGroupId   string `db:"auth_group_id"`   // 权限组别id
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
		UserId        string `db:"user_id"`         // 上传用户标识
		FinishAt      int64  `db:"finish_at"`       // 完成上传时间
		PlatId        string `db:"plat_id"`         // 应用id
		CreateAt      int64  `db:"create_at"`       // 创建时间戳
		UpdateAt      int64  `db:"update_at"`       // 更新时间戳
		DeleteAt      int64  `db:"delete_at"`       // 删除时间戳
	}
)

func newAssetNetdiskFileModel(conn sqlx.SqlConn, platId string) *defaultAssetNetdiskFileModel {
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
func (m *defaultAssetNetdiskFileModel) WhereId(id string) *defaultAssetNetdiskFileModel {
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
func (m *defaultAssetNetdiskFileModel) Limit(num int64) *defaultAssetNetdiskFileModel {
	m.dao.Limit(num)
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
func (m *defaultAssetNetdiskFileModel) Plat(id string) *defaultAssetNetdiskFileModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultAssetNetdiskFileModel) Reinit() *defaultAssetNetdiskFileModel {
	m.dao.Reinit()
	return m
}
func (m *defaultAssetNetdiskFileModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultAssetNetdiskFileModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAssetNetdiskFileModel) Find() (*AssetNetdiskFile, error) {
	resp := &AssetNetdiskFile{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetNetdiskFileModel) FindById(id string) (*AssetNetdiskFile, error) {
	resp := &AssetNetdiskFile{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
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
func (m *defaultAssetNetdiskFileModel) CacheFindById(redis *redisd.Redisd, id string) (*AssetNetdiskFile, error) {
	resp := &AssetNetdiskFile{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAssetNetdiskFileModel) Select() ([]*AssetNetdiskFile, error) {
	resp := make([]*AssetNetdiskFile, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultAssetNetdiskFileModel) SelectWithTotal() ([]*AssetNetdiskFile, int64, error) {
	resp := make([]*AssetNetdiskFile, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultAssetNetdiskFileModel) CacheSelect(redis *redisd.Redisd) ([]*AssetNetdiskFile, error) {
	resp := make([]*AssetNetdiskFile, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAssetNetdiskFileModel) Page(page int64, size int64) *defaultAssetNetdiskFileModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultAssetNetdiskFileModel) Insert(data *AssetNetdiskFile) (int64, error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultAssetNetdiskFileModel) TxInsert(tx *sql.Tx, data *AssetNetdiskFile) (int64, error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultAssetNetdiskFileModel) Update(data map[string]any) (int64, error) {
	return m.dao.Update(data)
}
func (m *defaultAssetNetdiskFileModel) TxUpdate(tx *sql.Tx, data map[string]any) (int64, error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultAssetNetdiskFileModel) Save(data *AssetNetdiskFile) (int64, error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultAssetNetdiskFileModel) TxSave(tx *sql.Tx, data *AssetNetdiskFile) (int64, error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultAssetNetdiskFileModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultAssetNetdiskFileModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
