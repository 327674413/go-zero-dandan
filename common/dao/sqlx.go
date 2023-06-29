package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/common/redisd"
	"go-zero-dandan/common/utild"
	"strconv"
)

// SqlxDao 自用orm
type SqlxDao struct {
	table           string
	defaultRowField string
	softDeleteField string
	softDeletable   bool
	fieldSql        string
	whereSql        string
	aliasSql        string
	orderSql        string
	platId          int64

	whereData []any
	err       error
}

func NewSqlxDao(tableName string, defaultRowField string, softDeletable bool, softDeleteField string) *SqlxDao {
	return &SqlxDao{
		table:           tableName,
		defaultRowField: defaultRowField,
		softDeletable:   softDeletable,
		softDeleteField: softDeleteField,
	}
}

// Count 查询数量，必须先设置where再使用
func (t *SqlxDao) Count(conn sqlx.SqlConn, ctx context.Context) (int, error) {
	return 0, nil
}

// Inc 对单个字段进行递增，必须先设置where再使用
func (t *SqlxDao) Inc(ctx context.Context, field string, num int) error {

	return nil
}

// TxInc 事务用的Inc
func (t *SqlxDao) TxInc(tx *sql.Tx, ctx context.Context, field string, num int) error {
	return nil
}

// Dec 对单个字段进行递减，必须先设置where再使用
func (t *SqlxDao) Dec(ctx context.Context, field string, num int) error {

	return nil
}

// TxDec 事务用的Dec
func (t *SqlxDao) TxDec(tx *sql.Tx, ctx context.Context, field string, num int) error {
	return nil
}

func (t *SqlxDao) Find(conn sqlx.SqlConn, ctx context.Context, id ...any) (map[string]string, error) {
	defer t.Reinit()
	var err error
	if err = t.err; err != nil {
		t.err = nil
		return nil, err
	}
	var resp map[string]string
	var sql string
	field := t.defaultRowField
	if t.fieldSql != "" {
		field = t.fieldSql
	}
	if len(id) > 0 {
		if t.whereSql == "" {
			t.whereSql = "1=1"
		}
		if t.platId != 0 {
			t.whereSql = t.whereSql + fmt.Sprintf(" AND id=%d AND plat_id=%d", id[0], t.platId)
		} else {
			t.whereSql = t.whereSql + fmt.Sprintf(" AND id=%d", id[0])
		}
		sql = fmt.Sprintf("select %s from %s where %s limit 1", field, t.table, t.whereSql)
		err = conn.QueryRowPartialCtx(ctx, &resp, sql) //QueryRowCtx 必须字段都覆盖
	} else {
		if t.whereSql == "" {
			t.whereSql = "1=1"
		}
		if t.platId != 0 {
			t.whereSql = t.whereSql + " AND plat_id=" + fmt.Sprintf("%d", t.platId)
		}
		sql = fmt.Sprintf("select %s from %s %s where "+t.whereSql+" limit 1", field, t.table, t.aliasSql)
		err = conn.QueryRowPartialCtx(ctx, &resp, sql, t.whereData...)
	}
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return resp, nil
	default:
		return nil, err
	}
}

// CacheFind 优先从缓存里查询数据，若缓存不存在则从数据库里查询，无失效时间
func (t *SqlxDao) CacheFind(conn sqlx.SqlConn, ctx context.Context, redis *redisd.Redisd, id ...int64) (map[string]string, error) {
	defer t.Reinit()
	var resp map[string]string
	cacheField := "model_" + t.table
	cacheKey := strconv.FormatInt(id[0], 10)
	// todo::需要把where条件一起放进去作为key，这样就能支持更多的自动缓存查询
	err := redis.GetData(cacheField, cacheKey, resp)
	if err == nil {
		return resp, nil
	}
	resp, err = t.Find(conn, ctx, id[0])
	if err != nil {
		return resp, err
	}
	_, ok := resp["Id"]
	if ok {
		redis.SetData(cacheField, cacheKey, resp)
	}
	return resp, nil
}

// Delete 优先根据传入的id删除，若未传则where条件必有
func (t *SqlxDao) Delete(conn sqlx.SqlConn, ctx context.Context, id ...int64) error {
	defer t.Reinit()
	query := fmt.Sprintf("delete from %s where `id` = ?", t.table)
	_, err := conn.ExecCtx(ctx, query, id)
	return err
}

// TxDelete 事务Delete
func (t *SqlxDao) TxDelete(conn sqlx.SqlConn, ctx context.Context, id ...int64) error {
	defer t.Reinit()
	query := fmt.Sprintf("delete from %s where `id` = ?", t.table)
	_, err := conn.ExecCtx(ctx, query, id)
	return err
}

