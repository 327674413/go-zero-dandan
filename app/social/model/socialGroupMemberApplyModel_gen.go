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
	socialGroupMemberApplyFieldNames          = builder.RawFieldNames(&SocialGroupMemberApply{})
	socialGroupMemberApplyRows                = strings.Join(socialGroupMemberApplyFieldNames, ",")
	defaultSocialGroupMemberApplyFields       = strings.Join(socialGroupMemberApplyFieldNames, ",")
	socialGroupMemberApplyRowsExpectAutoSet   = strings.Join(stringx.Remove(socialGroupMemberApplyFieldNames, "`delete_at`"), ",")
	socialGroupMemberApplyRowsWithPlaceHolder = strings.Join(stringx.Remove(socialGroupMemberApplyFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	SocialGroupMemberApply_Id             dao.TableField = "id"
	SocialGroupMemberApply_GroupId        dao.TableField = "group_id"
	SocialGroupMemberApply_UserId         dao.TableField = "user_id"
	SocialGroupMemberApply_ApplyMsg       dao.TableField = "apply_msg"
	SocialGroupMemberApply_ApplyAt        dao.TableField = "apply_at"
	SocialGroupMemberApply_JoinSourceEm   dao.TableField = "join_source_em"
	SocialGroupMemberApply_InviteUid      dao.TableField = "invite_uid"
	SocialGroupMemberApply_OperateUid     dao.TableField = "operate_uid"
	SocialGroupMemberApply_OperateAt      dao.TableField = "operate_at"
	SocialGroupMemberApply_OperateStateEm dao.TableField = "operate_state_em"
	SocialGroupMemberApply_OperateMsg     dao.TableField = "operate_msg"
	SocialGroupMemberApply_Remark         dao.TableField = "remark"
	SocialGroupMemberApply_PlatId         dao.TableField = "plat_id"
	SocialGroupMemberApply_CreateAt       dao.TableField = "create_at"
	SocialGroupMemberApply_UpdateAt       dao.TableField = "update_at"
	SocialGroupMemberApply_DeleteAt       dao.TableField = "delete_at"
)

type (
	socialGroupMemberApplyModel interface {
		Delete(id ...string) (effectRow int64, danErr error)
		TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error)
		Insert(data *SocialGroupMemberApply) (effectRow int64, danErr error)
		TxInsert(tx *sql.Tx, data *SocialGroupMemberApply) (effectRow int64, danErr error)
		Update(data map[dao.TableField]any) (effectRow int64, danErr error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, danErr error)
		Save(data *SocialGroupMemberApply) (effectRow int64, danErr error)
		TxSave(tx *sql.Tx, data *SocialGroupMemberApply) (effectRow int64, danErr error)
		Field(field string) *defaultSocialGroupMemberApplyModel
		Except(fields ...string) *defaultSocialGroupMemberApplyModel
		Alias(alias string) *defaultSocialGroupMemberApplyModel
		LeftJoin(joinTable string) *defaultSocialGroupMemberApplyModel
		RightJoin(joinTable string) *defaultSocialGroupMemberApplyModel
		InnerJoin(joinTable string) *defaultSocialGroupMemberApplyModel
		Where(whereStr string, whereData ...any) *defaultSocialGroupMemberApplyModel
		WhereId(id string) *defaultSocialGroupMemberApplyModel
		Order(order string) *defaultSocialGroupMemberApplyModel
		Limit(num int64) *defaultSocialGroupMemberApplyModel
		Plat(id string) *defaultSocialGroupMemberApplyModel
		Find() (*SocialGroupMemberApply, error)
		FindById(id string) (data *SocialGroupMemberApply, danErr error)
		CacheFind(redis *redisd.Redisd) (data *SocialGroupMemberApply, danErr error)
		CacheFindById(redis *redisd.Redisd, id string) (data *SocialGroupMemberApply, danErr error)
		Page(page int64, rows int64) *defaultSocialGroupMemberApplyModel
		Total() (total int64, danErr error)
		Select() (dataList []*SocialGroupMemberApply, danErr error)
		SelectWithTotal() (dataList []*SocialGroupMemberApply, total int64, danErr error)
		CacheSelect(redis *redisd.Redisd) (dataList []*SocialGroupMemberApply, danErr error)
		Count() (total int64, danErr error)
		Inc(field string, num int) (effectRow int64, danErr error)
		Dec(field string, num int) (effectRow int64, danErr error)
		StartTrans() (tx *sql.Tx, danErr error)
		Commit(tx *sql.Tx) (danErr error)
		Rollback(tx *sql.Tx) (danErr error)
		Ctx(ctx context.Context) *defaultSocialGroupMemberApplyModel
		Reinit() *defaultSocialGroupMemberApplyModel
		Dao() *dao.SqlxDao
	}

	defaultSocialGroupMemberApplyModel struct {
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

	SocialGroupMemberApply struct {
		Id             string `db:"id" json:"id"`
		GroupId        string `db:"group_id" json:"groupId"`                // 群组id
		UserId         string `db:"user_id" json:"userId"`                  // 用户id
		ApplyMsg       string `db:"apply_msg" json:"applyMsg"`              // 申请内容
		ApplyAt        int64  `db:"apply_at" json:"applyAt"`                // 申请时间
		JoinSourceEm   int64  `db:"join_source_em" json:"joinSourceEm"`     // 加入方式
		InviteUid      string `db:"invite_uid" json:"inviteUid"`            // 邀请人用户id
		OperateUid     string `db:"operate_uid" json:"operateUid"`          // 操作人用户id
		OperateAt      int64  `db:"operate_at" json:"operateAt"`            // 操作时间
		OperateStateEm int64  `db:"operate_state_em" json:"operateStateEm"` // 处理结果
		OperateMsg     string `db:"operate_msg" json:"operateMsg"`
		Remark         string `db:"remark" json:"remark"`      // 备注
		PlatId         string `db:"plat_id" json:"platId"`     // 应用id
		CreateAt       int64  `db:"create_at" json:"createAt"` // 创建时间戳
		UpdateAt       int64  `db:"update_at" json:"updateAt"` // 更新时间戳
		DeleteAt       int64  `db:"delete_at" json:"deleteAt"` // 删除时间戳
	}
)

// NewSocialGroupMemberApplyModel returns a model for the database table.
func NewSocialGroupMemberApplyModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) SocialGroupMemberApplyModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customSocialGroupMemberApplyModel{
		defaultSocialGroupMemberApplyModel: newSocialGroupMemberApplyModel(ctxOrNil, conn, platid),
		softDeletable:                      softDeletableSocialGroupMemberApply,
	}
}
func newSocialGroupMemberApplyModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultSocialGroupMemberApplyModel {
	dao := dao.NewSqlxDao(conn, "`social_group_member_apply`", defaultSocialGroupMemberApplyFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultSocialGroupMemberApplyModel{
		ctx:             ctx,
		conn:            conn,
		dao:             dao,
		table:           "`social_group_member_apply`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultSocialGroupMemberApplyModel) Ctx(ctx context.Context) *defaultSocialGroupMemberApplyModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) WhereId(id string) *defaultSocialGroupMemberApplyModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultSocialGroupMemberApplyModel) Where(whereStr string, whereData ...any) *defaultSocialGroupMemberApplyModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultSocialGroupMemberApplyModel) Alias(alias string) *defaultSocialGroupMemberApplyModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) LeftJoin(joinTable string) *defaultSocialGroupMemberApplyModel {
	m.dao.LeftJoin(joinTable)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) RightJoin(joinTable string) *defaultSocialGroupMemberApplyModel {
	m.dao.RightJoin(joinTable)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) InnerJoin(joinTable string) *defaultSocialGroupMemberApplyModel {
	m.dao.InnerJoin(joinTable)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) Field(field string) *defaultSocialGroupMemberApplyModel {
	m.dao.Field(field)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) Except(fields ...string) *defaultSocialGroupMemberApplyModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) Order(order string) *defaultSocialGroupMemberApplyModel {
	m.dao.Order(order)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) Limit(num int64) *defaultSocialGroupMemberApplyModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultSocialGroupMemberApplyModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultSocialGroupMemberApplyModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultSocialGroupMemberApplyModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialGroupMemberApplyModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultSocialGroupMemberApplyModel) Plat(id string) *defaultSocialGroupMemberApplyModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) Reinit() *defaultSocialGroupMemberApplyModel {
	m.dao.Reinit()
	return m
}
func (m *defaultSocialGroupMemberApplyModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultSocialGroupMemberApplyModel) Find() (*SocialGroupMemberApply, error) {
	resp := &SocialGroupMemberApply{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberApplyModel) FindById(id string) (*SocialGroupMemberApply, error) {
	resp := &SocialGroupMemberApply{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberApplyModel) CacheFind(redis *redisd.Redisd) (*SocialGroupMemberApply, error) {
	resp := &SocialGroupMemberApply{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberApplyModel) CacheFindById(redis *redisd.Redisd, id string) (*SocialGroupMemberApply, error) {
	resp := &SocialGroupMemberApply{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberApplyModel) Total() (total int64, danErr error) {
	return m.dao.Total()
}
func (m *defaultSocialGroupMemberApplyModel) Select() ([]*SocialGroupMemberApply, error) {
	resp := make([]*SocialGroupMemberApply, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultSocialGroupMemberApplyModel) SelectWithTotal() ([]*SocialGroupMemberApply, int64, error) {
	resp := make([]*SocialGroupMemberApply, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultSocialGroupMemberApplyModel) CacheSelect(redis *redisd.Redisd) ([]*SocialGroupMemberApply, error) {
	resp := make([]*SocialGroupMemberApply, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultSocialGroupMemberApplyModel) Page(page int64, size int64) *defaultSocialGroupMemberApplyModel {
	m.dao.Page(page, size)
	return m
}
func (m *defaultSocialGroupMemberApplyModel) Insert(data *SocialGroupMemberApply) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultSocialGroupMemberApplyModel) TxInsert(tx *sql.Tx, data *SocialGroupMemberApply) (effectRow int64, danErr error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}
func (m *defaultSocialGroupMemberApplyModel) Delete(id ...string) (effectRow int64, danErr error) {
	return m.dao.Delete(id...)
}
func (m *defaultSocialGroupMemberApplyModel) TxDelete(tx *sql.Tx, id ...string) (effectRow int64, danErr error) {
	return m.dao.TxDelete(tx, id...)
}
func (m *defaultSocialGroupMemberApplyModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultSocialGroupMemberApplyModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultSocialGroupMemberApplyModel) Save(data *SocialGroupMemberApply) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialGroupMemberApplyModel) TxSave(tx *sql.Tx, data *SocialGroupMemberApply) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultSocialGroupMemberApplyModel) StartTrans() (tx *sql.Tx, danErr error) {
	return dao.StartTrans(m.ctx, m.conn)
}
func (m *defaultSocialGroupMemberApplyModel) Commit(tx *sql.Tx) (danErr error) {
	return dao.Commit(tx)
}
func (m *defaultSocialGroupMemberApplyModel) Rollback(tx *sql.Tx) (danErr error) {
	return dao.Rollback(tx)
}
func (m *defaultSocialGroupMemberApplyModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultSocialGroupMemberApplyModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
