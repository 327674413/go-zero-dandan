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
	messageSmsTempFieldNames          = builder.RawFieldNames(&MessageSmsTemp{})
	messageSmsTempRows                = strings.Join(messageSmsTempFieldNames, ",")
	defaultMessageSmsTempFields       = strings.Join(messageSmsTempFieldNames, ",")
	messageSmsTempRowsExpectAutoSet   = strings.Join(stringx.Remove(messageSmsTempFieldNames, "`delete_at`"), ",")
	messageSmsTempRowsWithPlaceHolder = strings.Join(stringx.Remove(messageSmsTempFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	MessageSmsTemp_Id          dao.TableField = "id"
	MessageSmsTemp_Name        dao.TableField = "name"
	MessageSmsTemp_SecretId    dao.TableField = "secret_id"
	MessageSmsTemp_SecretKey   dao.TableField = "secret_key"
	MessageSmsTemp_Region      dao.TableField = "region"
	MessageSmsTemp_SmsSdkAppid dao.TableField = "sms_sdk_appid"
	MessageSmsTemp_SignName    dao.TableField = "sign_name"
	MessageSmsTemp_TemplateId  dao.TableField = "template_id"
	MessageSmsTemp_PlatId      dao.TableField = "plat_id"
	MessageSmsTemp_CreateAt    dao.TableField = "create_at"
	MessageSmsTemp_UpdateAt    dao.TableField = "update_at"
	MessageSmsTemp_DeleteAt    dao.TableField = "delete_at"
)

type (
	messageSmsTempModel interface {
		Insert(data *MessageSmsTemp) (effectRow int64, err error)
		TxInsert(tx *sql.Tx, data *MessageSmsTemp) (effectRow int64, err error)
		Update(data map[dao.TableField]any) (effectRow int64, err error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error)
		Save(data *MessageSmsTemp) (effectRow int64, err error)
		TxSave(tx *sql.Tx, data *MessageSmsTemp) (effectRow int64, err error)
		Delete(ctx context.Context, id string) error
		Field(field string) *defaultMessageSmsTempModel
		Except(fields ...string) *defaultMessageSmsTempModel
		Alias(alias string) *defaultMessageSmsTempModel
		Where(whereStr string, whereData ...any) *defaultMessageSmsTempModel
		WhereId(id string) *defaultMessageSmsTempModel
		Order(order string) *defaultMessageSmsTempModel
		Limit(num int64) *defaultMessageSmsTempModel
		Plat(id string) *defaultMessageSmsTempModel
		Find() (*MessageSmsTemp, error)
		FindById(id string) (*MessageSmsTemp, error)
		CacheFind(redis *redisd.Redisd) (*MessageSmsTemp, error)
		CacheFindById(redis *redisd.Redisd, id string) (*MessageSmsTemp, error)
		Page(page int64, rows int64) *defaultMessageSmsTempModel
		Total() (total int64, err error)
		Select() ([]*MessageSmsTemp, error)
		SelectWithTotal() ([]*MessageSmsTemp, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*MessageSmsTemp, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultMessageSmsTempModel
		Reinit() *defaultMessageSmsTempModel
		Dao() *dao.SqlxDao
	}

	defaultMessageSmsTempModel struct {
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

	MessageSmsTemp struct {
		Id          string `db:"id" json:"id"`
		Name        string `db:"name" json:"name"`
		SecretId    string `db:"secret_id" json:"secretId"`        // SecretId
		SecretKey   string `db:"secret_key" json:"secretKey"`      // SecretKey
		Region      string `db:"region" json:"region"`             // region
		SmsSdkAppid string `db:"sms_sdk_appid" json:"smsSdkAppid"` // SmsSdkAppId
		SignName    string `db:"sign_name" json:"signName"`        // SignName
		TemplateId  string `db:"template_id" json:"templateId"`    // TemplateId
		PlatId      string `db:"plat_id" json:"platId"`            // 应用id
		CreateAt    int64  `db:"create_at" json:"createAt"`        // 创建时间戳
		UpdateAt    int64  `db:"update_at" json:"updateAt"`        // 更新时间戳
		DeleteAt    int64  `db:"delete_at" json:"deleteAt"`        // 删除时间戳
	}
)

// NewMessageSmsTempModel returns a model for the database table.
func NewMessageSmsTempModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) MessageSmsTempModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customMessageSmsTempModel{
		defaultMessageSmsTempModel: newMessageSmsTempModel(ctxOrNil, conn, platid),
		softDeletable:              softDeletableMessageSmsTemp,
	}
}
func newMessageSmsTempModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultMessageSmsTempModel {
	dao := dao.NewSqlxDao(conn, "`message_sms_temp`", defaultMessageSmsTempFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultMessageSmsTempModel{
		ctx:             ctx,
		conn:            conn,
		dao:             dao,
		table:           "`message_sms_temp`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultMessageSmsTempModel) Ctx(ctx context.Context) *defaultMessageSmsTempModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultMessageSmsTempModel) WhereId(id string) *defaultMessageSmsTempModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultMessageSmsTempModel) Where(whereStr string, whereData ...any) *defaultMessageSmsTempModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultMessageSmsTempModel) Alias(alias string) *defaultMessageSmsTempModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultMessageSmsTempModel) Field(field string) *defaultMessageSmsTempModel {
	m.dao.Field(field)
	return m
}
func (m *defaultMessageSmsTempModel) Except(fields ...string) *defaultMessageSmsTempModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultMessageSmsTempModel) Order(order string) *defaultMessageSmsTempModel {
	m.dao.Order(order)
	return m
}
func (m *defaultMessageSmsTempModel) Limit(num int64) *defaultMessageSmsTempModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultMessageSmsTempModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultMessageSmsTempModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultMessageSmsTempModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultMessageSmsTempModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultMessageSmsTempModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultMessageSmsTempModel) Plat(id string) *defaultMessageSmsTempModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultMessageSmsTempModel) Reinit() *defaultMessageSmsTempModel {
	m.dao.Reinit()
	return m
}
func (m *defaultMessageSmsTempModel) Dao() *dao.SqlxDao {
	return m.dao
}