// List 查询所有数据
func (t *SqlxDao) List(conn sqlx.SqlConn, ctx context.Context) ([]map[string]string, error) {
	defer t.Reinit()
	var err error
	if err = t.err; err != nil {
		t.err = nil
		return nil, err
	}
	var resp []map[string]string
	var sql string
	field := t.defaultRowField
	if t.fieldSql != "" {
		field = t.fieldSql
	}
	if t.whereSql == "" {
		t.whereSql = "1=1"
	}
	sql = fmt.Sprintf("select %s from %s %s where "+t.whereSql, field, t.table, t.aliasSql)
	err = conn.QueryRowsPartialCtx(ctx, &resp, sql, t.whereData...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return resp, nil
	default:
		return nil, err
	}
}
func (t *SqlxDao) Page(ctx context.Context, page int, rows int) ([]map[string]string, error) {
	defer t.Reinit()
	return nil, nil
}

// Update 必须先设置where或在data中携带id，data中的id优先级高，若带id只能修改单个
func (t *SqlxDao) Update(conn sqlx.SqlConn, ctx context.Context, data map[string]string) error {
	defer t.Reinit()
	query, params, err := t.prepareUpdate(data)
	if err != nil {
		return err
	}
	_, err = conn.ExecCtx(ctx, query, params...)
	return err
}

// TxUpdate 同Update，事务专用
func (t *SqlxDao) TxUpdate(tx *sql.Tx, ctx context.Context, data map[string]string) error {
	defer t.Reinit()
	query, params, err := t.prepareUpdate(data)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, query, params...)
	return err
}

// prepareUpdate 内部封装，让事务和非事务公共代码复用
func (t *SqlxDao) prepareUpdate(data map[string]string) (string, []any, error) {
	//构造修改内容部分的sql
	var updateStr string
	params := make([]any, 0)
	var id int64
	for k, v := range data {
		if k == "Id" {
			id = utild.AnyToInt64(v)
			continue
		}
		updateStr = updateStr + fmt.Sprintf("%s=?,", utild.StrToSnake(k))
		params = append(params, v)
	}
	if len(updateStr) > 0 {
		updateStr = updateStr[:len(updateStr)-1]
	} else {
		return "", nil, errors.New("update data is empty")
	}
	//自动添加修改时间字段
	updateStr = updateStr + fmt.Sprintf(",update_at=%d", utild.GetStamp())
	//构造where部分sql
	whereStr := t.whereSql
	if id != 0 {
		//若data带id，则必为修改该id数据
		whereStr = fmt.Sprintf("id=%d", id)
	} else if whereStr == "" {
		//若data未带id，则必须给条件，即便全修改也要给条件1=1
		return "", nil, errors.New("update param where is empty")
	}
	// 多应用时自动增加多应用条件
	if t.platId != 0 {
		whereStr = whereStr + fmt.Sprintf(" AND plat_id=%d", t.platId)
	}
	query := fmt.Sprintf("update %s set %s where %s", t.table, updateStr, whereStr)
	return query, params, nil
}

// WhereId 根据id设置where，优先级最高，执行后将会覆盖原有条件
func (t *SqlxDao) WhereId(id int) *SqlxDao {
	t.whereSql = "id=?"
	t.whereData = append(t.whereData, id)
	return t
}

// WhereStr 在原有where条件上拼接一个 AND (条件)，注意sql注入，若有不可靠参数请使用whereRaw,如果是or也是自行用whereRaw拼接
func (t *SqlxDao) WhereStr(whereStr string) *SqlxDao {
	if t.whereSql != "" {
		t.whereSql += "AND (" + whereStr + ")"
	} else {
		t.whereSql = "(" + whereStr + ")"
	}
	return t
}

// WhereRaw 通过参数定义where条件，可防sql注入
func (t *SqlxDao) WhereRaw(whereStr string, whereData []any) *SqlxDao {
	if t.whereSql != "" {
		t.whereSql += "AND (" + whereStr + ")"
	} else {
		t.whereSql = "(" + whereStr + ")"
	}
	t.whereData = append(t.whereData, whereData...)
	return t
}

// Alias 设置主表别名，当联表查询时，必须通过Field指定字段
func (t *SqlxDao) Alias(field string) *SqlxDao {
	t.aliasSql = field
	return t
}

// Field 设置查询字段，若不限制则全字段查询
func (t *SqlxDao) Field(field string) *SqlxDao {
	t.fieldSql = field
	return t
}

// Order 设置排序字段
func (t *SqlxDao) Order(order string) *SqlxDao {

	return t
}

// Plat 设置应用id
func (t *SqlxDao) Plat(platId int64) *SqlxDao {
	return t
}

// Reinit 每次执行完数据库操作后恢复初始化，保证不干扰下次使用
func (t *SqlxDao) Reinit() {
	t.whereSql = ""
	t.fieldSql = ""
	t.aliasSql = ""
	t.orderSql = ""
	t.whereData = make([]any, 0)
	t.err = nil
}
