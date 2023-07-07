func (m *default{{.upperStartCamelObject}}Model) Insert(data map[string]string) (int64, error) {
    return m.dao.Insert(data)
}
func (m *default{{.upperStartCamelObject}}Model) TxInsert(tx *sql.Tx, data map[string]string) (int64, error) {
    return m.dao.TxInsert(tx, data)
}
