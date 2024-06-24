package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/common/redisd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"strings"
	"time"
)

// todo::错误返回全部封装， 用resd来包装，这样错误吗就不用每次写了

// SqlxDao 自用orm
type SqlxDao struct {
	conn             sqlx.SqlConn
	ctx              context.Context
	table            string
	defaultRowField  string
	exceptRowFields  []string
	defaultQuerySize int64
	softDeleteField  string
	softDeletable    bool

	tableAlias         string
	orderSql           string
	platId             string
	queryPage          int64
	querySize          int64
	limitNum           int64
	joinTables         []string
	whereData          []any
	fieldSql           string
	whereSql           string
	lastQuerySql       string
	lastQueryPage      int64
	lastQuerySize      int64
	lastQueryIsCache   bool
	err                error
	emptyQueryCacheSec int //开启缓存，未查到数据时，空数据的缓存秒数
	defaultCacheSec    int
}

func NewSqlxDao(conn sqlx.SqlConn, tableName string, defaultRowField string, softDeletable bool, softDeleteField string) *SqlxDao {
	return &SqlxDao{
		conn:               conn,
		table:              tableName,
		defaultRowField:    defaultRowField,
		softDeletable:      softDeletable,
		softDeleteField:    softDeleteField,
		whereData:          make([]any, 0),
		joinTables:         make([]string, 0),
		defaultQuerySize:   10,
		emptyQueryCacheSec: 0,
		defaultCacheSec:    0, //默认0，永久不失效
	}
}

// StartTrans 开启事务
func StartTrans(conn sqlx.SqlConn, ctx ...context.Context) (*sql.Tx, error) {
	var sqlCtx context.Context
	if len(ctx) > 0 {
		sqlCtx = ctx[0]
	} else {
		sqlCtx = context.Background()
	}
	db, err := conn.RawDB()
	if err != nil {
		return nil, resd.Error(err)
	}
	tx, err := db.BeginTx(sqlCtx, nil)
	if err != nil {
		return nil, resd.Error(err)
	}
	return tx, nil
}

// Commit 提交事务
func Commit(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return resd.Error(err)
	}
	return nil
}

// Rollback 回滚事务
func Rollback(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return resd.Error(err)
	}
	return nil
}

// Ctx 使用上下文执行sql
func (t *SqlxDao) Ctx(ctx context.Context) *SqlxDao {
	t.ctx = ctx
	return t
}

// Except 左联表
func (t *SqlxDao) Except(fields ...string) {
	t.exceptRowFields = fields
}

// Limit Select时限制数量，如果设置了Page则会被覆盖
func (t *SqlxDao) Limit(num int64) *SqlxDao {
	return t.Page(1, num)
}

// LeftJoin 左联表
func (t *SqlxDao) LeftJoin(joinStr string) {
	t.joinTables = append(t.joinTables, joinStr)
}

// RightJoin 左联表
func (t *SqlxDao) RightJoin(joinStr string) {
	t.joinTables = append(t.joinTables, joinStr)
}

