const(
    {{.constDatabaseFields}}
)
type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
        Delete(id ...string) (effectRow int64, danErr error)
        TxDelete(tx *sql.Tx,id ...string) (effectRow int64, danErr error)
        Insert(data *{{.upperStartCamelObject}}) (effectRow int64, danErr error)
        TxInsert(tx *sql.Tx,data *{{.upperStartCamelObject}}) (effectRow int64, danErr error)
        Update(data map[dao.TableField]any) (effectRow int64,danErr error)
        TxUpdate(tx *sql.Tx, data map[dao.TableField]any) (effectRow int64,danErr error)
        Save(data *{{.upperStartCamelObject}}) (effectRow int64,danErr error)
        TxSave(tx *sql.Tx, data *{{.upperStartCamelObject}}) (effectRow int64,danErr error)
		Field(field string) *default{{.upperStartCamelObject}}Model
		Except(fields ...string) *default{{.upperStartCamelObject}}Model
        Alias(alias string) *default{{.upperStartCamelObject}}Model
        LeftJoin(joinTable string) *default{{.upperStartCamelObject}}Model
        RightJoin(joinTable string) *default{{.upperStartCamelObject}}Model
        InnerJoin(joinTable string) *default{{.upperStartCamelObject}}Model
        Where(whereStr string, whereData ...any) *default{{.upperStartCamelObject}}Model
        WhereId(id string) *default{{.upperStartCamelObject}}Model
        Order(order string) *default{{.upperStartCamelObject}}Model
        Limit(num int64) *default{{.upperStartCamelObject}}Model
        Plat(id string) *default{{.upperStartCamelObject}}Model
        Find() (*{{.upperStartCamelObject}}, error)
        FindById(id string) (data *{{.upperStartCamelObject}}, danErr error)
        CacheFind(redis *redisd.Redisd) (data *{{.upperStartCamelObject}}, danErr error)
        CacheFindById(redis *redisd.Redisd, id string) (data *{{.upperStartCamelObject}}, danErr error)
        Page(page int64, rows int64) *default{{.upperStartCamelObject}}Model
        Total() (total int64,danErr error)
        Select() (dataList []*{{.upperStartCamelObject}}, danErr error)
        SelectWithTotal() (dataList []*{{.upperStartCamelObject}}, total int64, danErr error)
        CacheSelect(redis *redisd.Redisd) (dataList []*{{.upperStartCamelObject}}, danErr error)
        Count() (total int64, danErr error)
        Inc(field string, num int) (effectRow int64, danErr error)
        Dec(field string, num int) (effectRow int64, danErr error)
        StartTrans() (tx *sql.Tx, danErr error)
        Commit(tx *sql.Tx) (danErr error)
        Rollback(tx *sql.Tx) (danErr error)
        Ctx(ctx context.Context) *default{{.upperStartCamelObject}}Model
        Reinit() *default{{.upperStartCamelObject}}Model
        Dao() *dao.SqlxDao
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table           string
		dao             *dao.SqlxDao
		softDeleteField string
        softDeletable   bool
        fieldSql        string
        whereSql        string
        aliasSql 		string
        orderSql        string
        platId          string
        whereData       []any
        err             error
        ctx             context.Context
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
