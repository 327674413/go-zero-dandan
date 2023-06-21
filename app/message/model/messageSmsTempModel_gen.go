// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"go-zero-dandan/common/redisd"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	messageSmsTempFieldNames          = builder.RawFieldNames(&MessageSmsTemp{})
	messageSmsTempRows                = strings.Join(messageSmsTempFieldNames, ",")
	messageSmsTempRowsExpectAutoSet   = strings.Join(stringx.Remove(messageSmsTempFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messageSmsTempRowsWithPlaceHolder = strings.Join(stringx.Remove(messageSmsTempFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	messageSmsTempModel interface {
		Insert(ctx context.Context, data *MessageSmsTemp) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*MessageSmsTemp, error)
		Update(ctx context.Context, data *MessageSmsTemp) error
		Delete(ctx context.Context, id int64) error
		Field(field string) *defaultMessageSmsTempModel
		Alias(alias string) *defaultMessageSmsTempModel
		WhereStr(whereStr string) *defaultMessageSmsTempModel
		WhereId(id int) *defaultMessageSmsTempModel
		WhereMap(whereMap map[string]any) *defaultMessageSmsTempModel
		WhereRaw(whereStr string, whereData []any) *defaultMessageSmsTempModel
		Order(order string) *defaultMessageSmsTempModel
		Plat(id int) *defaultMessageSmsTempModel
		Find(ctx context.Context, id ...any) (*MessageSmsTemp, error)
		CacheFind(ctx context.Context, redis *redisd.Redisd, id ...int64) (*MessageSmsTemp, error)
		Page(ctx context.Context, page int, rows int) ([]*MessageSmsTemp, error)
		List(ctx context.Context) ([]*MessageSmsTemp, error)
		Count(ctx context.Context) int
		Inc(ctx context.Context, field string, num int) error
		Dec(ctx context.Context, field string, num int) error
	}

	defaultMessageSmsTempModel struct {
		conn            sqlx.SqlConn
		table           string
		softDeleteField string
		SoftDeletable   bool
		fieldSql        string
		whereSql        string
		aliasSql        string
		orderSql        string
		whereData       []any
		err             error
	}

	MessageSmsTemp struct {
		Id          int64  `db:"id"`
		Name        string `db:"name"`
		SecretId    string `db:"secret_id"`     // SecretId
		SecretKey   string `db:"secret_key"`    // SecretKey
		Region      string `db:"region"`        // region
		SmsSdkAppid string `db:"sms_sdk_appid"` // SmsSdkAppId
		SignName    string `db:"sign_name"`     // SignName
		TemplateId  string `db:"template_id"`   // TemplateId
		PlatId      int64  `db:"plat_id"`       // 应用id
		CreateAt    int64  `db:"create_at"`     // 创建时间戳
		UpdateAt    int64  `db:"update_at"`     // 更新时间戳
		DeleteAt    int64  `db:"delete_at"`     // 删除时间戳
	}
)

func newMessageSmsTempModel(conn sqlx.SqlConn) *defaultMessageSmsTempModel {
	return &defaultMessageSmsTempModel{
		conn:            conn,
		table:           "`message_sms_temp`",
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}

func (m *defaultMessageSmsTempModel) WhereId(id int) *defaultMessageSmsTempModel {
	m.whereSql = "id=?"
	m.whereData = append(m.whereData, id)
	return m
}
func (m *defaultMessageSmsTempModel) Page(ctx context.Context, page int, rows int) ([]*MessageSmsTemp, error) {

	return nil, nil
}
func (m *defaultMessageSmsTempModel) List(ctx context.Context) ([]*MessageSmsTemp, error) {

	return nil, nil
}
func (m *defaultMessageSmsTempModel) WhereStr(whereStr string) *defaultMessageSmsTempModel {
	return m
}

func (m *defaultMessageSmsTempModel) WhereMap(whereMap map[string]any) *defaultMessageSmsTempModel {
	return m
}
func (m *defaultMessageSmsTempModel) WhereRaw(whereStr string, whereData []any) *defaultMessageSmsTempModel {
	if m.whereSql != "" {
		m.whereSql += " AND (" + whereStr + ")"
	} else {
		m.whereSql = "(" + whereStr + ")"
	}
	m.whereData = append(m.whereData, whereData...)
	return m
}
func (m *defaultMessageSmsTempModel) Alias(field string) *defaultMessageSmsTempModel {
	m.aliasSql = field
	return m
}
func (m *defaultMessageSmsTempModel) Field(field string) *defaultMessageSmsTempModel {
	m.fieldSql = field
	return m
}
func (m *defaultMessageSmsTempModel) Order(order string) *defaultMessageSmsTempModel {

	return m
}
func (m *defaultMessageSmsTempModel) Count(ctx context.Context) int {

	return 0
}
func (m *defaultMessageSmsTempModel) Inc(ctx context.Context, field string, num int) error {

	return nil
}
func (m *defaultMessageSmsTempModel) Dec(ctx context.Context, field string, num int) error {

	return nil
}
func (m *defaultMessageSmsTempModel) Plat(id int) *defaultMessageSmsTempModel {

	return nil
}
func (m *defaultMessageSmsTempModel) CacheFind(ctx context.Context, redis *redisd.Redisd, id ...int64) (*MessageSmsTemp, error) {
	resp := &MessageSmsTemp{}
	cacheField := "model_" + m.tableName()
	cacheKey := strconv.FormatInt(id[0], 10)
	// todo::需要把where条件一起放进去作为key，这样就能支持更多的自动缓存查询
	err := redis.GetData(cacheField, cacheKey, resp)
	if err == nil {
		return resp, nil
	}
	resp, err = m.Find(ctx, id[0])
	fmt.Println(resp, err)
	if err != nil {
		return resp, err
	}
	if resp.Id != 0 {
		redis.SetData(cacheField, cacheKey, resp)
	}
	return resp, nil
}

func (m *defaultMessageSmsTempModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultMessageSmsTempModel) Find(ctx context.Context, id ...any) (*MessageSmsTemp, error) {
	var err error
	if err = m.err; err != nil {
		m.err = nil
		return nil, err
	}
	var resp MessageSmsTemp
	var sql string
	field := messageSmsTempRows
	if m.fieldSql != "" {
		field = m.fieldSql
	}
	if len(id) > 0 {
		sql = fmt.Sprintf("select %s from %s where id=? limit 1", field, m.table)
		err = m.conn.QueryRowPartialCtx(ctx, &resp, sql, id[0]) //QueryRowCtx 必须字段都覆盖
	} else {
		sql = fmt.Sprintf("select %s from %s %s where "+m.whereSql+" limit 1", field, m.table, m.aliasSql)
		err = m.conn.QueryRowPartialCtx(ctx, &resp, sql, m.whereData...)
	}
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return &resp, nil
	default:
		return nil, err
	}
}
func (m *defaultMessageSmsTempModel) FindOne(ctx context.Context, id int64) (*MessageSmsTemp, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", messageSmsTempRows, m.table)
	var resp MessageSmsTemp
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessageSmsTempModel) Insert(ctx context.Context, data *MessageSmsTemp) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, messageSmsTempRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.SecretId, data.SecretKey, data.Region, data.SmsSdkAppid, data.SignName, data.TemplateId, data.PlatId, data.DeleteAt)
	return ret, err
}

func (m *defaultMessageSmsTempModel) Update(ctx context.Context, data *MessageSmsTemp) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, messageSmsTempRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.SecretId, data.SecretKey, data.Region, data.SmsSdkAppid, data.SignName, data.TemplateId, data.PlatId, data.DeleteAt, data.Id)
	return err
}

func (m *defaultMessageSmsTempModel) tableName() string {
	return m.table
}