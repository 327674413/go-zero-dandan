func (m *default{{.upperStartCamelObject}}Model) Insert(data *{{.upperStartCamelObject}}) (int64, error) {
    insertData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
    return m.dao.Insert(insertData)
}
func (m *default{{.upperStartCamelObject}}Model) TxInsert(tx *sql.Tx, data *{{.upperStartCamelObject}}) (int64, error) {
    insertData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
    return m.dao.TxInsert(tx, insertData)
}
