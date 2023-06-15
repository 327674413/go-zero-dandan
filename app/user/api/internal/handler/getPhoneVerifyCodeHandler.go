package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/user/api/internal/logic"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/common/errd"
	"net/http"
)

func getPhoneVerifyCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPhoneVerifyCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, errd.Fail(err.Error()))
			return
		}

		l := logic.NewGetPhoneVerifyCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetPhoneVerifyCode(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, errd.Succ(resp))
		}
	}
}
