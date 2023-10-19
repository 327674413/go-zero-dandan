package wxpub

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/wechat/api/internal/logic/wxpub"
	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"
	"go-zero-dandan/common/land"
	"go-zero-dandan/common/resd"
	"net/http"
)

func ServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WxpubReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := wxpub.NewServiceLogic(r.Context(), svcCtx)
		resp, err := l.Service(&req, r)
		if err != nil {
			if danErr, ok := resd.AssertErr(err); ok {
				localizer, ok := r.Context().Value("lang").(*i18n.Localizer)
				if ok {
					danErr.Msg = land.Msg(localizer, danErr.Code, danErr.GetTemps())
				}
				httpx.OkJsonCtx(r.Context(), w, danErr)
			} else {
				httpx.OkJsonCtx(r.Context(), w, err)
			}
		} else {
			httpx.OkJsonCtx(r.Context(), w, resd.Succ(resp))
		}
	}
}
