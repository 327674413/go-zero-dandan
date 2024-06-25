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