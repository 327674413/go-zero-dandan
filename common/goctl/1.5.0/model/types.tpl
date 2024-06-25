const(
    {{.constDatabaseFields}}
)
type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
		Field(field string) *default{{.upperStartCamelObject}}Model
		Except(fields ...string) *default{{.upperStartCamelObject}}Model
        Alias(alias string) *default{{.upperStartCamelObject}}Model
        Where(whereStr string, whereData ...any) *default{{.upperStartCamelObject}}Model
        WhereId(id string) *default{{.upperStartCamelObject}}Model
        Order(order string) *default{{.upperStartCamelObject}}Model
        Limit(num int64) *default{{.upperStartCamelObject}}Model
        Plat(id string) *default{{.upperStartCamelObject}}Model
        Find() (*{{.upperStartCamelObject}}, error)
        FindById(id string) (*{{.upperStartCamelObject}}, error)
        CacheFind(redis *redisd.Redisd) (*{{.upperStartCamelObject}}, error)
        CacheFindById(redis *redisd.Redisd, id string) (*{{.upperStartCamelObject}}, error)
        Page(page int64, rows int64) *default{{.upperStartCamelObject}}Model
        Select() ([]*{{.upperStartCamelObject}}, error)
        SelectWithTotal() ([]*{{.upperStartCamelObject}}, int64, error)
        CacheSelect(redis *redisd.Redisd) ([]*{{.upperStartCamelObject}}, error)
        Count() (int64, error)
        Inc(field string, num int) (int64, error)
        Dec(field string, num int) (int64, error)
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
