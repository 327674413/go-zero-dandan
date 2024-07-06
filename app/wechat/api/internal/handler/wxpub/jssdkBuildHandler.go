package wxpub

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/wechat/api/internal/logic/wxpub"
	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func JssdkBuildHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JssdkBuildReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := wxpub.NewJssdkBuildLogic(r.Context(), svcCtx)
		resp, err := l.JssdkBuild(&req)
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
