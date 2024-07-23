package logic

import (
	"context"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/utild"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/ws/internal/svc"
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
		MsgFrom:        0,
		ChatType:       chat.ChatType,
		MsgType:        chat.MsgType,
		MsgContent:     chat.Content,
		SendTime:       utild.NowTime(),
		State:          0,
		ReadRecords:    nil,
		UpdateAt:       time.Time{},
		CreateAt:       time.Time{},
	}
	return l.svc.ChatLogModel.Insert(l.ctx, data)
}
