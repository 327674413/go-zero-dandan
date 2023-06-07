type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
		softDeleteField string
        softDeleteState bool
        fieldSql        string
        whereSql        string
        aliasSql 		string
        orderSql        string
        whereData       []any
        err             error
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
