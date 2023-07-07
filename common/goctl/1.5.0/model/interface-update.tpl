Update(data map[string]string) (int64,error)
TxUpdate(tx *sql.Tx, data map[string]string) (int64,error)
Save(data map[string]string) (int64,error)
TxSave(tx *sql.Tx, data map[string]string) (int64,error)