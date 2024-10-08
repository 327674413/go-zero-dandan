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
	messageSmsSendFieldNames          = builder.RawFieldNames(&MessageSmsSend{})
	messageSmsSendRows                = strings.Join(messageSmsSendFieldNames, ",")
	defaultMessageSmsSendFields       = strings.Join(messageSmsSendFieldNames, ",")
	messageSmsSendRowsExpectAutoSet   = strings.Join(stringx.Remove(messageSmsSendFieldNames, "`delete_at`"), ",")
	messageSmsSendRowsWithPlaceHolder = strings.Join(stringx.Remove(messageSmsSendFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	MessageSmsSend_Id       dao.TableField = "id"
	MessageSmsSend_Phone    dao.TableField = "phone"
	MessageSmsSend_Content  dao.TableField = "content"
	MessageSmsSend_StateEm  dao.TableField = "state_em"
	MessageSmsSend_Err      dao.TableField = "err"
	MessageSmsSend_TempId   dao.TableField = "temp_id"
	MessageSmsSend_PlatId   dao.TableField = "plat_id"
	MessageSmsSend_CreateAt dao.TableField = "create_at"
	MessageSmsSend_UpdateAt dao.TableField = "update_at"
	MessageSmsSend_DeleteAt dao.TableField = "delete_at"
)

type (
	messageSmsSendModel interface {
		Delete(id ...string) (effectRow int64, danErr error)
		TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error)
		Insert(data *MessageSmsSend) (effectRow int64, danErr error)
		TxInsert(tx *sql.Tx, data *MessageSmsSend) (effectRow int64, danErr error)
		Update(data map[dao.TableField]any) (effectRow int64, danErr error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, danErr error)
		Save(data *MessageSmsSend) (effectRow int64, danErr error)
		TxSave(tx *sql.Tx, data *MessageSmsSend) (effectRow int64, danErr error)
		Field(field string) *defaultMessageSmsSendModel
		Except(fields ...string) *defaultMessageSmsSendModel
		Alias(alias string) *defaultMessageSmsSendModel
		LeftJoin(joinTable string) *defaultMessageSmsSendModel
		RightJoin(joinTable string) *defaultMessageSmsSendModel
		InnerJoin(joinTable string) *defaultMessageSmsSendModel
		Where(whereStr string, whereData ...any) *defaultMessageSmsSendModel
		WhereId(id string) *defaultMessageSmsSendModel
		Order(order string) *defaultMessageSmsSendModel
		Limit(num int64) *defaultMessageSmsSendModel
		Plat(id string) *defaultMessageSmsSendModel
		Find() (*MessageSmsSend, error)
		FindById(id string) (data *MessageSmsSend, danErr error)
		CacheFind(redis *redisd.Redisd) (data *MessageSmsSend, danErr error)
		CacheFindById(redis *redisd.Redisd, id string) (data *MessageSmsSend, danErr error)
		Page(page int64, rows int64) *defaultMessageSmsSendModel
		Total() (total int64, danErr error)
		Select() (dataList []*MessageSmsSend, danErr error)
		SelectWithTotal() (dataList []*MessageSmsSend, total int64, danErr error)
		CacheSelect(redis *redisd.Redisd) (dataList []*MessageSmsSend, danErr error)
		Count() (total int64, danErr error)
		Inc(field string, num int) (effectRow int64, danErr error)
		Dec(field string, num int) (effectRow int64, danErr error)
		StartTrans() (tx *sql.Tx, danErr error)
		Commit(tx *sql.Tx) (danErr error)
		Rollback(tx *sql.Tx) (danErr error)
		Ctx(ctx context.Context) *defaultMessageSmsSendModel
		Reinit() *defaultMessageSmsSendModel
		Dao() *dao.SqlxDao
	}

	defaultMessageSmsSendModel struct {
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

	MessageSmsSend struct {
		Id       string `db:"id" json:"id"`
		Phone    string `db:"phone" json:"phone"`        // 发送手机号
		Content  string `db:"content" json:"content"`    // 发送手机号
		StateEm  int64  `db:"state_em" json:"stateEm"`   // 发送状态枚举
		Err      string `db:"err" json:"err"`            // 发送错误时的错误信息
		TempId   string `db:"temp_id" json:"tempId"`     // 模版id
		PlatId   string `db:"plat_id" json:"platId"`     // 应用id
		CreateAt int64  `db:"create_at" json:"createAt"` // 创建时间戳
		UpdateAt int64  `db:"update_at" json:"updateAt"` // 更新时间戳
		DeleteAt int64  `db:"delete_at" json:"deleteAt"` // 删除时间戳
	}
)

// NewMessageSmsSendModel returns a model for the database table.
func NewMessageSmsSendModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) MessageSmsSendModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customMessageSmsSendModel{
		defaultMessageSmsSendModel: newMessageSmsSendModel(ctxOrNil, conn, platid),
		softDeletable:              softDeletableMessageSmsSend,
	}
}
func newMessageSmsSendModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultMessageSmsSendModel {
	dao := dao.NewSqlxDao(conn, "`message_sms_send`", defaultMessageSmsSendFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultMessageSmsSendModel{
		ctx:             ctx,
		conn:            conn,
		dao:             dao,
		table:           "`message_sms_send`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultMessageSmsSendModel) Ctx(ctx context.Context) *defaultMessageSmsSendModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultMessageSmsSendModel) WhereId(id string) *defaultMessageSmsSendModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultMessageSmsSendModel) Where(whereStr string, whereData ...any) *defaultMessageSmsSendModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultMessageSmsSendModel) Alias(alias string) *defaultMessageSmsSendModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultMessageSmsSendModel) LeftJoin(joinTable string) *defaultMessageSmsSendModel {
	m.dao.LeftJoin(joinTable)
	return m
}
func (m *defaultMessageSmsSendModel) RightJoin(joinTable string) *defaultMessageSmsSendModel {
	m.dao.RightJoin(joinTable)
	return m
}
func (m *defaultMessageSmsSendModel) InnerJoin(joinTable string) *defaultMessageSmsSendModel {
	m.dao.InnerJoin(joinTable)
	return m
}
func (m *defaultMessageSmsSendModel) Field(field string) *defaultMessageSmsSendModel {
	m.dao.Field(field)
	return m
}
func (m *defaultMessageSmsSendModel) Except(fields ...string) *defaultMessageSmsSendModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultMessageSmsSendModel) Order(order string) *defaultMessageSmsSendModel {
	m.dao.Order(order)
	return m
}
func (m *defaultMessageSmsSendModel) Limit(num int64) *defaultMessageSmsSendModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultMessageSmsSendModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultMessageSmsSendModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultMessageSmsSendModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultMessageSmsSendModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultMessageSmsSendModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultMessageSmsSendModel) Plat(id string) *defaultMessageSmsSendModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultMessageSmsSendModel) Reinit() *defaultMessageSmsSendModel {
	m.dao.Reinit()
	return m
}
func (m *defaultMessageSmsSendModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultMessageSmsSendModel) Find() (*MessageSmsSend, error) {
	resp := &MessageSmsSend{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsSendModel) FindById(id string) (*MessageSmsSend, error) {
	resp := &MessageSmsSend{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsSendModel) CacheFind(redis *redisd.Redisd) (*MessageSmsSend, error) {
	resp := &MessageSmsSend{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsSendModel) CacheFindById(redis *redisd.Redisd, id string) (*MessageSmsSend, error) {
	resp := &MessageSmsSend{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsSendModel) Total() (total int64, danErr error) {
	return m.dao.Total()
}
func (m *defaultMessageSmsSendModel) Select() ([]*MessageSmsSend, error) {
	resp := make([]*MessageSmsSend, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultMessageSmsSendModel) SelectWithTotal() ([]*MessageSmsSend, int64, error) {
	resp := make([]*MessageSmsSend, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultMessageSmsSendModel) CacheSelect(redis *redisd.Redisd) ([]*MessageSmsSend, error) {
	resp := make([]*MessageSmsSend, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultMessageSmsSendModel) Page(page int64, size int64) *defaultMessageSmsSendModel {
	m.dao.Page(page, size)
	return m
}
func (m *defaultMessageSmsSendModel) Insert(data *MessageSmsSend) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultMessageSmsSendModel) TxInsert(tx *sql.Tx, data *MessageSmsSend) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}
func (m *defaultMessageSmsSendModel) Delete(id ...string) (effectRow int64, danErr error) {
	return m.dao.Delete(id...)
}
func (m *defaultMessageSmsSendModel) TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error) {
	return m.dao.TxDelete(tx, id...)
}
func (m *defaultMessageSmsSendModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultMessageSmsSendModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultMessageSmsSendModel) Save(data *MessageSmsSend) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultMessageSmsSendModel) TxSave(tx *sql.Tx, data *MessageSmsSend) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultMessageSmsSendModel) StartTrans() (tx *sql.Tx, danErr error) {
	return dao.StartTrans(m.ctx, m.conn)
}
func (m *defaultMessageSmsSendModel) Commit(tx *sql.Tx) (danErr error) {
	return dao.Commit(tx)
}
func (m *defaultMessageSmsSendModel) Rollback(tx *sql.Tx) (danErr error) {
	return dao.Rollback(tx)
}
func (m *defaultMessageSmsSendModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultMessageSmsSendModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
