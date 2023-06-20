func (m *default{{.upperStartCamelObject}}Model) Find(ctx context.Context, id ...any) (*{{.upperStartCamelObject}}, error) {
	defer m.Reinit()
	var err error
	if err = m.err; err != nil {
		m.err = nil
		return nil, err
	}
	var resp {{.upperStartCamelObject}}
	var sql string
    field := {{.lowerStartCamelObject}}Rows
	if m.fieldSql != "" {
		field = m.fieldSql
	}
	if len(id) > 0 {
		sql = fmt.Sprintf("select %s from %s where id=? limit 1", field, m.table)
		err = m.conn.QueryRowPartialCtx(ctx, &resp, sql, id[0]) //QueryRowCtx 必须字段都覆盖
	} else {
	    if m.whereSql == ""{
            m.whereSql = "1=1"
        }
		sql = fmt.Sprintf("select %s from %s %s where "+m.whereSql+" limit 1", field, m.table, m.aliasSql)
		err = m.conn.QueryRowPartialCtx(ctx, &resp, sql, m.whereData...)
	}
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return &resp, nil
	default:
		return nil, err
	}
}
func (m *default{{.upperStartCamelObject}}Model) List(ctx context.Context) ([]*{{.upperStartCamelObject}},error) {
	defer m.Reinit()
	var err error
	if err = m.err; err != nil {
		m.err = nil
		return nil, err
	}
	var resp []*{{.upperStartCamelObject}}
	var sql string
    field := {{.lowerStartCamelObject}}Rows
	if m.fieldSql != "" {
		field = m.fieldSql
	}
    if m.whereSql == ""{
        m.whereSql = "1=1"
    }
    sql = fmt.Sprintf("select %s from %s %s where "+m.whereSql, field, m.table, m.aliasSql)
    err = m.conn.QueryRowsPartialCtx(ctx, &resp, sql, m.whereData...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return resp, nil
	default:
		return nil, err
	}
}
func (m *default{{.upperStartCamelObject}}Model) Page(ctx context.Context, page int, rows int) ([]*{{.upperStartCamelObject}}, error) {
    defer m.Reinit()
	return nil, nil
}
func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query :=  fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		return conn.QueryRowCtx(ctx, v, query, {{.lowerStartCamelPrimaryKey}})
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}query := fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
	var resp {{.upperStartCamelObject}}
	err := m.conn.QueryRowCtx(ctx, &resp, query, {{.lowerStartCamelPrimaryKey}})
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}
func (m *default{{.upperStartCamelObject}}Model) Reinit(){
    m.whereSql = ""
    m.fieldSql = ""
    m.aliasSql = ""
    m.orderSql = ""
    m.whereData = make([]any, 0)
    m.err = nil
}