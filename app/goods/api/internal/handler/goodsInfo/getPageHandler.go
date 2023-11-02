package goodsInfo

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/goods/api/internal/logic/goodsInfo"
	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"
	"go-zero-dandan/common/land"
	"go-zero-dandan/common/resd"
	"net/http"
)

func GetPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}

		l := goodsInfo.NewGetPageLogic(r.Context(), svcCtx)
		resp, err := l.GetPage(&req)
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
