// Code generated by goctl. DO NOT EDIT.
package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type CreateGroupLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		Name string `json:"name,optional"`
	}
	hasReq struct {
		Name bool
	}
}

func NewCreateGroupLogicGen(ctx context.Context, svc *svc.ServiceContext) *CreateGroupLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &CreateGroupLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *CreateGroupLogicGen) initReq(in *types.CreateGroupReq) error {

	if in.Name != nil {
		l.req.Name = strings.TrimSpace(*in.Name)
		l.hasReq.Name = true
	} else {
		l.hasReq.Name = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}
