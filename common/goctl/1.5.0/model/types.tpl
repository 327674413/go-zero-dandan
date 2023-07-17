type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
		Field(field string) *default{{.upperStartCamelObject}}Model
        Alias(alias string) *default{{.upperStartCamelObject}}Model
        WhereStr(whereStr string) *default{{.upperStartCamelObject}}Model
        WhereId(id int64) *default{{.upperStartCamelObject}}Model
        WhereRaw(whereStr string, whereData []any) *default{{.upperStartCamelObject}}Model
        Order(order string) *default{{.upperStartCamelObject}}Model
        Plat(id int64) *default{{.upperStartCamelObject}}Model
        Find() (*{{.upperStartCamelObject}}, error)
        FindById(id int64) (*{{.upperStartCamelObject}}, error)
        CacheFind(redis *redisd.Redisd) (*{{.upperStartCamelObject}}, error)
        CacheFindById(redis *redisd.Redisd, id int64) (*{{.upperStartCamelObject}}, error)
        Page(page int64, rows int64) *default{{.upperStartCamelObject}}Model
        Select() ([]*{{.upperStartCamelObject}}, error)
        CacheSelect(redis *redisd.Redisd) ([]*{{.upperStartCamelObject}}, error)
        Count() (int64, error)
        Inc(field string, num int) (int64, error)
        Dec(field string, num int) (int64, error)
        Ctx(ctx context.Context) *default{{.upperStartCamelObject}}Model
        Reinit() *default{{.upperStartCamelObject}}Model
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
        platId          int64
        whereData       []any
        err             error
        ctx             context.Context
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
