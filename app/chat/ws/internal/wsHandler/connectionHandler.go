package wsHandler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/chat/ws/internal/wsLogic"
	"go-zero-dandan/app/chat/ws/internal/wsSvc"
	"go-zero-dandan/app/chat/ws/internal/wsTypes"
	"go-zero-dandan/common/resd"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func connectionHandler(svcCtx *wsSvc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req wsTypes.ConnectionReq
		if err := httpx.Parse(r, &req); err != nil {
			// todo::这里抄的还没改过
			resd.ErrorCtx(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 获取logic独享，单例模式
		hub := wsLogic.InitHub(r.Context(), svcCtx)
		// 处理连接请求
		res, err := hub.Connection(&req)
		status := http.StatusUnauthorized
		if err == nil {
			// 请求通过，升级ws协议
			err := hub.WsUpgrade(res.UserId, &req, w, r, nil)
			if err != nil {
				logx.WithContext(r.Context()).Errorf("ws.WsUpgrade error: %s", err)
				return
			}
		} else {
			// todo::这里抄的还没改过
			w.Header().Set("Sec-Websocket-Version", "13")
			w.Header().Set("ws_err_msg", "args err, need token, sendID, platformID")
			http.Error(w, http.StatusText(status), status)
		}
	}
}
