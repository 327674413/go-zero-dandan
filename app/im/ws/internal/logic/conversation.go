package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/utild"
)

type Conversation struct {
	ctx    context.Context
	svc    *svc.ServiceContext
	server *websocketd.Server
	logx.Logger
}

func NewConversationLogic(ctx context.Context, server *websocketd.Server, svc *svc.ServiceContext) *Conversation {
	return &Conversation{
		ctx:    ctx,
		svc:    svc,
		server: server,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Conversation) SingleChat(chat *websocketd.Chat, userId string) error {
	if chat.ConversationId == "" {
		chat.ConversationId = utild.CombineId(userId, chat.RecvId)
	}
	data := &modelMongo.ChatLog{
		ConversationId: chat.ConversationId,
		SendId:         userId,
		RecvId:         chat.RecvId,
		ChatType:       chat.ChatType,
		MsgType:        chat.MsgType,
		MsgContent:     chat.MsgContent,
		SendTime:       utild.NowTime(),
		SendAtMs:       utild.GetTimeMs(),
		MsgState:       0,
		MsgReads:       nil,
	}
	return l.svc.ChatLogModel.Insert(l.ctx, data)
}
