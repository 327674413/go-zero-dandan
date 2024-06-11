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
	socialFriendFieldNames          = builder.RawFieldNames(&SocialFriend{})
	socialFriendRows                = strings.Join(socialFriendFieldNames, ",")
	defaultSocialFriendFields       = strings.Join(socialFriendFieldNames, ",")
	socialFriendRowsExpectAutoSet   = strings.Join(stringx.Remove(socialFriendFieldNames, "`delete_at`"), ",")
	socialFriendRowsWithPlaceHolder = strings.Join(stringx.Remove(socialFriendFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

type (
	socialFriendModel interface {
		Insert(data *SocialFriend) (int64, error)
		TxInsert(tx *sql.Tx, data *SocialFriend) (int64, error)
		Update(data map[string]any) (int64, error)
		TxUpdate(tx *sql.Tx, data map[string]any) (int64, error)
		Save(data *SocialFriend) (int64, error)
		TxSave(tx *sql.Tx, data *SocialFriend) (int64, error)
		Delete(ctx context.Context, id int64) error
		Field(field string) *defaultSocialFriendModel
		Alias(alias string) *defaultSocialFriendModel
		Where(whereStr string, whereData ...any) *defaultSocialFriendModel
		WhereId(id int64) *defaultSocialFriendModel
		Order(order string) *defaultSocialFriendModel
		Limit(num int64) *defaultSocialFriendModel
		Plat(id int64) *defaultSocialFriendModel
		Find() (*SocialFriend, error)
		FindById(id int64) (*SocialFriend, error)
		CacheFind(redis *redisd.Redisd) (*SocialFriend, error)
		CacheFindById(redis *redisd.Redisd, id int64) (*SocialFriend, error)
		Page(page int64, rows int64) *defaultSocialFriendModel
		Select() ([]*SocialFriend, error)
		SelectWithTotal() ([]*SocialFriend, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*SocialFriend, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultSocialFriendModel
		Reinit() *defaultSocialFriendModel
		Dao() *dao.SqlxDao
	}

	defaultSocialFriendModel struct {
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

	SocialFriend struct {
		Id       int64  `db:"id"`
		UserId   int64  `db:"user_id"`   // 归属用户id
		FriendId int64  `db:"friend_id"` // 好友用户id
		SourceEm int64  `db:"source_em"` // 添加来源枚举
		Remark   string `db:"remark"`    // 备注
		PlatId   int64  `db:"plat_id"`   // 应用id
		CreateAt int64  `db:"create_at"` // 创建时间戳
		UpdateAt int64  `db:"update_at"` // 更新时间戳
		DeleteAt int64  `db:"delete_at"` // 删除时间戳
	}
)

func newSocialFriendModel(conn sqlx.SqlConn, platId int64) *defaultSocialFriendModel {
	dao := dao.NewSqlxDao(conn, "`social_friend`", defaultSocialFriendFields, true, "delete_at")
	dao.Plat(platId)
	return &defaultSocialFriendModel{
		conn:            conn,
		dao:             dao,
		table:           "`social_friend`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultSocialFriendModel) Ctx(ctx context.Context) *defaultSocialFriendModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultSocialFriendModel) WhereId(id int64) *defaultSocialFriendModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultSocialFriendModel) Where(whereStr string, whereData ...any) *defaultSocialFriendModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultSocialFriendModel) Alias(alias string) *defaultSocialFriendModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultSocialFriendModel) Field(field string) *defaultSocialFriendModel {
	m.dao.Field(field)
	return m
}
func (m *defaultSocialFriendModel) Order(order string) *defaultSocialFriendModel {
	m.dao.Order(order)
	return m
}
func (m *defaultSocialFriendModel) Limit(num int64) *defaultSocialFriendModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultSocialFriendModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultSocialFriendModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultSocialFriendModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultSocialFriendModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialFriendModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialFriendModel) Plat(id int64) *defaultSocialFriendModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultSocialFriendModel) Reinit() *defaultSocialFriendModel {
	m.dao.Reinit()
	return m
}
func (m *defaultSocialFriendModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultSocialFriendModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSocialFriendModel) Find() (*SocialFriend, error) {
	resp := &SocialFriend{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendModel) FindById(id int64) (*SocialFriend, error) {
	resp := &SocialFriend{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendModel) CacheFind(redis *redisd.Redisd) (*SocialFriend, error) {
	resp := &SocialFriend{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendModel) CacheFindById(redis *redisd.Redisd, id int64) (*SocialFriend, error) {
	resp := &SocialFriend{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultSocialFriendModel) Select() ([]*SocialFriend, error) {
	resp := make([]*SocialFriend, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendModel) SelectWithTotal() ([]*SocialFriend, int64, error) {
	resp := make([]*SocialFriend, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultSocialFriendModel) CacheSelect(redis *redisd.Redisd) ([]*SocialFriend, error) {
	resp := make([]*SocialFriend, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultSocialFriendModel) Page(page int64, size int64) *defaultSocialFriendModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultSocialFriendModel) Insert(data *SocialFriend) (int64, error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultSocialFriendModel) TxInsert(tx *sql.Tx, data *SocialFriend) (int64, error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultSocialFriendModel) Update(data map[string]any) (int64, error) {
	return m.dao.Update(data)
}
func (m *defaultSocialFriendModel) TxUpdate(tx *sql.Tx, data map[string]any) (int64, error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultSocialFriendModel) Save(data *SocialFriend) (int64, error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialFriendModel) TxSave(tx *sql.Tx, data *SocialFriend) (int64, error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultSocialFriendModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultSocialFriendModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
