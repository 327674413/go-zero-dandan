func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, data map[string]string) error {
    return m.dao.Update(ctx,data)
}
func (m *default{{.upperStartCamelObject}}Model) TxUpdate(tx *sql.Tx,ctx context.Context, data map[string]string) error {
    return m.dao.TxUpdate(tx,ctx,data)
}
func (m *default{{.upperStartCamelObject}}Model) Save(ctx context.Context, data map[string]string) error {
    var updateStr string
    params := make([]any, 0)
    var id int64
    var hasId bool
    for k, v := range data {
        if k == "Id" {
            id = utild.AnyToInt64(k)
            hasId = true
            continue
        }
        updateStr = updateStr + fmt.Sprintf("%s=?,", utild.StrToSnake(k))
        params = append(params, v)
    }
    if hasId == false{
        return errors.New("save data must need Id")
    }
    currData,err := m.Find(ctx,id)
    if err != nil{
        return err
    }
    if currData.Id == 0 {
        _,err = m.Insert(ctx, data)
        if err != nil{
            return err
        }
        return nil
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
            return errors.New("update data must need where")
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
}