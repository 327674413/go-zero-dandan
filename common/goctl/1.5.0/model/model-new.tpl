func new{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) *default{{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: sqlc.NewConn(conn, c){{else}}conn:conn{{end}},
		table:      {{.table}},
        softDeleteField: "delete_at",
        whereData:       make([]any, 0),
	}
}

func (m *default{{.upperStartCamelObject}}Model) WhereId(id int) *default{{.upperStartCamelObject}}Model {
	m.whereSql = "id=?"
	m.whereData = append(m.whereData, id)
	return m
}
func (m *default{{.upperStartCamelObject}}Model) Page(ctx context.Context, page int, rows int) ([]*{{.upperStartCamelObject}}, error) {

	return nil, nil
}
func (m *default{{.upperStartCamelObject}}Model) List(ctx context.Context) ([]*{{.upperStartCamelObject}}, error) {

	return nil, nil
}
func (m *default{{.upperStartCamelObject}}Model) WhereStr(whereStr string) *default{{.upperStartCamelObject}}Model {
	return m
}

func (m *default{{.upperStartCamelObject}}Model) WhereMap(whereMap map[string]any) *default{{.upperStartCamelObject}}Model {
	return m
}
func (m *default{{.upperStartCamelObject}}Model) WhereRaw(whereStr string, whereData []any) *default{{.upperStartCamelObject}}Model {
	if m.whereSql != "" {
		m.whereSql += " AND (" + whereStr + ")"
	} else {
		m.whereSql = "(" + whereStr + ")"
	}
	m.whereData = append(m.whereData, whereData...)
	return m
}
func (m *default{{.upperStartCamelObject}}Model) Alias(field string) *default{{.upperStartCamelObject}}Model {
	m.aliasSql = field
	return m
}
func (m *default{{.upperStartCamelObject}}Model) Field(field string) *default{{.upperStartCamelObject}}Model {
	m.fieldSql = field
	return m
}
func (m *default{{.upperStartCamelObject}}Model) Order(order string) *default{{.upperStartCamelObject}}Model {

	return m
}
func (m *default{{.upperStartCamelObject}}Model) Count(ctx context.Context) int {

	return 0
}
func (m *default{{.upperStartCamelObject}}Model) Inc(ctx context.Context, field string, num int) error {

	return nil
}
func (m *default{{.upperStartCamelObject}}Model) Dec(ctx context.Context, field string, num int) error {

	return nil
}
func (m *default{{.upperStartCamelObject}}Model) Plat(id int) *default{{.upperStartCamelObject}}Model {

	return nil
}
func (m *default{{.upperStartCamelObject}}Model) CacheFind(ctx context.Context, redis *redisd.Redisd, id ...int64) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
	cacheField := "model_" + m.tableName()
	cacheKey := strconv.FormatInt(id[0], 10)
	// todo::需要把where条件一起放进去作为key，这样就能支持更多的自动缓存查询
	err := redis.GetData(cacheField, cacheKey, resp)
	if err == nil {
		return resp, nil
	}
	resp, err = m.Find(ctx, id[0])
	fmt.Println(resp, err)
	if err != nil {
		return resp, err
	}
	if resp.Id != 0 {
		redis.SetData(cacheField, cacheKey, resp)
	}
	return resp, nil
}
