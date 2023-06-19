package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/plat/api/internal/logic"
	"go-zero-dandan/app/plat/api/internal/svc"
	"go-zero-dandan/app/plat/api/internal/types"
	"go-zero-dandan/common/respd"
	"net/http"
)

func GetTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, respd.Fail(err.Error()))
			return
		}

		l := logic.NewGetTokenLogic(r.Context(), svcCtx)
		resp, err := l.GetToken(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, respd.Succ(resp))
		}
	}
}
