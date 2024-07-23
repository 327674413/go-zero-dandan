// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(ctxOrNil context.Context,conn sqlx.SqlConn,platId ...string) {{.upperStartCamelObject}}Model {
	var platid string
    if len(platId) > 0 {
        platid = platId[0]
    } else {
        platid = ""
    }
    if ctxOrNil == nil {
        ctxOrNil = context.Background()
    }
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(ctxOrNil,conn,platid),
		softDeletable:softDeletable{{.upperStartCamelObject}},
	}
}
func new{{.upperStartCamelObject}}Model(ctx context.Context,conn sqlx.SqlConn,platId string) *default{{.upperStartCamelObject}}Model {
	dao := dao.NewSqlxDao(conn, {{.table}}, default{{.upperStartCamelObject}}Fields, true, "delete_at")
	dao.Plat(platId)
	dao.Ctx(ctx)
	return &default{{.upperStartCamelObject}}Model{
	    ctx:ctx,
		conn:       conn,
		dao:        dao,
		table:      {{.table}},
		platId:     platId,
        softDeleteField: "delete_at",
        whereData:       make([]any, 0),

	}
}
func (m *default{{.upperStartCamelObject}}Model) Ctx(ctx context.Context) *default{{.upperStartCamelObject}}Model {
	m.dao.Ctx(ctx)
	return m
}
func (m *default{{.upperStartCamelObject}}Model) WhereId(id string) *default{{.upperStartCamelObject}}Model {
	m.dao.WhereId(id)
    return m
}

func (m *default{{.upperStartCamelObject}}Model) Where(whereStr string, whereData ...any) *default{{.upperStartCamelObject}}Model {
	m.dao.Where(whereStr, whereData...)
    return m
}

