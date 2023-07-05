package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/asset/api/internal/logic"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func UploadImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadImgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Fail(err.Error()))
			return
		}

		l := logic.NewUploadImgLogic(r.Context(), svcCtx)
		resp, err := l.UploadImg(r, &req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resd.Succ(resp))
		}
	}
}
