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
	socialGroupMemberFieldNames          = builder.RawFieldNames(&SocialGroupMember{})
	socialGroupMemberRows                = strings.Join(socialGroupMemberFieldNames, ",")
	defaultSocialGroupMemberFields       = strings.Join(socialGroupMemberFieldNames, ",")
	socialGroupMemberRowsExpectAutoSet   = strings.Join(stringx.Remove(socialGroupMemberFieldNames, "`delete_at`"), ",")
	socialGroupMemberRowsWithPlaceHolder = strings.Join(stringx.Remove(socialGroupMemberFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	SocialGroupMember_Id           dao.TableField = "id"
	SocialGroupMember_GroupId      dao.TableField = "group_id"
	SocialGroupMember_UserId       dao.TableField = "user_id"
	SocialGroupMember_RoleLevel    dao.TableField = "role_level"
	SocialGroupMember_JoinAt       dao.TableField = "join_at"
	SocialGroupMember_JoinSourceEm dao.TableField = "join_source_em"
	SocialGroupMember_InviteUid    dao.TableField = "invite_uid"
	SocialGroupMember_OperateUid   dao.TableField = "operate_uid"
	SocialGroupMember_Remark       dao.TableField = "remark"
	SocialGroupMember_PlatId       dao.TableField = "plat_id"
	SocialGroupMember_CreateAt     dao.TableField = "create_at"
	SocialGroupMember_UpdateAt     dao.TableField = "update_at"
	SocialGroupMember_DeleteAt     dao.TableField = "delete_at"
)

type (
	socialGroupMemberModel interface {
		Insert(data *SocialGroupMember) (effectRow int64, err error)
		TxInsert(tx *sql.Tx, data *SocialGroupMember) (effectRow int64, err error)
		Update(data map[dao.TableField]any) (effectRow int64, err error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error)
		Save(data *SocialGroupMember) (effectRow int64, err error)
		TxSave(tx *sql.Tx, data *SocialGroupMember) (effectRow int64, err error)
		Delete(ctx context.Context, id string) error
		Field(field string) *defaultSocialGroupMemberModel
		Except(fields ...string) *defaultSocialGroupMemberModel
		Alias(alias string) *defaultSocialGroupMemberModel
		Where(whereStr string, whereData ...any) *defaultSocialGroupMemberModel
		WhereId(id string) *defaultSocialGroupMemberModel
		Order(order string) *defaultSocialGroupMemberModel
		Limit(num int64) *defaultSocialGroupMemberModel
		Plat(id string) *defaultSocialGroupMemberModel
		Find() (*SocialGroupMember, error)
		FindById(id string) (*SocialGroupMember, error)
		CacheFind(redis *redisd.Redisd) (*SocialGroupMember, error)
		CacheFindById(redis *redisd.Redisd, id string) (*SocialGroupMember, error)
		Page(page int64, rows int64) *defaultSocialGroupMemberModel
		Select() ([]*SocialGroupMember, error)
		SelectWithTotal() ([]*SocialGroupMember, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*SocialGroupMember, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultSocialGroupMemberModel
		Reinit() *defaultSocialGroupMemberModel
		Dao() *dao.SqlxDao
	}

	defaultSocialGroupMemberModel struct {
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

	SocialGroupMember struct {
		Id           string `db:"id"`
		GroupId      string `db:"group_id"`       // 群组id
		UserId       string `db:"user_id"`        // 用户id
		RoleLevel    int64  `db:"role_level"`     // 权限等级
		JoinAt       int64  `db:"join_at"`        // 加入时间
		JoinSourceEm int64  `db:"join_source_em"` // 加入方式
		InviteUid    string `db:"invite_uid"`     // 邀请人用户id
		OperateUid   string `db:"operate_uid"`    // 操作人用户id
		Remark       string `db:"remark"`         // 备注
		PlatId       string `db:"plat_id"`        // 应用id
		CreateAt     int64  `db:"create_at"`      // 创建时间戳
		UpdateAt     int64  `db:"update_at"`      // 更新时间戳
		DeleteAt     int64  `db:"delete_at"`      // 删除时间戳
	}
)

func newSocialGroupMemberModel(conn sqlx.SqlConn, platId string) *defaultSocialGroupMemberModel {
	dao := dao.NewSqlxDao(conn, "`social_group_member`", defaultSocialGroupMemberFields, true, "delete_at")
	dao.Plat(platId)
	return &defaultSocialGroupMemberModel{
		conn:            conn,
		dao:             dao,
		table:           "`social_group_member`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultSocialGroupMemberModel) Ctx(ctx context.Context) *defaultSocialGroupMemberModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultSocialGroupMemberModel) WhereId(id string) *defaultSocialGroupMemberModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultSocialGroupMemberModel) Where(whereStr string, whereData ...any) *defaultSocialGroupMemberModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultSocialGroupMemberModel) Alias(alias string) *defaultSocialGroupMemberModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultSocialGroupMemberModel) Field(field string) *defaultSocialGroupMemberModel {
	m.dao.Field(field)
	return m
}
func (m *defaultSocialGroupMemberModel) Except(fields ...string) *defaultSocialGroupMemberModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultSocialGroupMemberModel) Order(order string) *defaultSocialGroupMemberModel {
	m.dao.Order(order)
	return m
}
func (m *defaultSocialGroupMemberModel) Limit(num int64) *defaultSocialGroupMemberModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultSocialGroupMemberModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultSocialGroupMemberModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultSocialGroupMemberModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultSocialGroupMemberModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialGroupMemberModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialGroupMemberModel) Plat(id string) *defaultSocialGroupMemberModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultSocialGroupMemberModel) Reinit() *defaultSocialGroupMemberModel {
	m.dao.Reinit()
	return m
}
func (m *defaultSocialGroupMemberModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultSocialGroupMemberModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSocialGroupMemberModel) Find() (*SocialGroupMember, error) {
	resp := &SocialGroupMember{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberModel) FindById(id string) (*SocialGroupMember, error) {
	resp := &SocialGroupMember{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberModel) CacheFind(redis *redisd.Redisd) (*SocialGroupMember, error) {
	resp := &SocialGroupMember{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberModel) CacheFindById(redis *redisd.Redisd, id string) (*SocialGroupMember, error) {
	resp := &SocialGroupMember{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultSocialGroupMemberModel) Select() ([]*SocialGroupMember, error) {
	resp := make([]*SocialGroupMember, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberModel) SelectWithTotal() ([]*SocialGroupMember, int64, error) {
	resp := make([]*SocialGroupMember, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultSocialGroupMemberModel) CacheSelect(redis *redisd.Redisd) ([]*SocialGroupMember, error) {
	resp := make([]*SocialGroupMember, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultSocialGroupMemberModel) Page(page int64, size int64) *defaultSocialGroupMemberModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultSocialGroupMemberModel) Insert(data *SocialGroupMember) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultSocialGroupMemberModel) TxInsert(tx *sql.Tx, data *SocialGroupMember) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultSocialGroupMemberModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultSocialGroupMemberModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultSocialGroupMemberModel) Save(data *SocialGroupMember) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialGroupMemberModel) TxSave(tx *sql.Tx, data *SocialGroupMember) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultSocialGroupMemberModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultSocialGroupMemberModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
