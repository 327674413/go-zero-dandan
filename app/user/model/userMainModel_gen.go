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
	userMainFieldNames          = builder.RawFieldNames(&UserMain{})
	userMainRows                = strings.Join(userMainFieldNames, ",")
	defaultUserMainFields       = strings.Join(userMainFieldNames, ",")
	userMainRowsExpectAutoSet   = strings.Join(stringx.Remove(userMainFieldNames, "`delete_at`"), ",")
	userMainRowsWithPlaceHolder = strings.Join(stringx.Remove(userMainFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	UserMain_Id        dao.TableField = "id"
	UserMain_UnionId   dao.TableField = "union_id"
	UserMain_StateEm   dao.TableField = "state_em"
	UserMain_Account   dao.TableField = "account"
	UserMain_Password  dao.TableField = "password"
	UserMain_Code      dao.TableField = "code"
	UserMain_Nickname  dao.TableField = "nickname"
	UserMain_Phone     dao.TableField = "phone"
	UserMain_PhoneArea dao.TableField = "phone_area"
	UserMain_Email     dao.TableField = "email"
	UserMain_AvatarImg dao.TableField = "avatar_img"
	UserMain_Signature dao.TableField = "signature"
	UserMain_SexEm     dao.TableField = "sex_em"
	UserMain_PlatId    dao.TableField = "plat_id"
	UserMain_CreateAt  dao.TableField = "create_at"
	UserMain_UpdateAt  dao.TableField = "update_at"
	UserMain_DeleteAt  dao.TableField = "delete_at"
)

type (
	userMainModel interface {
		Insert(data *UserMain) (effectRow int64, err error)
		TxInsert(tx *sql.Tx, data *UserMain) (effectRow int64, err error)
		Update(data map[dao.TableField]any) (effectRow int64, err error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error)
		Save(data *UserMain) (effectRow int64, err error)
		TxSave(tx *sql.Tx, data *UserMain) (effectRow int64, err error)
		Delete(ctx context.Context, id string) error
		Field(field string) *defaultUserMainModel
		Except(fields ...string) *defaultUserMainModel
		Alias(alias string) *defaultUserMainModel
		Where(whereStr string, whereData ...any) *defaultUserMainModel
		WhereId(id string) *defaultUserMainModel
		Order(order string) *defaultUserMainModel
		Limit(num int64) *defaultUserMainModel
		Plat(id string) *defaultUserMainModel
		Find() (*UserMain, error)
		FindById(id string) (*UserMain, error)
		CacheFind(redis *redisd.Redisd) (*UserMain, error)
		CacheFindById(redis *redisd.Redisd, id string) (*UserMain, error)
		Page(page int64, rows int64) *defaultUserMainModel
		Total() (total int64, err error)
		Select() ([]*UserMain, error)
		SelectWithTotal() ([]*UserMain, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*UserMain, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultUserMainModel
		Reinit() *defaultUserMainModel
		Dao() *dao.SqlxDao
	}

	defaultUserMainModel struct {
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

	UserMain struct {
		Id        string `db:"id" json:"id"`
		UnionId   string `db:"union_id" json:"unionId"`     // 平台层用户唯一表示
		StateEm   int64  `db:"state_em" json:"stateEm"`     // 用户状态枚举
		Account   string `db:"account" json:"account"`      // 登录账号
		Password  string `db:"password" json:"password"`    // 登录密码
		Code      string `db:"code" json:"code"`            // 用户编号
		Nickname  string `db:"nickname" json:"nickname"`    // 昵称
		Phone     string `db:"phone" json:"phone"`          // 手机号(已验证)
		PhoneArea string `db:"phone_area" json:"phoneArea"` // 手机区号
		Email     string `db:"email" json:"email"`          // 邮箱地址
		AvatarImg string `db:"avatar_img" json:"avatarImg"` // 头像
		Signature string `db:"signature" json:"signature"`  // 个性签名
		SexEm     int64  `db:"sex_em" json:"sexEm"`         // 性别枚举
		PlatId    string `db:"plat_id" json:"platId"`       // 应用id
		CreateAt  int64  `db:"create_at" json:"createAt"`   // 创建时间戳
		UpdateAt  int64  `db:"update_at" json:"updateAt"`   // 更新时间戳
		DeleteAt  int64  `db:"delete_at" json:"deleteAt"`   // 删除时间戳
	}
)

// NewUserMainModel returns a model for the database table.
func NewUserMainModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) UserMainModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customUserMainModel{
		defaultUserMainModel: newUserMainModel(ctxOrNil, conn, platid),
		softDeletable:        softDeletableUserMain,
	}
}
func newUserMainModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultUserMainModel {
	dao := dao.NewSqlxDao(conn, "`user_main`", defaultUserMainFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultUserMainModel{
		ctx:             ctx,
		conn:            conn,
		dao:             dao,
		table:           "`user_main`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultUserMainModel) Ctx(ctx context.Context) *defaultUserMainModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultUserMainModel) WhereId(id string) *defaultUserMainModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultUserMainModel) Where(whereStr string, whereData ...any) *defaultUserMainModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultUserMainModel) Alias(alias string) *defaultUserMainModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultUserMainModel) Field(field string) *defaultUserMainModel {
	m.dao.Field(field)
	return m
}
func (m *defaultUserMainModel) Except(fields ...string) *defaultUserMainModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultUserMainModel) Order(order string) *defaultUserMainModel {
	m.dao.Order(order)
	return m
}
func (m *defaultUserMainModel) Limit(num int64) *defaultUserMainModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultUserMainModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultUserMainModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultUserMainModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultUserMainModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserMainModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultUserMainModel) Plat(id string) *defaultUserMainModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultUserMainModel) Reinit() *defaultUserMainModel {
	m.dao.Reinit()
	return m
}
func (m *defaultUserMainModel) Dao() *dao.SqlxDao {
	return m.dao
}

func (m *defaultUserMainModel) Find() (*UserMain, error) {
	resp := &UserMain{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) FindById(id string) (*UserMain, error) {
	resp := &UserMain{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) CacheFind(redis *redisd.Redisd) (*UserMain, error) {
	resp := &UserMain{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) CacheFindById(redis *redisd.Redisd, id string) (*UserMain, error) {
	resp := &UserMain{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) Total() (total int64, err error) {
	return m.dao.Total()
}
func (m *defaultUserMainModel) Select() ([]*UserMain, error) {
	resp := make([]*UserMain, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultUserMainModel) SelectWithTotal() ([]*UserMain, int64, error) {
	resp := make([]*UserMain, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultUserMainModel) CacheSelect(redis *redisd.Redisd) ([]*UserMain, error) {
	resp := make([]*UserMain, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserMainModel) Page(page int64, size int64) *defaultUserMainModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultUserMainModel) Insert(data *UserMain) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultUserMainModel) TxInsert(tx *sql.Tx, data *UserMain) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultUserMainModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultUserMainModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultUserMainModel) Save(data *UserMain) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultUserMainModel) TxSave(tx *sql.Tx, data *UserMain) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultUserMainModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultUserMainModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
