package conversation

import (
	"go-zero-dandan/common/resd"
	"net/http"

	"go-zero-dandan/app/im/api/internal/logic/conversation"
	"go-zero-dandan/app/im/api/internal/svc"
)

func GetConversationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := conversation.NewGetConversationListLogic(r.Context(), svcCtx)
		resp, err := l.GetConversationList()
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
