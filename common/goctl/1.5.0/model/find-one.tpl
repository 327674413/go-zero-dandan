func (m *default{{.upperStartCamelObject}}Model) Find() (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
	err := m.dao.Find(resp)
    if err != nil {
        if err == sqlx.ErrNotFound {
            return nil,nil
        }
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) FindById(id string) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
	err := m.dao.FindById(resp,id)
	if err != nil {
        if err == sqlx.ErrNotFound {
            return nil,nil
        }
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) CacheFind( redis *redisd.Redisd) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
    err := m.dao.CacheFind(redis, resp)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) CacheFindById( redis *redisd.Redisd, id string) (*{{.upperStartCamelObject}}, error) {
	resp := &{{.upperStartCamelObject}}{}
    err := m.dao.CacheFindById(redis, resp, id)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) Total() (total int64,err error) {
	return m.dao.Total()
}
func (m *default{{.upperStartCamelObject}}Model) Select() ([]*{{.upperStartCamelObject}},error) {
	resp := make([]*{{.upperStartCamelObject}},0)
	err := m.dao.Select(&resp)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
func (m *default{{.upperStartCamelObject}}Model) SelectWithTotal() ([]*{{.upperStartCamelObject}},int64,error) {
	resp := make([]*{{.upperStartCamelObject}},0)
	var total int64
	err := m.dao.Select(&resp,&total)
    if err != nil {
        return nil,0, err
    }
    return resp, total,nil
}
func (m *default{{.upperStartCamelObject}}Model) CacheSelect(redis *redisd.Redisd) ([]*{{.upperStartCamelObject}},error) {
	resp := make([]*{{.upperStartCamelObject}},0)
	err := m.dao.CacheSelect(redis,&resp)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

func (m *default{{.upperStartCamelObject}}Model) Page(page int64, size int64) *default{{.upperStartCamelObject}}Model {
    m.dao.Page(page,size)
    return m
}