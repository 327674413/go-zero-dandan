package friend

import (
	"go-zero-dandan/common/resd"
	"net/http"

	"go-zero-dandan/app/im/api/internal/logic/friend"
	"go-zero-dandan/app/im/api/internal/svc"
)

func GetMyFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := friend.NewGetMyFriendListLogic(r.Context(), svcCtx)
		resp, err := l.GetMyFriendList()
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
