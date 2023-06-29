func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data map[string]string) (sql.Result,error) {
	var insertField string
	var insertValue string
    params := make([]any, 0)
    var id int64
    for k, v := range data {
        if k == "Id" {
            id = utild.AnyToInt64(k)
            continue
        }
        updateStr = updateStr + fmt.Sprintf("%s=?,", utild.StrToSnake(k))
        params = append(params, v)
    }
    if len(updateStr) > 0 {
        updateStr = updateStr[:len(updateStr)-1]
    } else {
        return errors.New("update data is empty")
    }
    updateStr = updateStr + fmt.Sprintf(",update_at=%d", utild.GetStamp())
    whereStr := m.whereSql
    if whereStr == "" {
        if id == 0 {
            return errors.New("update param where is empty")
        } else {
            whereStr = fmt.Sprintf("id=%d", id)
        }

    }
    if m.platId != 0 {
        whereStr = whereStr + fmt.Sprintf(" AND plat_id=%d", m.platId)
    }
    query := fmt.Sprintf("update %s set %s where %s", m.table, updateStr, whereStr)
    _, err = m.conn.ExecCtx(ctx, query, params...)
    m.Reinit()
    return err
    return m.conn.ExecCtx(ctx, query, {{.expressionValues}})
}
func (m *default{{.upperStartCamelObject}}Model) TxInsert(tx *sql.Tx, ctx context.Context, data map[string]string) (sql.Result,error) {
	query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    data.CreateAt = time.Now().Unix()
    data.UpdateAt = time.Now().Unix()
    data.PlatId = m.platId
    return tx.ExecContext(ctx, query, {{.expressionValues}})
}
