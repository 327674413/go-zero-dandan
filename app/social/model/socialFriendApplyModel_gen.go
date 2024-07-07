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
	socialFriendApplyFieldNames          = builder.RawFieldNames(&SocialFriendApply{})
	socialFriendApplyRows                = strings.Join(socialFriendApplyFieldNames, ",")
	defaultSocialFriendApplyFields       = strings.Join(socialFriendApplyFieldNames, ",")
	socialFriendApplyRowsExpectAutoSet   = strings.Join(stringx.Remove(socialFriendApplyFieldNames, "`delete_at`"), ",")
	socialFriendApplyRowsWithPlaceHolder = strings.Join(stringx.Remove(socialFriendApplyFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	SocialFriendApply_Id           dao.TableField = "id"
	SocialFriendApply_UserId       dao.TableField = "user_id"
	SocialFriendApply_FriendUid    dao.TableField = "friend_uid"
	SocialFriendApply_ApplyLastMsg dao.TableField = "apply_last_msg"
	SocialFriendApply_ApplyStartAt dao.TableField = "apply_start_at"
	SocialFriendApply_ApplyLastAt  dao.TableField = "apply_last_at"
	SocialFriendApply_OperateMsg   dao.TableField = "operate_msg"
	SocialFriendApply_OperateAt    dao.TableField = "operate_at"
	SocialFriendApply_StateEm      dao.TableField = "state_em"
	SocialFriendApply_Remark       dao.TableField = "remark"
	SocialFriendApply_IsRead       dao.TableField = "is_read"
	SocialFriendApply_SourceEm     dao.TableField = "source_em"
	SocialFriendApply_PlatId       dao.TableField = "plat_id"
	SocialFriendApply_Content      dao.TableField = "content"
	SocialFriendApply_CreateAt     dao.TableField = "create_at"
	SocialFriendApply_UpdateAt     dao.TableField = "update_at"
	SocialFriendApply_DeleteAt     dao.TableField = "delete_at"
)

type (
	socialFriendApplyModel interface {
		Insert(data *SocialFriendApply) (effectRow int64, err error)
		TxInsert(tx *sql.Tx, data *SocialFriendApply) (effectRow int64, err error)
		Update(data map[dao.TableField]any) (effectRow int64, err error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error)
		Save(data *SocialFriendApply) (effectRow int64, err error)
		TxSave(tx *sql.Tx, data *SocialFriendApply) (effectRow int64, err error)
		Delete(ctx context.Context, id string) error
		Field(field string) *defaultSocialFriendApplyModel
		Except(fields ...string) *defaultSocialFriendApplyModel
		Alias(alias string) *defaultSocialFriendApplyModel
		Where(whereStr string, whereData ...any) *defaultSocialFriendApplyModel
		WhereId(id string) *defaultSocialFriendApplyModel
		Order(order string) *defaultSocialFriendApplyModel
		Limit(num int64) *defaultSocialFriendApplyModel
		Plat(id string) *defaultSocialFriendApplyModel
		Find() (*SocialFriendApply, error)
		FindById(id string) (*SocialFriendApply, error)
		CacheFind(redis *redisd.Redisd) (*SocialFriendApply, error)
		CacheFindById(redis *redisd.Redisd, id string) (*SocialFriendApply, error)
		Page(page int64, rows int64) *defaultSocialFriendApplyModel
		Total() (total int64, err error)
		Select() ([]*SocialFriendApply, error)
		SelectWithTotal() ([]*SocialFriendApply, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*SocialFriendApply, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultSocialFriendApplyModel
		Reinit() *defaultSocialFriendApplyModel
		Dao() *dao.SqlxDao
	}

	defaultSocialFriendApplyModel struct {
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

	SocialFriendApply struct {
		Id           string         `db:"id"`
		UserId       string         `db:"user_id"`        // 发起人id
		FriendUid    string         `db:"friend_uid"`     // 对方id
		ApplyLastMsg string         `db:"apply_last_msg"` // 最后一次申请验证信息
		ApplyStartAt int64          `db:"apply_start_at"` // 申请开始时间戳，用于过滤通过之前的历史申请
		ApplyLastAt  int64          `db:"apply_last_at"`  // 最后一次申请时间，用来好申请列表排序用
		OperateMsg   string         `db:"operate_msg"`    // 处理时填写的内容
		OperateAt    int64          `db:"operate_at"`     // 处理时间戳
		StateEm      int64          `db:"state_em"`       // 申请状态
		Remark       string         `db:"remark"`         // 备注
		IsRead       int64          `db:"is_read"`        // friend_uid被申请人是否已读
		SourceEm     int64          `db:"source_em"`      // 来源枚举
		PlatId       string         `db:"plat_id"`        // 应用id
		Content      sql.NullString `db:"content"`        // 添加沟通记录
		CreateAt     int64          `db:"create_at"`      // 创建时间戳
		UpdateAt     int64          `db:"update_at"`      // 更新时间戳
		DeleteAt     int64          `db:"delete_at"`      // 删除时间戳
	}
)

// NewSocialFriendApplyModel returns a model for the database table.
func NewSocialFriendApplyModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) SocialFriendApplyModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customSocialFriendApplyModel{
		defaultSocialFriendApplyModel: newSocialFriendApplyModel(ctxOrNil, conn, platid),
		softDeletable:                 softDeletableSocialFriendApply,
	}
}
func newSocialFriendApplyModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultSocialFriendApplyModel {
	dao := dao.NewSqlxDao(conn, "`social_friend_apply`", defaultSocialFriendApplyFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultSocialFriendApplyModel{
		ctx:             ctx,
		conn:            conn,
		dao:             dao,
		table:           "`social_friend_apply`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultSocialFriendApplyModel) Ctx(ctx context.Context) *defaultSocialFriendApplyModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultSocialFriendApplyModel) WhereId(id string) *defaultSocialFriendApplyModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultSocialFriendApplyModel) Where(whereStr string, whereData ...any) *defaultSocialFriendApplyModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultSocialFriendApplyModel) Alias(alias string) *defaultSocialFriendApplyModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultSocialFriendApplyModel) Field(field string) *defaultSocialFriendApplyModel {
	m.dao.Field(field)
	return m
}
func (m *defaultSocialFriendApplyModel) Except(fields ...string) *defaultSocialFriendApplyModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultSocialFriendApplyModel) Order(order string) *defaultSocialFriendApplyModel {
	m.dao.Order(order)
	return m
}
func (m *defaultSocialFriendApplyModel) Limit(num int64) *defaultSocialFriendApplyModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultSocialFriendApplyModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultSocialFriendApplyModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultSocialFriendApplyModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultSocialFriendApplyModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialFriendApplyModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialFriendApplyModel) Plat(id string) *defaultSocialFriendApplyModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultSocialFriendApplyModel) Reinit() *defaultSocialFriendApplyModel {
	m.dao.Reinit()
	return m
}
func (m *defaultSocialFriendApplyModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultSocialFriendApplyModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSocialFriendApplyModel) Find() (*SocialFriendApply, error) {
	resp := &SocialFriendApply{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendApplyModel) FindById(id string) (*SocialFriendApply, error) {
	resp := &SocialFriendApply{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendApplyModel) CacheFind(redis *redisd.Redisd) (*SocialFriendApply, error) {
	resp := &SocialFriendApply{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendApplyModel) CacheFindById(redis *redisd.Redisd, id string) (*SocialFriendApply, error) {
	resp := &SocialFriendApply{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendApplyModel) Total() (total int64, err error) {
	return m.dao.Total()
}
func (m *defaultSocialFriendApplyModel) Select() ([]*SocialFriendApply, error) {
	resp := make([]*SocialFriendApply, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialFriendApplyModel) SelectWithTotal() ([]*SocialFriendApply, int64, error) {
	resp := make([]*SocialFriendApply, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultSocialFriendApplyModel) CacheSelect(redis *redisd.Redisd) ([]*SocialFriendApply, error) {
	resp := make([]*SocialFriendApply, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultSocialFriendApplyModel) Page(page int64, size int64) *defaultSocialFriendApplyModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultSocialFriendApplyModel) Insert(data *SocialFriendApply) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultSocialFriendApplyModel) TxInsert(tx *sql.Tx, data *SocialFriendApply) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultSocialFriendApplyModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultSocialFriendApplyModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultSocialFriendApplyModel) Save(data *SocialFriendApply) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialFriendApplyModel) TxSave(tx *sql.Tx, data *SocialFriendApply) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultSocialFriendApplyModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultSocialFriendApplyModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
