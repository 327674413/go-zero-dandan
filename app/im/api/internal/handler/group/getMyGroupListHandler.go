package group

import (
	"go-zero-dandan/common/resd"
	"net/http"

	"go-zero-dandan/app/im/api/internal/logic/group"
	"go-zero-dandan/app/im/api/internal/svc"
)

func GetMyGroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := group.NewGetMyGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GetMyGroupList()
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
