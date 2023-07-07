func (m *default{{.upperStartCamelObject}}Model) Find( id ...any) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
	err := m.dao.Find(resp, id...)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) CacheFind( redis *redisd.Redisd, id ...int64) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
    err := m.dao.CacheFind(redis, resp, id...)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) Select() ([]*{{.upperStartCamelObject}},error) {
	var resp []*{{.upperStartCamelObject}}
	err := m.dao.Select(resp)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) Page(page int, rows int) *default{{.upperStartCamelObject}}Model {
    m.dao.Page(page,rows)
    return m
}