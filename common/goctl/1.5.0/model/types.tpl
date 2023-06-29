type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
		Field(field string) *default{{.upperStartCamelObject}}Model
        Alias(alias string) *default{{.upperStartCamelObject}}Model
        WhereStr(whereStr string) *default{{.upperStartCamelObject}}Model
        WhereId(id int) *default{{.upperStartCamelObject}}Model
        WhereMap(whereMap map[string]any) *default{{.upperStartCamelObject}}Model
        WhereRaw(whereStr string, whereData []any) *default{{.upperStartCamelObject}}Model
        Order(order string) *default{{.upperStartCamelObject}}Model
        Plat(id int) *default{{.upperStartCamelObject}}Model
        Find(ctx context.Context, id ...any) (*{{.upperStartCamelObject}}, error)
        CacheFind(ctx context.Context, redis *redisd.Redisd, id ...int64) (*{{.upperStartCamelObject}}, error)
        Page(ctx context.Context, page int, rows int) ([]*{{.upperStartCamelObject}}, error)
        List(ctx context.Context) ([]*{{.upperStartCamelObject}}, error)
        Count(ctx context.Context) int
        Inc(ctx context.Context, field string, num int) error
        Dec(ctx context.Context, field string, num int) error
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table           string
		softDeleteField string
        softDeletable   bool
        fieldSql        string
        whereSql        string
        aliasSql 		string
        orderSql        string
        platId          int64
        whereData       []any
        err             error
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
