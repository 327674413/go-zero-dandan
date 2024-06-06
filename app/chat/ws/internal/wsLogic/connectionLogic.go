package wsLogic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/chat/ws/internal/wsSvc"
	"go-zero-dandan/app/chat/ws/internal/wsTypes"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/resd"
	"net/http"
	"sync"
	"time"
)

type UserConn struct {
	*websocket.Conn
	w *sync.Mutex
}

var once sync.Once

type Hub struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *wsSvc.ServiceContext
	wsMaxConnNum int
	wsUpGrader   *websocket.Upgrader
	wsConnToUser map[*UserConn]map[int64]int64
	wsUserToConn map[int64]map[int64]*UserConn
	//rep          *wsrepository.Rep
}

var (
	userCount       uint64
	rwLock          *sync.RWMutex
	sendMsgAllCount uint64
)

func init() {
	rwLock = new(sync.RWMutex)
}

var connectionLogic *Hub

func setHub(wsConn *Hub) {
	connectionLogic = wsConn
}
func GetHub() *Hub {
	return connectionLogic
}
func InitHub(ctx context.Context, svcCtx *wsSvc.ServiceContext) *Hub {
	//实现单例模式
	if connectionLogic != nil {
		return connectionLogic
	}
	ws := &Hub{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//rep:    wsrepository.NewRep(svcCtx),
	}
	ws.wsMaxConnNum = ws.svcCtx.Config.WebsocketConfig.MaxConnNum
	ws.wsConnToUser = make(map[*UserConn]map[int64]int64)
	ws.wsUserToConn = make(map[int64]map[int64]*UserConn)
	ws.wsUpGrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(ws.svcCtx.Config.WebsocketConfig.TimeOut) * time.Second,
		ReadBufferSize:   ws.svcCtx.Config.WebsocketConfig.MaxMsgLen,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	once.Do(func() {
		connectionLogic = ws
		//setHub(ws)
	})

	return connectionLogic
}
func (l *Hub) Connection(req *wsTypes.ConnectionReq) (*wsTypes.ConnectionRes, error) {
	//除url外，还可以通过Subprotocol: []string{"chat"}来接收
	//前端用 let sock = new WebSocket('ws://example.com', ['chat', $token])
	if req.PlatformEm == 0 {
		return nil, resd.NewErrCtx(l.ctx, "平台标识未传")
	}
	// 这里不能用l.ctx，会报错：rpc error: code = Canceled desc = context canceled，可能是单例复用的缘故
	user, err := l.svcCtx.UserRpc.GetUserByToken(context.Background(), &pb.TokenReq{
		Token: req.Token,
	})
	if err != nil {
		//conn.WriteMessage(websocket.CloseMessage, []byte("登录验证失败")) //CloseMessage这个会进onerror
		//conn.Close()                                                //这个直接断开连接，前端触发onclose
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if existClient := l.GetUserConn(user.Id, req.PlatformEm); existClient != nil {
		existClient.Conn.WriteMessage(websocket.TextMessage, []byte("您的账号在其他设备登录"))
		existClient.Conn.Close()
	}

	return &wsTypes.ConnectionRes{UserId: user.Id}, nil
}
func (l *Hub) WsUpgrade(userId int64, req *wsTypes.ConnectionReq, w http.ResponseWriter, r *http.Request, header http.Header) error {
	conn, err := l.wsUpGrader.Upgrade(w, r, header)
	if err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}
	newConn := &UserConn{conn, new(sync.Mutex)}
	userCount++
	l.addUserConn(userId, req.PlatformEm, newConn)
	go l.readMsg(newConn, userId, req.PlatformEm)
	return nil
}

func (l *Hub) readMsg(conn *UserConn, userId int64, platformEm int64) {
	for {
		messageType, msg, err := conn.ReadMessage()
		//todo::这里读取消息，判断消息内容，然后发送消息
		s := &wsTypes.Message{}
		if messageType == websocket.PingMessage {
			l.sendMsg(l.ctx, conn, "Pong")
		} else if messageType == websocket.TextMessage {
			json.Unmarshal(msg, &s)
			if s.TypeEm == wsTypes.MsgTypeEmChat {
				l.ChatTextHandler(s, conn)
			}
		}
		if err != nil {
			//uid, platform := l.getUserUid(conn)
			//logx.Error("WS ReadMsg error ", " userIP ", conn.RemoteAddr().String(), " userUid ", uid, " platform ", platform, " error ", err.Error())
			userCount--
			l.delUserConn(conn)
			return
		}
		/*// xtrace
		xtrace.RunWithTrace("", func(ctx context.Context) {
			l.msgParse(ctx, conn, msg, uid, platform)
		}, attribute.KeyValue{
			Key:   "uid",
			Value: attribute.StringValue(uid),
		}, attribute.KeyValue{
			Key:   "platform",
			Value: attribute.StringValue(platform),
		})*/
	}
}
func (l *Hub) ChatTextHandler(message *wsTypes.Message, fromUserConn *UserConn) {
	switch message.TargetTypeEm {
	case wsTypes.TargetTypeEmCrony:
		if userConns, ok := l.wsUserToConn[message.TargetId]; ok {
			//在线消息
			sendByte, err := json.Marshal(message)
			if err != nil {
				logc.Error(l.ctx, err)
				return
			}
			sendState := false
			for platform, v := range userConns {
				err = l.SendMsgToUser(l.ctx, v, sendByte, platform)
				if err == nil {
					sendState = true
				}
			}
			if sendState {
				l.responseChatSucc(message, fromUserConn)
			}

		} else {
			fmt.Println("进来了3")
			//离线消息
		}

	}
	//获取目标对象的conn

}

// responseChatSucc 回复发送状态
func (l *Hub) responseChatSucc(message *wsTypes.Message, conn *UserConn) error {
	byte, _ := json.Marshal(&wsTypes.MessageResp{
		Code:    message.Code,
		StateEm: wsTypes.ChatMsgStateEmSent,
		ErrCode: 0,
		TypeEm:  wsTypes.MsgTypeEmChatResp,
	})
	fmt.Println("触发发送状态回复")
	l.writeMsg(conn, websocket.TextMessage, byte)
	return nil
}
func (l *Hub) SendMsgToUser(ctx context.Context, conn *UserConn, bMsg []byte, platformEm int64) error {
	fmt.Println("SendMsgToUser start")
	err := l.writeMsg(conn, websocket.TextMessage, bMsg)
	if err != nil {
		return resd.NewErrCtx(l.ctx, err.Error())
	} else {
		return nil
	}
}
