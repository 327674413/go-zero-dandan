package wxpub

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/wechat/api/internal/logic/wxpub"
	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func AuthByCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthByCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := wxpub.NewAuthByCodeLogic(r.Context(), svcCtx)
		resp, err := l.AuthByCode(&req)
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
