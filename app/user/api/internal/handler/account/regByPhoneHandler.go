package account

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/user/api/internal/logic/account"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func RegByPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegByPhoneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}

		l := account.NewRegByPhoneLogic(r.Context(), svcCtx)
		resp, err := l.RegByPhone(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resd.Succ(resp))
		}
	}
}
