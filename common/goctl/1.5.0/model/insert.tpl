func (m *default{{.upperStartCamelObject}}Model) Insert(data *{{.upperStartCamelObject}}) (effectRow int64, err error) {
    insertData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
    return m.dao.Insert(insertData)
}
func (m *default{{.upperStartCamelObject}}Model) TxInsert(tx *sql.Tx, data *{{.upperStartCamelObject}}) (effectRow int64,err error) {
    insertData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
    return m.dao.TxInsert(tx, insertData)
}
