package account

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/user/api/internal/logic/account"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func GetPhoneVerifyCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPhoneVerifyCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := account.NewGetPhoneVerifyCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetPhoneVerifyCode(&req)
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
