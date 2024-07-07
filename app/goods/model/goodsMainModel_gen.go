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
	goodsMainFieldNames          = builder.RawFieldNames(&GoodsMain{})
	goodsMainRows                = strings.Join(goodsMainFieldNames, ",")
	defaultGoodsMainFields       = strings.Join(goodsMainFieldNames, ",")
	goodsMainRowsExpectAutoSet   = strings.Join(stringx.Remove(goodsMainFieldNames, "`delete_at`"), ",")
	goodsMainRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsMainFieldNames, "`id`", "`delete_at`"), "=?,") + "=?"
)

const (
	GoodsMain_Id        dao.TableField = "id"
	GoodsMain_Code      dao.TableField = "code"
	GoodsMain_Name      dao.TableField = "name"
	GoodsMain_Spec      dao.TableField = "spec"
	GoodsMain_Cover     dao.TableField = "cover"
	GoodsMain_SellPrice dao.TableField = "sell_price"
	GoodsMain_StoreQty  dao.TableField = "store_qty"
	GoodsMain_State     dao.TableField = "state"
	GoodsMain_IsSpecial dao.TableField = "is_special"
	GoodsMain_UnitId    dao.TableField = "unit_id"
	GoodsMain_UnitName  dao.TableField = "unit_name"
	GoodsMain_ViewNum   dao.TableField = "view_num"
	GoodsMain_PlatId    dao.TableField = "plat_id"
	GoodsMain_CreateAt  dao.TableField = "create_at"
	GoodsMain_EditAt    dao.TableField = "edit_at"
	GoodsMain_DeleteAt  dao.TableField = "delete_at"
)

type (
	goodsMainModel interface {
		Insert(data *GoodsMain) (effectRow int64, err error)
		TxInsert(tx *sql.Tx, data *GoodsMain) (effectRow int64, err error)
		Update(data map[dao.TableField]any) (effectRow int64, err error)
		TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error)
		Save(data *GoodsMain) (effectRow int64, err error)
		TxSave(tx *sql.Tx, data *GoodsMain) (effectRow int64, err error)
		Delete(ctx context.Context, id string) error
		Field(field string) *defaultGoodsMainModel
		Except(fields ...string) *defaultGoodsMainModel
		Alias(alias string) *defaultGoodsMainModel
		Where(whereStr string, whereData ...any) *defaultGoodsMainModel
		WhereId(id string) *defaultGoodsMainModel
		Order(order string) *defaultGoodsMainModel
		Limit(num int64) *defaultGoodsMainModel
		Plat(id string) *defaultGoodsMainModel
		Find() (*GoodsMain, error)
		FindById(id string) (*GoodsMain, error)
		CacheFind(redis *redisd.Redisd) (*GoodsMain, error)
		CacheFindById(redis *redisd.Redisd, id string) (*GoodsMain, error)
		Page(page int64, rows int64) *defaultGoodsMainModel
		Total() (total int64, err error)
		Select() ([]*GoodsMain, error)
		SelectWithTotal() ([]*GoodsMain, int64, error)
		CacheSelect(redis *redisd.Redisd) ([]*GoodsMain, error)
		Count() (int64, error)
		Inc(field string, num int) (int64, error)
		Dec(field string, num int) (int64, error)
		Ctx(ctx context.Context) *defaultGoodsMainModel
		Reinit() *defaultGoodsMainModel
		Dao() *dao.SqlxDao
	}

	defaultGoodsMainModel struct {
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

	GoodsMain struct {
		Id        string `db:"id"`
		Code      string `db:"code"`       // 商品编号
		Name      string `db:"name"`       // 商品名称
		Spec      string `db:"spec"`       // 商品规格
		Cover     string `db:"cover"`      // 商品封面
		SellPrice int64  `db:"sell_price"` // 商品售价
		StoreQty  int64  `db:"store_qty"`  // 当前库存
		State     int64  `db:"state"`      // 0未上架，1上架
		IsSpecial int64  `db:"is_special"` // 是否活动专用的特殊商品
		UnitId    string `db:"unit_id"`    // 单位
		UnitName  string `db:"unit_name"`  // 单位名称
		ViewNum   int64  `db:"view_num"`   // 浏览数量
		PlatId    string `db:"plat_id"`
		CreateAt  int64  `db:"create_at"`
		EditAt    int64  `db:"edit_at"`
		DeleteAt  int64  `db:"delete_at"`
	}
)

