package logic

import (
	"context"
	"go-zero-dandan/app/chat/ws/mgmodel"
	"go-zero-dandan/app/chat/ws/websocketd"
	"go-zero-dandan/pkg/numd"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/chat/ws/internal/svc"
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

func (l *Conversation) SingleChat(chat *websocketd.Chat, userId int64) error {
	if chat.ConversationCode == "" {
		chat.ConversationCode = numd.CombineInt64(userId, chat.RecvId)
	}
	logx.Info(chat)
	data := &mgmodel.ChatLog{
		ConversationCode: chat.ConversationCode,
		SendId:           userId,
		RecvId:           chat.RecvId,
		MsgFrom:          0,
		ChatType:         chat.ChatType,
		MsgType:          chat.MsgType,
		MsgContent:       chat.Content,
		SendTime:         time.Now().UnixNano(),
		State:            0,
		ReadRecords:      nil,
		UpdateAt:         time.Time{},
		CreateAt:         time.Time{},
	}
	return l.svc.ChatLogModel.Insert(l.ctx, data)
}
