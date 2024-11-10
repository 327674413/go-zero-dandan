package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dandan/app/dify/api/internal/logic"
	"go-zero-dandan/app/dify/api/internal/svc"
	"go-zero-dandan/app/dify/api/internal/types"
	"go-zero-dandan/common/resd"
	"net/http"
)

func chatMessagesStreamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatMessagesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}
		l := logic.NewChatMessagesStreamLogic(r.Context(), svcCtx)
		err := l.ChatMessagesStream(w, &req)
		if err != nil {
			resd.ApiFail(w, r, err)
		} else {
			httpx.Ok(w)
		}
	}
}
