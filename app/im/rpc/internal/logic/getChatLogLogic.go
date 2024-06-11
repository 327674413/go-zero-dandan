package logic

import (
	"context"
	"go-zero-dandan/app/im/rpc/im"
	"go-zero-dandan/common/resd"

	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetChatLog 获取会话记录
func (l *GetChatLogLogic) GetChatLog(in *pb.GetChatLogReq) (*pb.GetChatLogResp, error) {
	// 根据id
	if in.MsgId != "" {
		chatLog, err := l.svcCtx.ChatLogModel.FindOne(l.ctx, in.MsgId)
		if err != nil {
			return nil, resd.ErrorCtx(l.ctx, err)
		}
		return &pb.GetChatLogResp{
			List: []*pb.ChatLog{
				{
					Id:             chatLog.ID.Hex(),
					SendTime:       chatLog.SendTime,
					MsgType:        int64(chatLog.MsgType),
					ChatType:       int64(chatLog.ChatType),
					ConversationId: chatLog.ConversationId,
					RecvId:         chatLog.RecvId,
					SendId:         chatLog.SendId,
					MsgContent:     chatLog.MsgContent,
					ReadRecords:    chatLog.ReadRecords,
				},
			},
		}, nil
	}
	// 根据时间段
	data, err := l.svcCtx.ChatLogModel.ListBySendTime(l.ctx, in.ConversationId, in.StartSendTime, in.EndSendTime, in.Count)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	list := make([]*im.ChatLog, 0, len(data))
	for _, item := range data {
		list = append(list, &im.ChatLog{
			Id:             item.ID.Hex(),
			SendTime:       item.SendTime,
			MsgType:        int64(item.MsgType),
			ChatType:       int64(item.ChatType),
			ConversationId: item.ConversationId,
			RecvId:         item.RecvId,
			SendId:         item.SendId,
			MsgContent:     item.MsgContent,
			ReadRecords:    item.ReadRecords,
		})
	}
	return &pb.GetChatLogResp{
		List: list,
	}, nil
}
