// Code generated by goctl. DO NOT EDIT.
package wxpub

import (
	"context"

	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type JssdkBuildLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		Url         string   `json:"url"`
		JsApiList   []string `json:"jsApiList"`
		OpenTagList []string `json:"openTagList"`
	}
	hasReq struct {
		Url         bool
		JsApiList   bool
		OpenTagList bool
	}
}

func NewJssdkBuildLogicGen(ctx context.Context, svc *svc.ServiceContext) *JssdkBuildLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &JssdkBuildLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, resd.I18n.NewLang(meta.Lang)),
		meta:   meta,
	}
}

func (l *JssdkBuildLogicGen) initReq(req *types.JssdkBuildReq) error {

	if req.Url != nil {
		l.req.Url = strings.TrimSpace(*req.Url)
		l.hasReq.Url = true
	} else {
		l.hasReq.Url = false
	}

	if req.JsApiList != nil {
		l.req.JsApiList = req.JsApiList
		l.hasReq.JsApiList = true
	} else {
		l.hasReq.JsApiList = false
	}

	if req.OpenTagList != nil {
		l.req.OpenTagList = req.OpenTagList
		l.hasReq.OpenTagList = true
	} else {
		l.hasReq.OpenTagList = false
	}

	return nil
}
