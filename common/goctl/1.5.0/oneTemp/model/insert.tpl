func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    data.CreateAt = time.Now().Unix()
    data.UpdateAt = time.Now().Unix()
    return m.conn.ExecCtx(ctx, query, {{.expressionValues}})
}
func (m *default{{.upperStartCamelObject}}Model) TxInsert(tx *sql.Tx,ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    data.CreateAt = time.Now().Unix()
    data.UpdateAt = time.Now().Unix()
    return tx.ExecContext(ctx, query, {{.expressionValues}})
}