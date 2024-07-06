package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/asset/api/internal/logic"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func MultipartUploadSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MultipartUploadSendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := logic.NewMultipartUploadSendLogic(r.Context(), svcCtx)
		resp, err := l.MultipartUploadSend(&req)
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
