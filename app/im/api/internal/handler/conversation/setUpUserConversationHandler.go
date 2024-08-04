package conversation

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/im/api/internal/logic/conversation"
	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func SetUpUserConversationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUpUserConversationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := conversation.NewSetUpUserConversationLogic(r.Context(), svcCtx)
		resp, err := l.SetUpUserConversation(&req)
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			resd.ApiOk(w, r, resp)
		}
	}
}
