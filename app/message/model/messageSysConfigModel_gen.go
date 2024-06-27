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
	messageSysConfigFieldNames          = builder.RawFieldNames(&MessageSysConfig{})
	messageSysConfigRows                = strings.Join(messageSysConfigFieldNames, ",")
	defaultMessageSysConfigFields       = strings.Join(messageSysConfigFieldNames, ",")
	messageSysConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(messageSysConfigFieldNames, "`delete_at`"), ",")
	messageSysConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(messageSysConfigFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	MessageSysConfig_Id              dao.TableField = "id"
	MessageSysConfig_SmsLimitHourNum dao.TableField = "sms_limit_hour_num"
	MessageSysConfig_SmsLimitDayNum  dao.TableField = "sms_limit_day_num"
	MessageSysConfig_CreateAt        dao.TableField = "create_at"
	MessageSysConfig_UpdateAt        dao.TableField = "update_at"
	MessageSysConfig_DeleteAt        dao.TableField = "delete_at"
)

type (
	messageSysConfigModel interface {
		Insert(data *MessageSysConfig) (effectRow int64, err error)
		TxInsert(tx *sql.Tx, data *MessageSysConfig) (effectRow int64, err error)
		Update(data map[dao.TableField]any) (effectRow int64, err error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error)
		Save(data *MessageSysConfig) (effectRow int64, err error)
		TxSave(tx *sql.Tx, data *MessageSysConfig) (effectRow int64, err error)
		Delete(ctx context.Context, id string) error
		Field(field string) *defaultMessageSysConfigModel
		Except(fields ...string) *defaultMessageSysConfigModel
		Alias(alias string) *defaultMessageSysConfigModel
		Where(whereStr string, whereData ...any) *defaultMessageSysConfigModel
		WhereId(id string) *defaultMessageSysConfigModel
		Order(order string) *defaultMessageSysConfigModel
		Limit(num int64) *defaultMessageSysConfigModel
		Plat(id string) *defaultMessageSysConfigModel
		Find() (*MessageSysConfig, error)
		FindById(id string) (*MessageSysConfig, error)
		CacheFind(redis *redisd.Redisd) (*MessageSysConfig, error)
		CacheFindById(redis *redisd.Redisd, id string) (*MessageSysConfig, error)
		Page(page int64, rows int64) *defaultMessageSysConfigModel
		Total() (total int64, err error)
		Select() ([]*MessageSysConfig, error)
		SelectWithTotal() ([]*MessageSysConfig, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*MessageSysConfig, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultMessageSysConfigModel
		Reinit() *defaultMessageSysConfigModel
		Dao() *dao.SqlxDao
	}

	defaultMessageSysConfigModel struct {
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

	MessageSysConfig struct {
		Id              string `db:"id"`
		SmsLimitHourNum int64  `db:"sms_limit_hour_num"` // 最近一小时内获取上限,0则不管控
		SmsLimitDayNum  int64  `db:"sms_limit_day_num"`  // 每日获取上限,0则不管控
		CreateAt        int64  `db:"create_at"`          // 创建时间戳
		UpdateAt        int64  `db:"update_at"`          // 更新时间戳
		DeleteAt        int64  `db:"delete_at"`          // 删除时间戳
	}
)

func newMessageSysConfigModel(conn sqlx.SqlConn, platId string) *defaultMessageSysConfigModel {
	dao := dao.NewSqlxDao(conn, "`message_sys_config`", defaultMessageSysConfigFields, true, "delete_at")
	dao.Plat(platId)
	return &defaultMessageSysConfigModel{
		conn:            conn,
		dao:             dao,
		table:           "`message_sys_config`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultMessageSysConfigModel) Ctx(ctx context.Context) *defaultMessageSysConfigModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultMessageSysConfigModel) WhereId(id string) *defaultMessageSysConfigModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultMessageSysConfigModel) Where(whereStr string, whereData ...any) *defaultMessageSysConfigModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultMessageSysConfigModel) Alias(alias string) *defaultMessageSysConfigModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultMessageSysConfigModel) Field(field string) *defaultMessageSysConfigModel {
	m.dao.Field(field)
	return m
}
func (m *defaultMessageSysConfigModel) Except(fields ...string) *defaultMessageSysConfigModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultMessageSysConfigModel) Order(order string) *defaultMessageSysConfigModel {
	m.dao.Order(order)
	return m
}
func (m *defaultMessageSysConfigModel) Limit(num int64) *defaultMessageSysConfigModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultMessageSysConfigModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultMessageSysConfigModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultMessageSysConfigModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultMessageSysConfigModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultMessageSysConfigModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultMessageSysConfigModel) Plat(id string) *defaultMessageSysConfigModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultMessageSysConfigModel) Reinit() *defaultMessageSysConfigModel {
	m.dao.Reinit()
	return m
}
func (m *defaultMessageSysConfigModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultMessageSysConfigModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultMessageSysConfigModel) Find() (*MessageSysConfig, error) {
	resp := &MessageSysConfig{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSysConfigModel) FindById(id string) (*MessageSysConfig, error) {
	resp := &MessageSysConfig{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSysConfigModel) CacheFind(redis *redisd.Redisd) (*MessageSysConfig, error) {
	resp := &MessageSysConfig{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSysConfigModel) CacheFindById(redis *redisd.Redisd, id string) (*MessageSysConfig, error) {
	resp := &MessageSysConfig{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSysConfigModel) Total() (total int64, err error) {
	return m.dao.Total()
}
func (m *defaultMessageSysConfigModel) Select() ([]*MessageSysConfig, error) {
	resp := make([]*MessageSysConfig, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSysConfigModel) SelectWithTotal() ([]*MessageSysConfig, int64, error) {
	resp := make([]*MessageSysConfig, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultMessageSysConfigModel) CacheSelect(redis *redisd.Redisd) ([]*MessageSysConfig, error) {
	resp := make([]*MessageSysConfig, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultMessageSysConfigModel) Page(page int64, size int64) *defaultMessageSysConfigModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultMessageSysConfigModel) Insert(data *MessageSysConfig) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultMessageSysConfigModel) TxInsert(tx *sql.Tx, data *MessageSysConfig) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultMessageSysConfigModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultMessageSysConfigModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultMessageSysConfigModel) Save(data *MessageSysConfig) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultMessageSysConfigModel) TxSave(tx *sql.Tx, data *MessageSysConfig) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultMessageSysConfigModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultMessageSysConfigModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
