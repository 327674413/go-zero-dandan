package userInfo

import (
	"go-zero-dandan/common/resd"
	"net/http"

	"go-zero-dandan/app/user/api/internal/logic/userInfo"
	"go-zero-dandan/app/user/api/internal/svc"
)

func GetMyUserMainInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userInfo.NewGetMyUserMainInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetMyUserMainInfo()
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
