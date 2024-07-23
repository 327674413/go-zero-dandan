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
	socialGroupFieldNames          = builder.RawFieldNames(&SocialGroup{})
	socialGroupRows                = strings.Join(socialGroupFieldNames, ",")
	defaultSocialGroupFields       = strings.Join(socialGroupFieldNames, ",")
	socialGroupRowsExpectAutoSet   = strings.Join(stringx.Remove(socialGroupFieldNames, "`delete_at`"), ",")
	socialGroupRowsWithPlaceHolder = strings.Join(stringx.Remove(socialGroupFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	SocialGroup_Id          dao.TableField = "id"
	SocialGroup_Code        dao.TableField = "code"
	SocialGroup_Name        dao.TableField = "name"
	SocialGroup_StateEm     dao.TableField = "state_em"
	SocialGroup_TypeEm      dao.TableField = "type_em"
	SocialGroup_CreateUid   dao.TableField = "create_uid"
	SocialGroup_IsVerify    dao.TableField = "is_verify"
	SocialGroup_NotiContent dao.TableField = "noti_content"
	SocialGroup_NotiUid     dao.TableField = "noti_uid"
	SocialGroup_Remark      dao.TableField = "remark"
	SocialGroup_PlatId      dao.TableField = "plat_id"
	SocialGroup_CreateAt    dao.TableField = "create_at"
	SocialGroup_UpdateAt    dao.TableField = "update_at"
	SocialGroup_DeleteAt    dao.TableField = "delete_at"
)

type (
	socialGroupModel interface {
		Delete(id ...string) (effectRow int64, danErr error)
		TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error)
		Insert(data *SocialGroup) (effectRow int64, danErr error)
		TxInsert(tx *sql.Tx, data *SocialGroup) (effectRow int64, danErr error)
		Update(data map[dao.TableField]any) (effectRow int64, danErr error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, danErr error)
		Save(data *SocialGroup) (effectRow int64, danErr error)
		TxSave(tx *sql.Tx, data *SocialGroup) (effectRow int64, danErr error)
		Field(field string) *defaultSocialGroupModel
		Except(fields ...string) *defaultSocialGroupModel
		Alias(alias string) *defaultSocialGroupModel
		Where(whereStr string, whereData ...any) *defaultSocialGroupModel
		WhereId(id string) *defaultSocialGroupModel
		Order(order string) *defaultSocialGroupModel
		Limit(num int64) *defaultSocialGroupModel
		Plat(id string) *defaultSocialGroupModel
		Find() (*SocialGroup, error)
		FindById(id string) (data *SocialGroup, danErr error)
		CacheFind(redis *redisd.Redisd) (data *SocialGroup, danErr error)
		CacheFindById(redis *redisd.Redisd, id string) (data *SocialGroup, danErr error)
		Page(page int64, rows int64) *defaultSocialGroupModel
		Total() (total int64, danErr error)
		Select() (dataList []*SocialGroup, danErr error)
		SelectWithTotal() (dataList []*SocialGroup, total int64, danErr error)
		CacheSelect(redis *redisd.Redisd) (dataList []*SocialGroup, danErr error)
		Count() (total int64, danErr error)
		Inc(field string, num int) (effectRow int64, danErr error)
		Dec(field string, num int) (effectRow int64, danErr error)
		StartTrans() (tx *sql.Tx, danErr error)
		Commit(tx *sql.Tx) (danErr error)
		Rollback(tx *sql.Tx) (danErr error)
		Ctx(ctx context.Context) *defaultSocialGroupModel
		Reinit() *defaultSocialGroupModel
		Dao() *dao.SqlxDao
	}

	defaultSocialGroupModel struct {
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

	SocialGroup struct {
		Id          string `db:"id" json:"id"`
		Code        string `db:"code" json:"code"`                // 群号
		Name        string `db:"name" json:"name"`                // 群组名
		StateEm     int64  `db:"state_em" json:"stateEm"`         // 群组状态
		TypeEm      int64  `db:"type_em" json:"typeEm"`           // 群组类型
		CreateUid   string `db:"create_uid" json:"createUid"`     // 创建人id
		IsVerify    int64  `db:"is_verify" json:"isVerify"`       // 是否需要验证
		NotiContent string `db:"noti_content" json:"notiContent"` // 群公告内容
		NotiUid     string `db:"noti_uid" json:"notiUid"`         // 群公告编写人
		Remark      string `db:"remark" json:"remark"`            // 备注
		PlatId      string `db:"plat_id" json:"platId"`           // 应用id
		CreateAt    int64  `db:"create_at" json:"createAt"`       // 创建时间戳
		UpdateAt    int64  `db:"update_at" json:"updateAt"`       // 更新时间戳
		DeleteAt    int64  `db:"delete_at" json:"deleteAt"`       // 删除时间戳
	}
)

// NewSocialGroupModel returns a model for the database table.
func NewSocialGroupModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) SocialGroupModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customSocialGroupModel{
		defaultSocialGroupModel: newSocialGroupModel(ctxOrNil, conn, platid),
		softDeletable:           softDeletableSocialGroup,
	}
}
func newSocialGroupModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultSocialGroupModel {
	dao := dao.NewSqlxDao(conn, "`social_group`", defaultSocialGroupFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultSocialGroupModel{
		ctx:             ctx,
		conn:            conn,
		dao:             dao,
		table:           "`social_group`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultSocialGroupModel) Ctx(ctx context.Context) *defaultSocialGroupModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultSocialGroupModel) WhereId(id string) *defaultSocialGroupModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultSocialGroupModel) Where(whereStr string, whereData ...any) *defaultSocialGroupModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultSocialGroupModel) Alias(alias string) *defaultSocialGroupModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultSocialGroupModel) Field(field string) *defaultSocialGroupModel {
	m.dao.Field(field)
	return m
}
func (m *defaultSocialGroupModel) Except(fields ...string) *defaultSocialGroupModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultSocialGroupModel) Order(order string) *defaultSocialGroupModel {
	m.dao.Order(order)
	return m
}
func (m *defaultSocialGroupModel) Limit(num int64) *defaultSocialGroupModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultSocialGroupModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultSocialGroupModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultSocialGroupModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultSocialGroupModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialGroupModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialGroupModel) Plat(id string) *defaultSocialGroupModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultSocialGroupModel) Reinit() *defaultSocialGroupModel {
	m.dao.Reinit()
	return m
}
func (m *defaultSocialGroupModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultSocialGroupModel) Find() (*SocialGroup, error) {
	resp := &SocialGroup{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupModel) FindById(id string) (*SocialGroup, error) {
	resp := &SocialGroup{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupModel) CacheFind(redis *redisd.Redisd) (*SocialGroup, error) {
	resp := &SocialGroup{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupModel) CacheFindById(redis *redisd.Redisd, id string) (*SocialGroup, error) {
	resp := &SocialGroup{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupModel) Total() (total int64, danErr error) {
	return m.dao.Total()
}
func (m *defaultSocialGroupModel) Select() ([]*SocialGroup, error) {
	resp := make([]*SocialGroup, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupModel) SelectWithTotal() ([]*SocialGroup, int64, error) {
	resp := make([]*SocialGroup, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultSocialGroupModel) CacheSelect(redis *redisd.Redisd) ([]*SocialGroup, error) {
	resp := make([]*SocialGroup, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultSocialGroupModel) Page(page int64, size int64) *defaultSocialGroupModel {
	m.dao.Page(page, size)
	return m
}
func (m *defaultSocialGroupModel) Insert(data *SocialGroup) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultSocialGroupModel) TxInsert(tx *sql.Tx, data *SocialGroup) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}
func (m *defaultSocialGroupModel) Delete(id ...string) (effectRow int64, danErr error) {
	return m.dao.Delete(id...)
}
func (m *defaultSocialGroupModel) TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error) {
	return m.dao.TxDelete(tx, id...)
}
func (m *defaultSocialGroupModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultSocialGroupModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultSocialGroupModel) Save(data *SocialGroup) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialGroupModel) TxSave(tx *sql.Tx, data *SocialGroup) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialGroupModel) StartTrans() (tx *sql.Tx, danErr error) {
	return dao.StartTrans(m.ctx, m.conn)
}
func (m *defaultSocialGroupModel) Commit(tx *sql.Tx) (danErr error) {
	return dao.Commit(tx)
}
func (m *defaultSocialGroupModel) Rollback(tx *sql.Tx) (danErr error) {
	return dao.Rollback(tx)
}
func (m *defaultSocialGroupModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultSocialGroupModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
