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

const (
	SocialFriend_Id          dao.TableField = "id"
	SocialFriend_UserId      dao.TableField = "user_id"
	SocialFriend_FriendUid   dao.TableField = "friend_uid"
	SocialFriend_FriendName  dao.TableField = "friend_name"
	SocialFriend_FriendAlias dao.TableField = "friend_alias"
	SocialFriend_FriendIcon  dao.TableField = "friend_icon"
	SocialFriend_SourceEm    dao.TableField = "source_em"
	SocialFriend_StateEm     dao.TableField = "state_em"
	SocialFriend_Remark      dao.TableField = "remark"
	SocialFriend_PlatId      dao.TableField = "plat_id"
	SocialFriend_CreateAt    dao.TableField = "create_at"
	SocialFriend_UpdateAt    dao.TableField = "update_at"
	SocialFriend_DeleteAt    dao.TableField = "delete_at"
)

type (
	socialFriendModel interface {
		Delete(id ...string) (effectRow int64, danErr error)
		TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error)
		Insert(data *SocialFriend) (effectRow int64, danErr error)
		TxInsert(tx *sql.Tx, data *SocialFriend) (effectRow int64, danErr error)
		Update(data map[dao.TableField]any) (effectRow int64, danErr error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, danErr error)
		Save(data *SocialFriend) (effectRow int64, danErr error)
		TxSave(tx *sql.Tx, data *SocialFriend) (effectRow int64, danErr error)
		Field(field string) *defaultSocialFriendModel
		Except(fields ...string) *defaultSocialFriendModel
		Alias(alias string) *defaultSocialFriendModel
		Where(whereStr string, whereData ...any) *defaultSocialFriendModel
		WhereId(id string) *defaultSocialFriendModel
		Order(order string) *defaultSocialFriendModel
		Limit(num int64) *defaultSocialFriendModel
		Plat(id string) *defaultSocialFriendModel
		Find() (*SocialFriend, error)
		FindById(id string) (data *SocialFriend, danErr error)
		CacheFind(redis *redisd.Redisd) (data *SocialFriend, danErr error)
		CacheFindById(redis *redisd.Redisd, id string) (data *SocialFriend, danErr error)
		Page(page int64, rows int64) *defaultSocialFriendModel
		Total() (total int64, danErr error)
		Select() (dataList []*SocialFriend, danErr error)
		SelectWithTotal() (dataList []*SocialFriend, total int64, danErr error)
		CacheSelect(redis *redisd.Redisd) (dataList []*SocialFriend, danErr error)
		Count() (total int64, danErr error)
		Inc(field string, num int) (effectRow int64, danErr error)
		Dec(field string, num int) (effectRow int64, danErr error)
		StartTrans() (tx *sql.Tx, danErr error)
		Commit(tx *sql.Tx) (danErr error)
		Rollback(tx *sql.Tx) (danErr error)
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
		platId          string
		whereData       []any
		err             error
		ctx             context.Context
	}

	SocialFriend struct {
		Id          string `db:"id" json:"id"`
		UserId      string `db:"user_id" json:"userId"`           // 归属用户id
		FriendUid   string `db:"friend_uid" json:"friendUid"`     // 好友用户id
		FriendName  string `db:"friend_name" json:"friendName"`   // 冗余好友名称
		FriendAlias string `db:"friend_alias" json:"friendAlias"` // 好友别名备注
		FriendIcon  string `db:"friend_icon" json:"friendIcon"`   // 冗余好友头像
		SourceEm    int64  `db:"source_em" json:"sourceEm"`       // 添加来源枚举
		StateEm     int64  `db:"state_em" json:"stateEm"`         // 好友状态
		Remark      string `db:"remark" json:"remark"`            // 备注
		PlatId      string `db:"plat_id" json:"platId"`           // 应用id
		CreateAt    int64  `db:"create_at" json:"createAt"`       // 创建时间戳
		UpdateAt    int64  `db:"update_at" json:"updateAt"`       // 更新时间戳
		DeleteAt    int64  `db:"delete_at" json:"deleteAt"`       // 删除时间戳
	}
)

// NewSocialFriendModel returns a model for the database table.
func NewSocialFriendModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) SocialFriendModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customSocialFriendModel{
		defaultSocialFriendModel: newSocialFriendModel(ctxOrNil, conn, platid),
		softDeletable:            softDeletableSocialFriend,
	}
}
func newSocialFriendModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultSocialFriendModel {
	dao := dao.NewSqlxDao(conn, "`social_friend`", defaultSocialFriendFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultSocialFriendModel{
		ctx:             ctx,
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
func (m *defaultSocialFriendModel) WhereId(id string) *defaultSocialFriendModel {
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
func (m *defaultSocialFriendModel) Except(fields ...string) *defaultSocialFriendModel {
	m.dao.Except(fields...)
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
func (m *defaultSocialFriendModel) Plat(id string) *defaultSocialFriendModel {
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
func (m *defaultSocialFriendModel) FindById(id string) (*SocialFriend, error) {
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
func (m *defaultSocialFriendModel) CacheFindById(redis *redisd.Redisd, id string) (*SocialFriend, error) {
	resp := &SocialFriend{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendModel) Total() (total int64, danErr error) {
	return m.dao.Total()
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
func (m *defaultSocialFriendModel) Insert(data *SocialFriend) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultSocialFriendModel) TxInsert(tx *sql.Tx, data *SocialFriend) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}
func (m *defaultSocialFriendModel) Delete(id ...string) (effectRow int64, danErr error) {
	return m.dao.Delete(id...)
}
func (m *defaultSocialFriendModel) TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error) {
	return m.dao.TxDelete(tx, id...)
}
func (m *defaultSocialFriendModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultSocialFriendModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultSocialFriendModel) Save(data *SocialFriend) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialFriendModel) TxSave(tx *sql.Tx, data *SocialFriend) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialFriendModel) StartTrans() (tx *sql.Tx, danErr error) {
	return dao.StartTrans(m.conn, m.ctx)
}
func (m *defaultSocialFriendModel) Commit(tx *sql.Tx) (danErr error) {
	return dao.Commit(tx)
}
func (m *defaultSocialFriendModel) Rollback(tx *sql.Tx) (danErr error) {
	return dao.Rollback(tx)
}
func (m *defaultSocialFriendModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultSocialFriendModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
