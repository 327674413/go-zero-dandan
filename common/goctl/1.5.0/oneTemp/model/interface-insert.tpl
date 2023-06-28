Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error)
TxInsert(tx *sql.Tx,ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error)