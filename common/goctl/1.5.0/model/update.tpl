func (m *default{{.upperStartCamelObject}}Model) Update(data map[string]string) (int64, error) {
    return m.dao.Update(data)
}
func (m *default{{.upperStartCamelObject}}Model) TxUpdate(tx *sql.Tx, data map[string]string) (int64, error) {
    return m.dao.TxUpdate(tx, data)
}
func (m *default{{.upperStartCamelObject}}Model) Save(data map[string]string) (int64, error) {
   return m.dao.Save(data)
}
func (m *default{{.upperStartCamelObject}}Model) TxSave(tx *sql.Tx,data map[string]string) (int64, error) {
   return m.dao.Save(data)
}