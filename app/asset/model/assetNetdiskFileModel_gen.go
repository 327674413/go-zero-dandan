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

const (
	AssetNetdiskFile_Id            dao.TableField = "id"
	AssetNetdiskFile_AssetId       dao.TableField = "asset_id"
	AssetNetdiskFile_Name          dao.TableField = "name"
	AssetNetdiskFile_StateEm       dao.TableField = "state_em"
	AssetNetdiskFile_ClasId        dao.TableField = "clas_id"
	AssetNetdiskFile_AuthGroupId   dao.TableField = "auth_group_id"
	AssetNetdiskFile_AuthGroupName dao.TableField = "auth_group_name"
	AssetNetdiskFile_ClasName      dao.TableField = "clas_name"
	AssetNetdiskFile_Sha1          dao.TableField = "sha1"
	AssetNetdiskFile_OriginalName  dao.TableField = "original_name"
	AssetNetdiskFile_ModeEm        dao.TableField = "mode_em"
	AssetNetdiskFile_Mime          dao.TableField = "mime"
	AssetNetdiskFile_SizeNum       dao.TableField = "size_num"
	AssetNetdiskFile_SizeText      dao.TableField = "size_text"
	AssetNetdiskFile_Ext           dao.TableField = "ext"
	AssetNetdiskFile_Path          dao.TableField = "path"
	AssetNetdiskFile_Url           dao.TableField = "url"
	AssetNetdiskFile_UserId        dao.TableField = "user_id"
	AssetNetdiskFile_FinishAt      dao.TableField = "finish_at"
	AssetNetdiskFile_PlatId        dao.TableField = "plat_id"
	AssetNetdiskFile_CreateAt      dao.TableField = "create_at"
	AssetNetdiskFile_UpdateAt      dao.TableField = "update_at"
	AssetNetdiskFile_DeleteAt      dao.TableField = "delete_at"
)

type (
	assetNetdiskFileModel interface {
		Delete(id ...string) (effectRow int64, danErr error)
		TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error)
		Insert(data *AssetNetdiskFile) (effectRow int64, danErr error)
		TxInsert(tx *sql.Tx, data *AssetNetdiskFile) (effectRow int64, danErr error)
		Update(data map[dao.TableField]any) (effectRow int64, danErr error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, danErr error)
		Save(data *AssetNetdiskFile) (effectRow int64, danErr error)
		TxSave(tx *sql.Tx, data *AssetNetdiskFile) (effectRow int64, danErr error)
		Field(field string) *defaultAssetNetdiskFileModel
		Except(fields ...string) *defaultAssetNetdiskFileModel
		Alias(alias string) *defaultAssetNetdiskFileModel
		LeftJoin(joinTable string) *defaultAssetNetdiskFileModel
		RightJoin(joinTable string) *defaultAssetNetdiskFileModel
		InnerJoin(joinTable string) *defaultAssetNetdiskFileModel
		Where(whereStr string, whereData ...any) *defaultAssetNetdiskFileModel
		WhereId(id string) *defaultAssetNetdiskFileModel
		Order(order string) *defaultAssetNetdiskFileModel
		Limit(num int64) *defaultAssetNetdiskFileModel
		Plat(id string) *defaultAssetNetdiskFileModel
		Find() (*AssetNetdiskFile, error)
		FindById(id string) (data *AssetNetdiskFile, danErr error)
		CacheFind(redis *redisd.Redisd) (data *AssetNetdiskFile, danErr error)
		CacheFindById(redis *redisd.Redisd, id string) (data *AssetNetdiskFile, danErr error)
		Page(page int64, rows int64) *defaultAssetNetdiskFileModel
		Total() (total int64, danErr error)
		Select() (dataList []*AssetNetdiskFile, danErr error)
		SelectWithTotal() (dataList []*AssetNetdiskFile, total int64, danErr error)
		CacheSelect(redis *redisd.Redisd) (dataList []*AssetNetdiskFile, danErr error)
		Count() (total int64, danErr error)
		Inc(field string, num int) (effectRow int64, danErr error)
		Dec(field string, num int) (effectRow int64, danErr error)
		StartTrans() (tx *sql.Tx, danErr error)
		Commit(tx *sql.Tx) (danErr error)
		Rollback(tx *sql.Tx) (danErr error)
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
		Id            string `db:"id" json:"id"`
		AssetId       string `db:"asset_id" json:"assetId"`              // 关联资源id
		Name          string `db:"name" json:"name"`                     // 自定义文件别名
		StateEm       int64  `db:"state_em" json:"stateEm"`              // 文件状态
		ClasId        string `db:"clas_id" json:"clasId"`                // 所属目录id
		AuthGroupId   string `db:"auth_group_id" json:"authGroupId"`     // 权限组别id
		AuthGroupName string `db:"auth_group_name" json:"authGroupName"` // 权限组别名称
		ClasName      string `db:"clas_name" json:"clasName"`            // 所属目录名称
		Sha1          string `db:"sha1" json:"sha1"`                     // 文件哈希
		OriginalName  string `db:"original_name" json:"originalName"`    // 上传时的原始名称
		ModeEm        int64  `db:"mode_em" json:"modeEm"`                // 存储模式枚举
		Mime          string `db:"mime" json:"mime"`                     // 文件类型
		SizeNum       int64  `db:"size_num" json:"sizeNum"`              // 文件字节
		SizeText      string `db:"size_text" json:"sizeText"`            // 文件大小
		Ext           string `db:"ext" json:"ext"`                       // 文件后缀
		Path          string `db:"path" json:"path"`                     // 文件路径
		Url           string `db:"url" json:"url"`                       // 文件链接
		UserId        string `db:"user_id" json:"userId"`                // 上传用户标识
		FinishAt      int64  `db:"finish_at" json:"finishAt"`            // 完成上传时间
		PlatId        string `db:"plat_id" json:"platId"`                // 应用id
		CreateAt      int64  `db:"create_at" json:"createAt"`            // 创建时间戳
		UpdateAt      int64  `db:"update_at" json:"updateAt"`            // 更新时间戳
		DeleteAt      int64  `db:"delete_at" json:"deleteAt"`            // 删除时间戳
	}
)

