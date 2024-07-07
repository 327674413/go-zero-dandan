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