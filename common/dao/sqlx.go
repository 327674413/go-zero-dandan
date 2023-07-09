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
	"time"
)

// SqlxDao 自用orm
type SqlxDao struct {
	conn            sqlx.SqlConn
	ctx             context.Context
	table           string
	defaultRowField string
	softDeleteField string
	softDeletable   bool
	fieldSql        string
	whereSql        string
	aliasSql        string
	orderSql        string
	platId          int64
	queryPage       int
	queryRows       int

	whereData []any
	err       error
}

func NewSqlxDao(conn sqlx.SqlConn, tableName string, defaultRowField string, softDeletable bool, softDeleteField string) *SqlxDao {
	return &SqlxDao{
		conn:            conn,
		table:           tableName,
		defaultRowField: defaultRowField,
		softDeletable:   softDeletable,
		softDeleteField: softDeleteField,
	}
}

// Ctx 使用上下文执行sql
func (t *SqlxDao) Ctx(ctx context.Context) *SqlxDao {
	t.ctx = ctx
	return t
}

// Count 查询数量，必须先设置where再使用
func (t *SqlxDao) Count() (int64, error) {
	return 0, nil
}

// Inc 对单个字段进行递增，必须先设置where再使用
func (t *SqlxDao) Inc(field string, num int) (int64, error) {
	return 0, nil
}

// TxInc 事务，对单个字段进行递减，必须先设置where再使用
func (t *SqlxDao) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return 0, nil
}

// Dec 对单个字段进行递减，必须先设置where再使用
func (t *SqlxDao) Dec(field string, num int) (int64, error) {
	return 0, nil
}

// TxDec 使用事务，对单个字段进行递减，必须先设置where再使用
func (t *SqlxDao) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
	return 0, nil
}

func (t *SqlxDao) Find(targetStructPtr any, id ...any) error {
	defer t.Reinit()
	var err error
	if err = t.err; err != nil {
		t.err = nil
		return err
	}
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
		if t.ctx != nil {
			err = t.conn.QueryRowPartialCtx(t.ctx, targetStructPtr, sql) //若用QueryRowCtx 必须字段都覆盖
		} else {
			err = t.conn.QueryRowPartial(targetStructPtr, sql)
		}

	} else {
		if t.whereSql == "" {
			t.whereSql = "1=1"
		}
		if t.platId != 0 {
			t.whereSql = t.whereSql + " AND plat_id=" + fmt.Sprintf("%d", t.platId)
		}
		sql = fmt.Sprintf("select %s from %s %s where "+t.whereSql+" limit 1", field, t.table, t.aliasSql)
		if t.ctx != nil {
			err = t.conn.QueryRowPartialCtx(t.ctx, targetStructPtr, sql, t.whereData...)
		} else {
			err = t.conn.QueryRowPartial(targetStructPtr, sql, t.whereData...)
		}

	}
	switch err {
	case nil:
		return nil
	case sqlx.ErrNotFound:
		return nil
	default:
		return err
	}
}

// Select 查询所有数据,需传入目标结构体切片的指针
func (t *SqlxDao) Select(targetStructPtr any) error {
	defer t.Reinit()
	var err error
	if err = t.err; err != nil {
		t.err = nil
		return err
	}
	var sql string
	field := t.defaultRowField
	if t.fieldSql != "" {
		field = t.fieldSql
	}
	if t.whereSql == "" {
		t.whereSql = "1=1"
	}
	sql = fmt.Sprintf("select %s from %s %s where "+t.whereSql, field, t.table, t.aliasSql)
	if t.ctx != nil {
		err = t.conn.QueryRowsPartialCtx(t.ctx, targetStructPtr, sql, t.whereData...)
	} else {
		err = t.conn.QueryRowsPartial(targetStructPtr, sql, t.whereData...)
	}
	switch err {
	case nil:
		return nil
	case sqlx.ErrNotFound:
		return nil
	default:
		return err
	}
}

