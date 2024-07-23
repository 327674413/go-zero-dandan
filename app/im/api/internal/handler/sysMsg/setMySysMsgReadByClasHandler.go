package sysMsg

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/im/api/internal/logic/sysMsg"
	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func SetMySysMsgReadByClasHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetMySysMsgReadByClasReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := sysMsg.NewSetMySysMsgReadByClasLogic(r.Context(), svcCtx)
		resp, err := l.SetMySysMsgReadByClas(&req)
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
