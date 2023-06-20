func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{.keys}}
	data.CreateAt = time.Now().Unix()
    data.UpdateAt = time.Now().Unix()
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    data.CreateAt = time.Now().Unix()
    data.UpdateAt = time.Now().Unix()
    ret,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return ret,err
}
