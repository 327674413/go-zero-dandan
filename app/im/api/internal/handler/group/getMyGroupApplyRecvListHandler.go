package group

import (
	"go-zero-dandan/common/resd"
	"net/http"

	"go-zero-dandan/app/im/api/internal/logic/group"
	"go-zero-dandan/app/im/api/internal/svc"
)

func GetMyGroupApplyRecvListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := group.NewGetMyGroupApplyRecvListLogic(r.Context(), svcCtx)
		resp, err := l.GetMyGroupApplyRecvList()
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
