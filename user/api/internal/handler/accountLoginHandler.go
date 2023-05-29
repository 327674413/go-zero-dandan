package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/common/api"
	"go-zero-dandan/user/api/internal/logic"
	"go-zero-dandan/user/api/internal/svc"
	"go-zero-dandan/user/api/internal/types"
	"net/http"
)

func AccountLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccountLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, api.Fail(err.Error()))
			return
		}

		l := logic.NewAccountLoginLogic(r.Context(), svcCtx)
		resp, err := l.AccountLogin(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, api.Fail(err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
