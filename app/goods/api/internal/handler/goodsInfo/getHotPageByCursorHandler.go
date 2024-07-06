package goodsInfo

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/goods/api/internal/logic/goodsInfo"
	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func GetHotPageByCursorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetHotPageByCursorReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := goodsInfo.NewGetHotPageByCursorLogic(r.Context(), svcCtx)
		resp, err := l.GetHotPageByCursor(&req)
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