// NewGoodsMainModel returns a model for the database table.
func NewGoodsMainModel(ctxOrNil context.Context, conn sqlx.SqlConn, platId ...string) GoodsMainModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	if ctxOrNil == nil {
		ctxOrNil = context.Background()
	}
	return &customGoodsMainModel{
		defaultGoodsMainModel: newGoodsMainModel(ctxOrNil, conn, platid),
		softDeletable:         softDeletableGoodsMain,
	}
}
func newGoodsMainModel(ctx context.Context, conn sqlx.SqlConn, platId string) *defaultGoodsMainModel {
	dao := dao.NewSqlxDao(conn, "`goods_main`", defaultGoodsMainFields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &defaultGoodsMainModel{
		ctx:             ctx,
		conn:            conn,
		dao:             dao,
		table:           "`goods_main`",
		platId:          platId,
		softDeleteField: "delete_at",
		whereData:       make([]any, 0),
	}
}
func (m *defaultGoodsMainModel) Ctx(ctx context.Context) *defaultGoodsMainModel {
	m.dao.Ctx(ctx)
	return m
}
func (m *defaultGoodsMainModel) WhereId(id string) *defaultGoodsMainModel {
	m.dao.WhereId(id)
	return m
}

func (m *defaultGoodsMainModel) Where(whereStr string, whereData ...any) *defaultGoodsMainModel {
	m.dao.Where(whereStr, whereData...)
	return m
}

func (m *defaultGoodsMainModel) Alias(alias string) *defaultGoodsMainModel {
	m.dao.Alias(alias)
	return m
}
func (m *defaultGoodsMainModel) Field(field string) *defaultGoodsMainModel {
	m.dao.Field(field)
	return m
}
func (m *defaultGoodsMainModel) Except(fields ...string) *defaultGoodsMainModel {
	m.dao.Except(fields...)
	return m
}
func (m *defaultGoodsMainModel) Order(order string) *defaultGoodsMainModel {
	m.dao.Order(order)
	return m
}
func (m *defaultGoodsMainModel) Limit(num int64) *defaultGoodsMainModel {
	m.dao.Limit(num)
	return m
}
func (m *defaultGoodsMainModel) Count() (int64, error) {
	return m.dao.Count()
}
func (m *defaultGoodsMainModel) Inc(field string, num int) (int64, error) {
	return m.dao.Inc(field, num)
}
func (m *defaultGoodsMainModel) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *defaultGoodsMainModel) Dec(field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultGoodsMainModel) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.Dec(field, num)
}
func (m *defaultGoodsMainModel) Plat(id string) *defaultGoodsMainModel {
	m.dao.Plat(id)
	return m
}
func (m *defaultGoodsMainModel) Reinit() *defaultGoodsMainModel {
	m.dao.Reinit()
	return m
}
func (m *defaultGoodsMainModel) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *defaultGoodsMainModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGoodsMainModel) Find() (*GoodsMain, error) {
	resp := &GoodsMain{}
	err := m.dao.Find(resp)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultGoodsMainModel) FindById(id string) (*GoodsMain, error) {
	resp := &GoodsMain{}
	err := m.dao.FindById(resp, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}
func (m *defaultGoodsMainModel) CacheFind(redis *redisd.Redisd) (*GoodsMain, error) {
	resp := &GoodsMain{}
	err := m.dao.CacheFind(redis, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultGoodsMainModel) CacheFindById(redis *redisd.Redisd, id string) (*GoodsMain, error) {
	resp := &GoodsMain{}
	err := m.dao.CacheFindById(redis, resp, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultGoodsMainModel) Total() (total int64, err error) {
	return m.dao.Total()
}
func (m *defaultGoodsMainModel) Select() ([]*GoodsMain, error) {
	resp := make([]*GoodsMain, 0)
	err := m.dao.Select(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *defaultGoodsMainModel) SelectWithTotal() ([]*GoodsMain, int64, error) {
	resp := make([]*GoodsMain, 0)
	var total int64
	err := m.dao.Select(&resp, &total)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (m *defaultGoodsMainModel) CacheSelect(redis *redisd.Redisd) ([]*GoodsMain, error) {
	resp := make([]*GoodsMain, 0)
	err := m.dao.CacheSelect(redis, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultGoodsMainModel) Page(page int64, size int64) *defaultGoodsMainModel {
	m.dao.Page(page, size)
	return m
}

func (m *defaultGoodsMainModel) Insert(data *GoodsMain) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Insert(insertData)
}
func (m *defaultGoodsMainModel) TxInsert(tx *sql.Tx, data *GoodsMain) (effectRow int64, err error) {
	insertData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.TxInsert(tx, insertData)
}

func (m *defaultGoodsMainModel) Update(data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.Update(data)
}
func (m *defaultGoodsMainModel) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64, err error) {
	return m.dao.TxUpdate(tx, data)
}
func (m *defaultGoodsMainModel) Save(data *GoodsMain) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}
func (m *defaultGoodsMainModel) TxSave(tx *sql.Tx, data *GoodsMain) (effectRow int64, err error) {
	saveData, err := dao.PrepareData(data)
	if err != nil {
		return 0, err
	}
	return m.dao.Save(saveData)
}

func (m *defaultGoodsMainModel) tableName() string {
	return m.table
}

// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *defaultGoodsMainModel) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