func (m *default{{.upperStartCamelObject}}Model) Alias(alias string) *default{{.upperStartCamelObject}}Model {
	m.dao.Alias(alias)
    return m
}
func (m *default{{.upperStartCamelObject}}Model) Field(field string) *default{{.upperStartCamelObject}}Model {
	m.dao.Field(field)
    return m
}
func (m *default{{.upperStartCamelObject}}Model) Except(fields ...string) *default{{.upperStartCamelObject}}Model {
	m.dao.Except(fields...)
    return m
}
func (m *default{{.upperStartCamelObject}}Model) Order(order string) *default{{.upperStartCamelObject}}Model {
    m.dao.Order(order)
	return m
}
func (m *default{{.upperStartCamelObject}}Model) Limit(num int64) *default{{.upperStartCamelObject}}Model {
    m.dao.Limit(num)
	return m
}
func (m *default{{.upperStartCamelObject}}Model) Count() (int64, error) {
    return m.dao.Count()
}
func (m *default{{.upperStartCamelObject}}Model) Inc(field string, num int) (int64, error) {
    return m.dao.Inc(field, num)
}
func (m *default{{.upperStartCamelObject}}Model) TxInc(tx *sql.Tx, field string, num int) (int64, error) {
	return m.dao.TxInc(tx, field, num)
}
func (m *default{{.upperStartCamelObject}}Model) Dec(field string, num int) (int64, error) {
    return m.dao.Dec(field, num)
}
func (m *default{{.upperStartCamelObject}}Model) TxDec(tx *sql.Tx, field string, num int) (int64, error) {
    return m.dao.Dec(field, num)
}
func (m *default{{.upperStartCamelObject}}Model) Plat(id string) *default{{.upperStartCamelObject}}Model {
    m.dao.Plat(id)
    return m
}
func (m *default{{.upperStartCamelObject}}Model) Reinit() *default{{.upperStartCamelObject}}Model {
	m.dao.Reinit()
	return m
}
func (m *default{{.upperStartCamelObject}}Model) Dao() *dao.SqlxDao {
	return m.dao
}
func (m *default{{.upperStartCamelObject}}Model) Find() (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
	err := m.dao.Find(resp)
    if err != nil {
        if err == sqlx.ErrNotFound {
            return nil,nil
        }
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) FindById(id string) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
	err := m.dao.FindById(resp,id)
	if err != nil {
        if err == sqlx.ErrNotFound {
            return nil,nil
        }
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) CacheFind( redis *redisd.Redisd) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
    err := m.dao.CacheFind(redis, resp)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) CacheFindById( redis *redisd.Redisd, id string) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
    err := m.dao.CacheFindById(redis, resp, id)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) Total() (total int64,danErr error) {
	return m.dao.Total()
}
func (m *default{{.upperStartCamelObject}}Model) Select() ([]*{{.upperStartCamelObject}},error) {
	resp := make([]*{{.upperStartCamelObject}},0)
	err := m.dao.Select(&resp)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) SelectWithTotal() ([]*{{.upperStartCamelObject}},int64,error) {
	resp := make([]*{{.upperStartCamelObject}},0)
	var total int64
	err := m.dao.Select(&resp,&total)
    if err != nil {
        return nil,0, err
    }
    return resp, total,nil
}
func (m *default{{.upperStartCamelObject}}Model) CacheSelect(redis *redisd.Redisd) ([]*{{.upperStartCamelObject}},error) {
	resp := make([]*{{.upperStartCamelObject}},0)
	err := m.dao.CacheSelect(redis,&resp)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

func (m *default{{.upperStartCamelObject}}Model) Page(page int64, size int64) *default{{.upperStartCamelObject}}Model {
    m.dao.Page(page,size)
    return m
}
func (m *default{{.upperStartCamelObject}}Model) Insert(data *{{.upperStartCamelObject}}) (effectRow int64, danErr error) {
    insertData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
    return m.dao.Insert(insertData)
}
func (m *default{{.upperStartCamelObject}}Model) TxInsert(tx *sql.Tx, data *{{.upperStartCamelObject}}) (effectRow int64,danErr error) {
    insertData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
    return m.dao.TxInsert(tx, insertData)
}
func (m *default{{.upperStartCamelObject}}Model) Delete(id ...string) (effectRow int64,danErr error) {
    return m.dao.Delete(id...)
}
func (m *default{{.upperStartCamelObject}}Model) TxDelete(tx *sql.Tx,id ...string) (effectRow int64,danErr error) {
    return m.dao.TxDelete(tx,id...)
}
func (m *default{{.upperStartCamelObject}}Model) Update(data map[dao.TableField]any) (effectRow int64,err error) {
    return m.dao.Update(data)
}
func (m *default{{.upperStartCamelObject}}Model) TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64,err error) {
    return m.dao.TxUpdate(tx, data)
}
func (m *default{{.upperStartCamelObject}}Model) Save(data *{{.upperStartCamelObject}}) (effectRow int64,err error) {
    saveData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
   return m.dao.Save(saveData)
}
func (m *default{{.upperStartCamelObject}}Model) TxSave(tx *sql.Tx,data *{{.upperStartCamelObject}}) (effectRow int64,err error) {
    saveData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
   return m.dao.Save(saveData)
}
func (m *default{{.upperStartCamelObject}}Model) StartTrans() (tx *sql.Tx,danErr error) {
    return dao.StartTrans(m.ctx,m.conn)
}
func (m *default{{.upperStartCamelObject}}Model) Commit(tx *sql.Tx) (danErr error) {
    return dao.Commit(tx)
}
func (m *default{{.upperStartCamelObject}}Model) Rollback(tx *sql.Tx) (danErr error) {
    return dao.Rollback(tx)
}
func (m *default{{.upperStartCamelObject}}Model) tableName() string {
	return m.table
}
// forGoctl 避免有的model没有time.Time类型时，goctl生成模版会因引入未使用的包而报错
func (m *default{{.upperStartCamelObject}}Model) forGoctl() {
	t := time.Time{}
	fmt.Println(t)
}
