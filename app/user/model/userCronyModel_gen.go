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
	userCronyFieldNames          = builder.RawFieldNames(&UserCrony{})
	userCronyRows                = strings.Join(userCronyFieldNames, ",")
	defaultUserCronyFields       = strings.Join(userCronyFieldNames, ",")
	userCronyRowsExpectAutoSet   = strings.Join(stringx.Remove(userCronyFieldNames, "`delete_at`"), ",")
	userCronyRowsWithPlaceHolder = strings.Join(stringx.Remove(userCronyFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

type (
	userCronyModel interface {
		Insert(data *UserCrony) (int64, error)
		TxInsert(tx *sql.Tx, data *UserCrony) (int64, error)
		Update(data map[string]any) (int64, error)
		TxUpdate(tx *sql.Tx, data map[string]any) (int64, error)
		Save(data *UserCrony) (int64, error)
		TxSave(tx *sql.Tx, data *UserCrony) (int64, error)
		Delete(ctx context.Context, id int64) error
		Field(field string) *defaultUserCronyModel
		Alias(alias string) *defaultUserCronyModel
		Where(whereStr string, whereData ...any) *defaultUserCronyModel
		WhereId(id string) *defaultUserCronyModel
		Order(order string) *defaultUserCronyModel
		Limit(num int64) *defaultUserCronyModel
		Plat(id string) *defaultUserCronyModel
		Find() (*UserCrony, error)
		FindById(id string) (*UserCrony, error)
		CacheFind(redis *redisd.Redisd) (*UserCrony, error)
		CacheFindById(redis *redisd.Redisd, id string) (*UserCrony, error)
		Page(page int64, rows int64) *defaultUserCronyModel
		Select() ([]*UserCrony, error)
		SelectWithTotal() ([]*UserCrony, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*UserCrony, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultUserCronyModel
		Reinit() *defaultUserCronyModel
		Dao() *dao.SqlxDao
	}

	defaultUserCronyModel struct {
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

	UserCrony struct {
		Id               string `db:"id"`
		OwnerUserId      string `db:"owner_user_id"`      // 归属人id
		TargetUserId     string `db:"target_user_id"`     // 好友id
		TargetUserName   string `db:"target_user_name"`   // 好友名称
		TargetUserAvatar string `db:"target_user_avatar"` // 好友头像
		NameNote         string `db:"name_note"`          // 好友别名
		Remark           string `db:"remark"`             // 好友备注
		TypeEm           int64  `db:"type_em"`            // 好友类型
		GroupId          string `db:"group_id"`           // 组别id
		GroupName        string `db:"group_name"`         // 组别名称
		TagIds           string `db:"tag_ids"`            // 标签集合id
		PlatId           string `db:"plat_id"`
		CreateAt         int64  `db:"create_at"`
		EditAt           int64  `db:"edit_at"`
		DeleteAt         int64  `db:"delete_at"`
	}
)

func newUserCronyModel(conn sqlx.SqlConn, platId string) *defaultUserCronyModel {
	dao := dao.NewSqlxDao(conn, "`user_crony`", defaultUserCronyFields, true, "delete_at")
	dao.Plat(platId)
	return &defaultUserCronyModel{
		conn:            conn,
		dao:             dao,
		table:           "`user_crony`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultUserCronyModel) Ctx(ctx context.Context) *defaultUserCronyModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultUserCronyModel) WhereId(id string) *defaultUserCronyModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultUserCronyModel) Where(whereStr string, whereData ...any) *defaultUserCronyModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultUserCronyModel) Alias(alias string) *defaultUserCronyModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultUserCronyModel) Field(field string) *defaultUserCronyModel {
	m.dao.Field(field)
	return m
}
func (m *defaultUserCronyModel) Order(order string) *defaultUserCronyModel {
	m.dao.Order(order)
	return m
}
func (m *defaultUserCronyModel) Limit(num int64) *defaultUserCronyModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultUserCronyModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultUserCronyModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultUserCronyModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultUserCronyModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserCronyModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserCronyModel) Plat(id string) *defaultUserCronyModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultUserCronyModel) Reinit() *defaultUserCronyModel {
	m.dao.Reinit()
	return m
}
func (m *defaultUserCronyModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultUserCronyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserCronyModel) Find() (*UserCrony, error) {
	resp := &UserCrony{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserCronyModel) FindById(id string) (*UserCrony, error) {
	resp := &UserCrony{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserCronyModel) CacheFind(redis *redisd.Redisd) (*UserCrony, error) {
	resp := &UserCrony{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserCronyModel) CacheFindById(redis *redisd.Redisd, id string) (*UserCrony, error) {
	resp := &UserCrony{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserCronyModel) Select() ([]*UserCrony, error) {
	resp := make([]*UserCrony, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserCronyModel) SelectWithTotal() ([]*UserCrony, int64, error) {
	resp := make([]*UserCrony, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultUserCronyModel) CacheSelect(redis *redisd.Redisd) ([]*UserCrony, error) {
	resp := make([]*UserCrony, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserCronyModel) Page(page int64, size int64) *defaultUserCronyModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultUserCronyModel) Insert(data *UserCrony) (int64, error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultUserCronyModel) TxInsert(tx *sql.Tx, data *UserCrony) (int64, error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultUserCronyModel) Update(data map[string]any) (int64, error) {
	return m.dao.Update(data)
}
func (m *defaultUserCronyModel) TxUpdate(tx *sql.Tx, data map[string]any) (int64, error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultUserCronyModel) Save(data *UserCrony) (int64, error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultUserCronyModel) TxSave(tx *sql.Tx, data *UserCrony) (int64, error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultUserCronyModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultUserCronyModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
