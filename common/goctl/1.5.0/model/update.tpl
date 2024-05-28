func (m *default{{.upperStartCamelObject}}Model) Update(data map[string]any) (int64, error) {
    return m.dao.Update(data)
}
func (m *default{{.upperStartCamelObject}}Model) TxUpdate(tx *sql.Tx, data map[string]any) (int64, error) {
    return m.dao.TxUpdate(tx, data)
}
func (m *default{{.upperStartCamelObject}}Model) Save(data *{{.upperStartCamelObject}}) (int64, error) {
    saveData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
   return m.dao.Save(saveData)
}
func (m *default{{.upperStartCamelObject}}Model) TxSave(tx *sql.Tx,data *{{.upperStartCamelObject}}) (int64, error) {
    saveData,err := dao.PrepareData(data)
    if err != nil{
        return 0,err
    }
   return m.dao.Save(saveData)
}