// CacheFind 优先从缓存里查询数据，若缓存不存在则从数据库里查询，无失效时间
func (t *SqlxDao) CacheFind(redis *redisd.Redisd, targetStructPtr any, id ...int64) error {
	defer t.Reinit()
	cacheField := "model_" + t.table
	cacheKey := strconv.FormatInt(id[0], 10)
	// todo::需要把where条件一起放进去作为key，这样就能支持更多的自动缓存查询
	err := redis.GetData(cacheField, cacheKey, targetStructPtr)
	if err == nil {
		return nil
	}
	err = t.Find(&targetStructPtr, id[0])
	if err != nil {
		return err
	}
	// todo::如果没查到，是不是会有问题
	redis.SetData(cacheField, cacheKey, targetStructPtr)
	return nil
}
func (t *SqlxDao) Insert(data map[string]string) (int64, error) {
	var sqlRes sql.Result
	var err error
	query, insertData, err := t.prepareInsert(data)
	if err != nil {
		return 0, err
	}

	if t.ctx != nil {
		sqlRes, err = t.conn.ExecCtx(t.ctx, query, insertData...)
	} else {
		sqlRes, err = t.conn.Exec(query, insertData...)
	}

	if err != nil {
		return 0, err
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}
func (t *SqlxDao) TxInsert(tx *sql.Tx, data map[string]string) (int64, error) {
	var sqlRes sql.Result
	var err error
	query, insertData, err := t.prepareInsert(data)
	if err != nil {
		return 0, err
	}
	if t.ctx != nil {
		sqlRes, err = tx.ExecContext(t.ctx, query, insertData...)
	} else {
		sqlRes, err = tx.Exec(query, insertData...)
	}

	if err != nil {
		return 0, err
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}
func (t *SqlxDao) prepareInsert(data map[string]string) (string, []any, error) {
	insertFields := ""
	insertValues := ""
	insertData := make([]any, 0)
	for k, v := range data {
		insertFields = insertFields + k + ","
		insertValues = insertValues + "?,"
		insertData = append(insertData, v)
	}
	if len(insertFields) > 0 {
		insertFields = insertFields[:len(insertFields)-1]
		insertValues = insertValues[:len(insertValues)-1]
	} else {
		return "", nil, errors.New("insert data is empty")
	}
	query := fmt.Sprintf("insert into %s (%s,create_at,update_at,plat_id) values (%s,?,?,?)", t.table, insertFields, insertValues)
	nowStamp := fmt.Sprintf("%d", time.Now().Unix())
	insertData = append(insertData, nowStamp)
	insertData = append(insertData, nowStamp)
	insertData = append(insertData, fmt.Sprintf("%d", t.platId))
	return query, insertData, nil
}

// Delete 优先根据传入的id删除，若未传则where条件必有
func (t *SqlxDao) Delete(id ...int64) (int64, error) {
	defer t.Reinit()
	query := fmt.Sprintf("delete from %s where `id` = ?", t.table)
	var sqlRes sql.Result
	var err error
	if t.ctx != nil {
		sqlRes, err = t.conn.ExecCtx(t.ctx, query, id)
	} else {
		sqlRes, err = t.conn.Exec(query, id)
	}

	if err != nil {
		return 0, err
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// TxDelete 事务Delete
func (t *SqlxDao) TxDelete(tx *sql.Tx, id ...int64) (int64, error) {
	defer t.Reinit()
	query := fmt.Sprintf("delete from %s where `id` = ?", t.table)
	var sqlRes sql.Result
	var err error
	if t.ctx != nil {
		sqlRes, err = tx.ExecContext(t.ctx, query, id)
	} else {
		sqlRes, err = tx.Exec(query, id)
	}

	if err != nil {
		return 0, err
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// Page 设置当前查询第几页，查多少行
func (t *SqlxDao) Page(page int, rows int) {
	t.queryPage = page
	t.queryRows = rows
}

// Update 必须先设置where或在data中携带id，data中的id优先级高，若带id只能修改单个
func (t *SqlxDao) Update(data map[string]string) (int64, error) {
	defer t.Reinit()
	query, params, err := t.prepareUpdate(data)
	if err != nil {
		return 0, err
	}

	var sqlRes sql.Result
	if t.ctx != nil {
		sqlRes, err = t.conn.ExecCtx(t.ctx, query, params...)
	} else {
		sqlRes, err = t.conn.Exec(query, params...)
	}

	if err != nil {
		return 0, err
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// TxUpdate 同Update，事务专用
func (t *SqlxDao) TxUpdate(tx *sql.Tx, data map[string]string) (int64, error) {
	defer t.Reinit()
	query, params, err := t.prepareUpdate(data)
	if err != nil {
		return 0, err
	}
	var sqlRes sql.Result
	if t.ctx != nil {
		sqlRes, err = tx.ExecContext(t.ctx, query, params...)
	} else {
		sqlRes, err = tx.Exec(query, params...)
	}

	if err != nil {
		return 0, err
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil

}

// TxSave 事务用Save
func (t *SqlxDao) TxSave(tx *sql.Tx, data map[string]string) (int64, error) {
	var updateStr string
	params := make([]any, 0)
	var id int64
	var hasId bool
	for k, v := range data {
		if k == "id" {
			id = utild.AnyToInt64(k)
			hasId = true
			continue
		}
		updateStr = updateStr + fmt.Sprintf("%s=?,", utild.StrToSnake(k))
		params = append(params, v)
	}
	if hasId == false {
		return 0, errors.New("save data must need id")
	}
	var currData map[string]string
	err := t.Find(&currData, id)
	if err != nil {
		return 0, err
	}
	if len(currData) == 0 {
		_, err = t.TxInsert(tx, data)
		if err != nil {
			return 0, err
		}
		return 0, nil
	}
	if len(updateStr) > 0 {
		updateStr = updateStr[:len(updateStr)-1]
	} else {
		return 0, errors.New("update data is empty")
	}
	updateStr = updateStr + fmt.Sprintf(",update_at=%d", utild.GetStamp())
	whereStr := t.whereSql
	if whereStr == "" {
		if id == 0 {
			return 0, errors.New("update data must need where")
		} else {
			whereStr = fmt.Sprintf("id=%d", id)
		}

	}
	if t.platId != 0 {
		whereStr = whereStr + fmt.Sprintf(" AND plat_id=%d", t.platId)
	}
	query := fmt.Sprintf("update %s set %s where %s", t.table, updateStr, whereStr)
	var sqlRes sql.Result
	if t.ctx != nil {
		sqlRes, err = tx.ExecContext(t.ctx, query, params...)
	} else {
		sqlRes, err = tx.Exec(query, params...)
	}

	if err != nil {
		return 0, err
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// Save 如果数据存在则修改，不存在则新增，data中必须有id
func (t *SqlxDao) Save(data map[string]string) (int64, error) {
	var updateStr string
	params := make([]any, 0)
	var id int64
	var hasId bool
	for k, v := range data {
		if k == "id" {
			id = utild.AnyToInt64(k)
			hasId = true
			continue
		}
		updateStr = updateStr + fmt.Sprintf("%s=?,", utild.StrToSnake(k))
		params = append(params, v)
	}
	if hasId == false {
		return 0, errors.New("save data must need id")
	}
	var currData map[string]string
	err := t.Find(&currData, id)
	if err != nil {
		return 0, err
	}
	if len(currData) == 0 {
		_, err = t.Insert(data)
		if err != nil {
			return 0, err
		}
		return 0, nil
	}
	if len(updateStr) > 0 {
		updateStr = updateStr[:len(updateStr)-1]
	} else {
		return 0, errors.New("update data is empty")
	}
	updateStr = updateStr + fmt.Sprintf(",update_at=%d", utild.GetStamp())
	whereStr := t.whereSql
	if whereStr == "" {
		if id == 0 {
			return 0, errors.New("update data must need where")
		} else {
			whereStr = fmt.Sprintf("id=%d", id)
		}

	}
	if t.platId != 0 {
		whereStr = whereStr + fmt.Sprintf(" AND plat_id=%d", t.platId)
	}
	query := fmt.Sprintf("update %s set %s where %s", t.table, updateStr, whereStr)
	var sqlRes sql.Result
	if t.ctx != nil {
		sqlRes, err = t.conn.ExecCtx(t.ctx, query, params...)
	} else {
		sqlRes, err = t.conn.Exec(query, params...)
	}

	if err != nil {
		return 0, err
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// prepareUpdate 内部封装，让事务和非事务公共代码复用
func (t *SqlxDao) prepareUpdate(data map[string]string) (string, []any, error) {
	//构造修改内容部分的sql
	var updateStr string
	params := make([]any, 0)
	var id int64
	for k, v := range data {
		if k == "id" {
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
func (t *SqlxDao) WhereId(id int64) *SqlxDao {
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
	t.platId = platId
	return t
}

// Reinit 每次执行完数据库操作后恢复初始化，保证不干扰下次使用
func (t *SqlxDao) Reinit() {
	t.whereSql = ""
	t.fieldSql = ""
	t.aliasSql = ""
	t.orderSql = ""
	t.whereData = make([]any, 0)
	t.queryRows = 0
	t.queryPage = 0
	t.err = nil
	t.ctx = nil
}
