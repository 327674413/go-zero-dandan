package authd

import (
	"github.com/casbin/casbin/v2/model"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/redisd"
)

type CasbinAdapter struct {
	Redis *redisd.Redisd
	Dao   *dao.SqlxDao
}

func (z *CasbinAdapter) LoadPolicy(model model.Model) error {
	//TODO implement me
	panic("implement me")
}

func (z *CasbinAdapter) SavePolicy(model model.Model) error {
	//TODO implement me
	panic("implement me")
}

func (z *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	//TODO implement me
	panic("implement me")
}

func (z *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	//TODO implement me
	panic("implement me")
}

func (z *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	//TODO implement me
	panic("implement me")
}