// NewAssetNetdiskFileModel returns a model for the database table.
func NewAssetNetdiskFileModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) AssetNetdiskFileModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customAssetNetdiskFileModel{
		defaultAssetNetdiskFileModel: newAssetNetdiskFileModel(ctxOrNil, conn, platid),
		softDeletable:                softDeletableAssetNetdiskFile,
	}
}
func newAssetNetdiskFileModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultAssetNetdiskFileModel {
	dao := dao.NewSqlxDao(conn, "`asset_netdisk_file`", defaultAssetNetdiskFileFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultAssetNetdiskFileModel{
		ctx:             ctx,
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
func (m *defaultAssetNetdiskFileModel) LeftJoin(joinTable string) *defaultAssetNetdiskFileModel {
	m.dao.LeftJoin(joinTable)
	return m
}
func (m *defaultAssetNetdiskFileModel) RightJoin(joinTable string) *defaultAssetNetdiskFileModel {
	m.dao.RightJoin(joinTable)
	return m
}
func (m *defaultAssetNetdiskFileModel) InnerJoin(joinTable string) *defaultAssetNetdiskFileModel {
	m.dao.InnerJoin(joinTable)
	return m
}
func (m *defaultAssetNetdiskFileModel) Field(field string) *defaultAssetNetdiskFileModel {
	m.dao.Field(field)
	return m
}
func (m *defaultAssetNetdiskFileModel) Except(fields ...string) *defaultAssetNetdiskFileModel {
	m.dao.Except(fields...)
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
func (m *defaultAssetNetdiskFileModel) Total() (total int64, danErr error) {
	return m.dao.Total()
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
func (m *defaultAssetNetdiskFileModel) Insert(data *AssetNetdiskFile) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultAssetNetdiskFileModel) TxInsert(tx *sql.Tx, data *AssetNetdiskFile) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}
func (m *defaultAssetNetdiskFileModel) Delete(id ...string) (effectRow int64, danErr error) {
	return m.dao.Delete(id...)
}
func (m *defaultAssetNetdiskFileModel) TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error) {
	return m.dao.TxDelete(tx, id...)
}
func (m *defaultAssetNetdiskFileModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultAssetNetdiskFileModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultAssetNetdiskFileModel) Save(data *AssetNetdiskFile) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultAssetNetdiskFileModel) TxSave(tx *sql.Tx, data *AssetNetdiskFile) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultAssetNetdiskFileModel) StartTrans() (tx *sql.Tx, danErr error) {
	return dao.StartTrans(m.ctx, m.conn)
}
func (m *defaultAssetNetdiskFileModel) Commit(tx *sql.Tx) (danErr error) {
	return dao.Commit(tx)
}
func (m *defaultAssetNetdiskFileModel) Rollback(tx *sql.Tx) (danErr error) {
	return dao.Rollback(tx)
}
func (m *defaultAssetNetdiskFileModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultAssetNetdiskFileModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
