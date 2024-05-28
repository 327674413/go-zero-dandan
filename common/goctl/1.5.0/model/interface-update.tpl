Update(data map[string]any) (int64,error)
TxUpdate(tx *sql.Tx, data map[string]any) (int64,error)
Save(data *{{.upperStartCamelObject}}) (int64,error)
TxSave(tx *sql.Tx, data *{{.upperStartCamelObject}}) (int64,error)