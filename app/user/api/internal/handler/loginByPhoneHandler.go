package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/user/api/internal/logic"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/common/respd"
	"net/http"
)

func LoginByPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginByPhoneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, respd.Fail(err.Error()))
			return
		}

		l := logic.NewLoginByPhoneLogic(r.Context(), svcCtx)
		resp, err := l.LoginByPhone(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, respd.Succ(resp))
		}
	}
}
