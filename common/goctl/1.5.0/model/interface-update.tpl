Update(data map[dao.TableField]any) (effectRow int64,err error)
TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64,err error)
Save(data *{{.upperStartCamelObject}}) (effectRow int64,err error)
TxSave(tx *sql.Tx, data *{{.upperStartCamelObject}}) (effectRow int64,err error)