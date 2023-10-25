package authd

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	stringadapter "github.com/casbin/casbin/v2/persist/string-adapter"
	"go-zero-dandan/common/resd"
	"strings"
)

type Auth struct {
	ctx       context.Context
	ModelType ModelType
	Enforcer  *casbin.Enforcer
}

// NewAuth 创建权限
func NewAuth(ctx context.Context, modelType ModelType, policyLines string) (*Auth, error) {
	t := &Auth{
		ModelType: modelType,
		ctx:       ctx,
	}
	model, err := model.NewModelFromString(string(t.ModelType))
	if err != nil {
		return nil, resd.ErrorCtx(t.ctx, err)
	}
	e, err := casbin.NewEnforcer(model, stringadapter.NewAdapter(policyLines))
	if err != nil {
		return nil, resd.ErrorCtx(t.ctx, err)
	}
	t.Enforcer = e
	return t, nil
}

// Check 校验权限是否通过
func (t *Auth) Check(sub string, obj string, act string) (bool, error) {
	return t.Enforcer.Enforce(sub, obj, act)
}

/*
// 自己用的话，应该根据角色id来判断
func (t *Auth) Check(roleId int64, obj string, act string) (bool, error) {
	return t.Enforcer.Enforce(string(roleId), obj, act)
}
*/

// PolicyToLine 将二维数组的权限规则转成一行一行的casbin规则
func PolicyToLine(policy [][]string) string {
	var sb strings.Builder
	for _, p := range policy {
		line := strings.Join(p, ", ")
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	return sb.String()
}