func (m *defaultMessageSmsTempModel) Find() (*MessageSmsTemp, error) {
	resp := &MessageSmsTemp{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsTempModel) FindById(id string) (*MessageSmsTemp, error) {
	resp := &MessageSmsTemp{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsTempModel) CacheFind(redis *redisd.Redisd) (*MessageSmsTemp, error) {
	resp := &MessageSmsTemp{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsTempModel) CacheFindById(redis *redisd.Redisd, id string) (*MessageSmsTemp, error) {
	resp := &MessageSmsTemp{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsTempModel) Total() (total int64, err error) {
	return m.dao.Total()
}
func (m *defaultMessageSmsTempModel) Select() ([]*MessageSmsTemp, error) {
	resp := make([]*MessageSmsTemp, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsTempModel) SelectWithTotal() ([]*MessageSmsTemp, int64, error) {
	resp := make([]*MessageSmsTemp, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultMessageSmsTempModel) CacheSelect(redis *redisd.Redisd) ([]*MessageSmsTemp, error) {
	resp := make([]*MessageSmsTemp, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultMessageSmsTempModel) Page(page int64, size int64) *defaultMessageSmsTempModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultMessageSmsTempModel) Insert(data *MessageSmsTemp) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultMessageSmsTempModel) TxInsert(tx *sql.Tx, data *MessageSmsTemp) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultMessageSmsTempModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultMessageSmsTempModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultMessageSmsTempModel) Save(data *MessageSmsTemp) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultMessageSmsTempModel) TxSave(tx *sql.Tx, data *MessageSmsTemp) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultMessageSmsTempModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultMessageSmsTempModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
