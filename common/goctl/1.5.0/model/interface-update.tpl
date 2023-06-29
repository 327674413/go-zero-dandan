Update(ctx context.Context, data map[string]string) error
TxUpdate(tx *sql.Tx, ctx context.Context, data map[string]string) error
Save(ctx context.Context, data map[string]string) error
TxSave(tx *sql.Tx, ctx context.Context, data map[string]string) error