// InnerJoin 左联表
func (t *SqlxDao) InnerJoin(joinStr string) {
	t.joinTables = append(t.joinTables, joinStr)
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
func (t *SqlxDao) FindById(targetStructPtr any, id string) error {
	t.whereSql = ""
	t.WhereId(id)
	return t.Find(targetStructPtr)
}
func (t *SqlxDao) getFieldParam() string {
	field := t.defaultRowField
	if t.fieldSql != "" {
		field = t.fieldSql
	}
	// 移除要除去的字段
	if len(t.exceptRowFields) > 0 {
		fieldMap := make(map[string]bool)
		for _, v := range t.exceptRowFields {
			fieldMap[v] = true
		}
		fields := strings.Split(field, ",")
		resField := ""
		for _, v := range fields {
			if _, ok := fieldMap[v]; !ok {
				resField = v + ","
			}
		}
		if len(resField) > 0 {
			field = resField[:len(resField)-1]
		}
	}
	return field
}
func (t *SqlxDao) getTableParam() string {
	table := t.table
	if t.tableAlias != "" {
		table = table + " " + t.tableAlias
	}
	if len(t.joinTables) > 0 {
		for _, v := range t.joinTables {
			table = table + " LEFT JOIN " + v
		}
	}
	return table
}

func (t *SqlxDao) getWhereParam() string {
	where := t.whereSql
	if where == "" {
		where = "1=1"
	}
	plat := "plat_id"
	if t.tableAlias != "" {
		plat = t.tableAlias + "." + plat
	}
	if t.platId != "" {
		where = where + "" + fmt.Sprintf(" AND %s=%s", plat, t.platId)
	}
	if t.softDeletable {
		aliasDelete := ""
		if t.tableAlias != "" {
			aliasDelete = t.tableAlias + "."
		}
		where = where + "" + fmt.Sprintf(" AND %s%s=0", aliasDelete, t.softDeleteField)
	}
	return where
}
func (t *SqlxDao) getPageParam() string {
	if t.querySize == 0 {
		return ""
		t.querySize = t.defaultQuerySize
	}
	if t.queryPage <= 0 {
		t.queryPage = 1
	}
	offset := (t.queryPage - 1) * t.querySize
	return fmt.Sprintf("LIMIT %d, %d", offset, t.querySize)
}
func (t *SqlxDao) validate() (err error) {
	if err = t.err; err != nil {
		t.err = nil
		return err
	}
	if len(t.joinTables) > 0 && t.fieldSql == "" {
		return resd.NewErr("联表查询时必须设置Field")
	}
	return nil
}
func (t *SqlxDao) Find(targetStructPtr any) error {
	defer t.Reinit()
	err := t.validate()
	if err != nil {
		return err
	}
	fieldParam := t.getFieldParam()
	tableParam := t.getTableParam()
	whereParam := t.getWhereParam()
	if t.orderSql != "" {
		t.orderSql = " ORDER BY " + t.orderSql
	}
	checkParamsNum := strings.Count(whereParam, "?")
	if checkParamsNum > 0 && checkParamsNum != len(t.whereData) {
		return resd.NewErr(fmt.Sprintf("查询参数不匹配，预置%d，实际%d", checkParamsNum, len(t.whereData)))
	}
	sql := fmt.Sprintf("select %s from %s where "+whereParam+t.orderSql+" limit 1", fieldParam, tableParam)
	if t.ctx != nil {
		fmt.Println(sql, t.whereData)
		err = t.conn.QueryRowPartialCtx(t.ctx, targetStructPtr, sql, t.whereData...)
	} else {
		err = t.conn.QueryRowPartial(targetStructPtr, sql, t.whereData...)
	}
	if err != nil {
		if err == sqlx.ErrNotFound {
			return err
		}
		return resd.Error(err)
	} else {
		return nil
	}
}

// Select 查询所有数据,需传入目标结构体切片的指针
func (t *SqlxDao) Select(targetStructPtr any, totalPtr ...any) error {
	defer t.Reinit()
	err := t.validate()
	if err != nil {
		return err
	}
	fieldParam := t.getFieldParam()
	tableParam := t.getTableParam()
	whereParam := t.getWhereParam()
	pageParam := t.getPageParam()
	orderSql := ""
	if t.orderSql != "" {
		orderSql = " ORDER BY " + t.orderSql
	}
	checkParamsNum := strings.Count(whereParam, "?")
	if checkParamsNum > 0 && checkParamsNum != len(t.whereData) {
		return resd.NewErr(fmt.Sprintf("查询参数不匹配，预置%d，实际%d", checkParamsNum, len(t.whereData)))
	}
	sql := fmt.Sprintf("select %s from %s where "+whereParam+orderSql+" "+pageParam, fieldParam, tableParam)
	if t.ctx != nil {
		err = t.conn.QueryRowsPartialCtx(t.ctx, targetStructPtr, sql, t.whereData...)
	} else {
		err = t.conn.QueryRowsPartial(targetStructPtr, sql, t.whereData...)
	}
	// select传入的应该是切片指针，似乎往切片写入数据时，没查到也不会进err
	if err != nil {
		logc.Error(t.ctx, fmt.Sprintf("【error sql】 : %s , 【data】 : %v", sql, t.whereData))
		return resd.Error(err)
	}
	t.lastQuerySql = sql //Todo::这里不是真实执行的sql，不知道怎么获取
	t.lastQueryPage = t.queryPage
	t.lastQuerySize = t.querySize
	t.lastQueryIsCache = false
	if len(totalPtr) > 0 {
		sql = fmt.Sprintf("select COUNT(*) from %s where "+whereParam, tableParam)
		if t.ctx != nil {
			err = t.conn.QueryRowPartialCtx(t.ctx, totalPtr[0], sql, t.whereData...)
		} else {
			err = t.conn.QueryRowPartial(totalPtr[0], sql, t.whereData...)
		}
	}
	return nil
}
func (t *SqlxDao) getSelectCacheKey() (cacheField string, cacheKey string) {
	cacheField = t.getCachePrefixField()
	if t.queryPage > 0 {
		page := t.getPageParam()
		cacheKey = fmt.Sprintf("%s_%s_%s_%s", t.whereSql, fmt.Sprint(t.whereData), t.orderSql, page)
	} else {
		cacheKey = fmt.Sprintf("%s_%s_%s", t.whereSql, fmt.Sprint(t.whereData), t.orderSql)
	}

	return cacheField, cacheKey
}

// CacheSelect 缓存查数据
func (t *SqlxDao) CacheSelect(redis *redisd.Redisd, targetStructPtr any) error {
	defer t.Reinit()
	if targetStructPtr == nil {
		return resd.NewErrCtx(t.ctx, "赋值对象未初始化，为nil")
	}
	cacheField, cacheKey := t.getSelectCacheKey()
	err := redis.GetData(cacheField, cacheKey, targetStructPtr)
	if err == nil {
		t.lastQueryPage = t.queryPage
		t.lastQuerySize = t.querySize
		t.lastQueryIsCache = true
		return nil
	}
	err = t.Select(targetStructPtr)
	if err != nil {
		return resd.Error(err)
	}
	if fmt.Sprint(targetStructPtr) != "&[]" {
		//有查到数据，按默认缓存走
		redis.SetDataEx(cacheField, cacheKey, targetStructPtr, t.defaultCacheSec)
	} else if t.emptyQueryCacheSec > 0 {
		//没查到数据，且需要设置空数据缓存，避免持续打到数据库中
		redis.SetDataEx(cacheField, cacheKey, targetStructPtr, t.emptyQueryCacheSec)
	}
	return nil
}

// CacheFind 优先从缓存里查询数据，若缓存不存在则从数据库里查询，无失效时间
func (t *SqlxDao) CacheFind(redis *redisd.Redisd, targetStructPtr any) error {
	defer t.Reinit()
	if targetStructPtr == nil {
		return resd.NewErrCtx(t.ctx, "赋值对象未初始化，为nil")
	}
	cacheField, cacheKey := t.getSelectCacheKey()
	err := redis.GetData(cacheField, cacheKey, targetStructPtr)
	if err == nil {
		t.lastQueryIsCache = true
		return nil
	}
	err = t.Find(targetStructPtr)
	if err != nil && err != sqlx.ErrNotFound {
		return resd.Error(err)
	}
	if err == nil {
		//有查到数据，按默认缓存走
		redis.SetDataEx(cacheField, cacheKey, targetStructPtr, t.defaultCacheSec)
	} else if t.emptyQueryCacheSec > 0 {
		//没查到数据，且需要设置空数据缓存，避免持续打到数据库中
		redis.SetDataEx(cacheField, cacheKey, targetStructPtr, t.emptyQueryCacheSec)
	}
	return nil
}

// CacheFindById 优先从缓存里查询数据，若缓存不存在则从数据库里查询，无失效时间
func (t *SqlxDao) CacheFindById(redis *redisd.Redisd, targetStructPtr any, id string) (err error) {
	defer t.Reinit()
	cacheField := t.getCachePrefixField()
	if t.ctx == nil {
		err = redis.GetData(cacheField, id, targetStructPtr)
	} else {
		err = redis.GetDataCtx(t.ctx, cacheField, id, targetStructPtr)
	}

	if err == nil {
		t.lastQueryIsCache = true
		return nil
	}
	err = t.FindById(targetStructPtr, id)
	if err != nil {
		return err
	}
	// todo::如果没查到，是不是会有问题
	if t.ctx == nil {
		redis.SetData(cacheField, id, targetStructPtr)
	} else {

		redis.SetDataCtx(t.ctx, cacheField, id, targetStructPtr)
	}

	return nil
}

// ClearCache 清空自带缓存
func (t *SqlxDao) ClearCache(redis *redisd.Redisd) (delNum int64, err error) {
	result, err := redis.DelKeyByPrefix(t.getCachePrefixField() + "*")
	if err != nil {
		return 0, resd.ErrorCtx(t.ctx, err)
	}
	num, ok := result.(int64)
	if ok {
		delNum = num
	}
	return delNum, nil
}
func (t *SqlxDao) Insert(data map[string]any) (int64, error) {
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
		return 0, resd.Error(err)
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}
func (t *SqlxDao) TxInsert(tx *sql.Tx, data map[string]any) (int64, error) {
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
		return 0, resd.Error(err)
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}
func (t *SqlxDao) prepareInsert(data map[string]any) (string, []any, error) {
	insertFields := ""
	insertValues := ""
	insertData := make([]any, 0)
	hasPlatId := false
	for k, v := range data {
		//自动填充字段过滤
		if k == "create_at" || k == "update_at" {
			continue
		}
		if k == "plat_id" {
			//如果存在plat_id且非零值，则按目标值来
			if v != "0" && v != "" {
				hasPlatId = true
			} else {
				continue
			}
		}
		insertFields = insertFields + k + ","
		insertValues = insertValues + "?,"
		insertData = append(insertData, v)
	}
	if len(insertFields) > 0 {
		insertFields = insertFields[:len(insertFields)-1]
		insertValues = insertValues[:len(insertValues)-1]
	} else {
		return "", nil, resd.NewErr("insert data is empty", 4) //这里用了第4层能定位到业务调用代码处，暂不确定是否靠谱
	}
	if !hasPlatId && t.platId != "" {
		insertFields = insertFields + ",plat_id"
		insertValues = insertValues + ",?"
		insertData = append(insertData, t.platId)
	}
	query := fmt.Sprintf("insert into %s (%s,create_at,update_at) values (%s,?,?)", t.table, insertFields, insertValues)
	nowStamp := fmt.Sprintf("%d", time.Now().Unix())
	insertData = append(insertData, nowStamp)
	insertData = append(insertData, nowStamp)
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
		return 0, resd.Error(err)
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
		return 0, resd.Error(err)
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// Page 设置当前查询第几页，查多少行
func (t *SqlxDao) Page(page int64, size int64) *SqlxDao {
	t.queryPage = page
	t.querySize = size
	return t
}

// Update 必须先设置where或在data中携带id，data中的id优先级高，若带id只能修改单个
func (t *SqlxDao) Update(data map[string]any) (int64, error) {
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
		return 0, resd.Error(err)
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// TxUpdate 同Update，事务专用
func (t *SqlxDao) TxUpdate(tx *sql.Tx, data map[string]any) (int64, error) {
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
		return 0, resd.Error(err)
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil

}

// TxSave 事务用Save
func (t *SqlxDao) TxSave(tx *sql.Tx, data map[string]any) (int64, error) {
	var updateStr string
	params := make([]any, 0)
	var id string
	var hasId bool
	for k, v := range data {
		if k == "id" {
			id = k
			hasId = true
			continue
		}
		updateStr = updateStr + fmt.Sprintf("%s=?,", utild.StrToSnake(k))
		params = append(params, v)
	}
	if hasId == false {
		return 0, resd.NewErr("save data must need id")
	}
	var currData map[string]string
	err := t.FindById(&currData, id)
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
		return 0, resd.NewErr("update data is empty")
	}
	updateStr = updateStr + fmt.Sprintf(",update_at=%d", utild.GetStamp())
	whereStr := t.whereSql
	if whereStr == "" {
		if id == "" {
			return 0, resd.NewErr("update data must need where")
		} else {
			whereStr = fmt.Sprintf("id=%d", id)
		}

	}
	if t.platId != "" {
		whereStr = whereStr + fmt.Sprintf(" AND plat_id=%s", t.platId)
	}
	query := fmt.Sprintf("update %s set %s where %s", t.table, updateStr, whereStr)
	var sqlRes sql.Result
	if t.ctx != nil {
		sqlRes, err = tx.ExecContext(t.ctx, query, params...)
	} else {
		sqlRes, err = tx.Exec(query, params...)
	}

	if err != nil {
		return 0, resd.Error(err)
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// Save 如果数据存在则修改，不存在则新增，data中必须有id
func (t *SqlxDao) Save(data map[string]any) (int64, error) {
	var updateStr string
	params := make([]any, 0)
	var id string
	var hasId bool
	for k, v := range data {
		if k == "id" {
			id = k
			hasId = true
			continue
		}
		updateStr = updateStr + fmt.Sprintf("%s=?,", utild.StrToSnake(k))
		params = append(params, v)
	}
	if hasId == false {
		return 0, resd.NewErr("save data must need id")
	}
	var currData map[string]string
	err := t.FindById(&currData, id)
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
		return 0, resd.NewErr("update data is empty")
	}
	updateStr = updateStr + fmt.Sprintf(",update_at=%d", utild.GetStamp())
	whereStr := t.whereSql
	if whereStr == "" {
		if id == "" {
			return 0, resd.NewErr("update data must need where")
		} else {
			whereStr = fmt.Sprintf("id=%d", id)
		}

	}
	if t.platId != "" {
		whereStr = whereStr + fmt.Sprintf(" AND plat_id=%s", t.platId)
	}
	query := fmt.Sprintf("update %s set %s where %s", t.table, updateStr, whereStr)
	var sqlRes sql.Result
	if t.ctx != nil {
		sqlRes, err = t.conn.ExecCtx(t.ctx, query, params...)
	} else {
		sqlRes, err = t.conn.Exec(query, params...)
	}

	if err != nil {
		return 0, resd.Error(err)
	}
	affectedRow, _ := sqlRes.RowsAffected()
	return affectedRow, nil
}

// prepareUpdate 内部封装，让事务和非事务公共代码复用
func (t *SqlxDao) prepareUpdate(data map[string]any) (string, []any, error) {
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
		return "", nil, resd.NewErr("update data is empty")
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
		return "", nil, resd.NewErr("update param where is empty")
	}
	// 多应用时自动增加多应用条件
	if t.platId != "" {
		whereStr = whereStr + fmt.Sprintf(" AND plat_id=%s", t.platId)
	}
	query := fmt.Sprintf("update %s set %s where %s", t.table, updateStr, whereStr)
	return query, params, nil
}

// WhereId 根据id设置where，优先级最高，执行后将会覆盖原有条件
func (t *SqlxDao) WhereId(id string) *SqlxDao {
	t.whereSql = "id=?"
	t.whereData = append(t.whereData, id)
	return t
}

// Where 在原有where条件上拼接一个 AND (条件)，通过？占位，可防sql注入
func (t *SqlxDao) Where(whereStr string, whereData ...any) *SqlxDao {
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
	t.tableAlias = field
	return t
}

// Field 设置查询字段，若不限制则全字段查询
func (t *SqlxDao) Field(field string) *SqlxDao {
	t.fieldSql = field
	return t
}

// Order 设置排序字段
func (t *SqlxDao) Order(order string) *SqlxDao {
	t.orderSql = order
	return t
}

// Plat 设置应用id
func (t *SqlxDao) Plat(platId string) *SqlxDao {
	t.platId = platId
	return t
}

// GetLastQuerySql 获取最后查询的sql,暂未实现获取
func (t *SqlxDao) GetLastQuerySql() string {
	return t.lastQuerySql
}

// GetLastQueryPageAndSize 获取最后查询的分页页码 和 条数
func (t *SqlxDao) GetLastQueryPageAndSize() (page int64, size int64) {
	return t.lastQueryPage, t.lastQuerySize
}

// GetLastQueryIsCache 获取最后查询的分页页码 和 条数
func (t *SqlxDao) GetLastQueryIsCache() bool {
	return t.lastQueryIsCache
}

// Reinit 每次执行完数据库操作后恢复初始化，保证不干扰下次使用
func (t *SqlxDao) Reinit() {
	t.whereSql = ""
	t.fieldSql = ""
	t.tableAlias = ""
	t.orderSql = ""
	t.whereData = make([]any, 0)
	t.querySize = 0
	t.queryPage = 0
	t.err = nil
	t.ctx = nil
}

// Conn 获取当前的sqlx.SqlConn
func (t *SqlxDao) Conn() sqlx.SqlConn {
	return t.conn
}

func (t *SqlxDao) getCachePrefixField() string {
	return "model_" + strings.Replace(t.table, "`", "", -1)
}
