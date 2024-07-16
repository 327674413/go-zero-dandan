// Code generated by goctl. DO NOT EDIT.
package friend

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type SearchNewFriendPageLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		Keyword string `json:"keyword,optional"`
	}
	hasReq struct {
		Keyword bool
	}
}

func NewSearchNewFriendPageLogicGen(ctx context.Context, svc *svc.ServiceContext) *SearchNewFriendPageLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &SearchNewFriendPageLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, resd.I18n.NewLang(meta.Lang)),
		meta:   meta,
	}
}

func (l *SearchNewFriendPageLogicGen) initReq(req *types.SearchNewFriendReq) error {

	if req.Keyword != nil {
		l.req.Keyword = strings.TrimSpace(*req.Keyword)
		l.hasReq.Keyword = true
	} else {
		l.hasReq.Keyword = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}
