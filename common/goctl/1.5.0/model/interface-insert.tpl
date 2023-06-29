Insert(ctx context.Context, data map[string]string) (sql.Result,error)
TxInsert(tx *sql.Tx,ctx context.Context, data map[string]string) (sql.Result,error)