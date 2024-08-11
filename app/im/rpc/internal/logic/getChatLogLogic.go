package logic

import (
	"context"
	"go-zero-dandan/app/im/rpc/im"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
)

type GetChatLogLogic struct {
	*GetChatLogLogicGen
}

func NewGetChatLogLogic(ctx context.Context, svc *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		GetChatLogLogicGen: NewGetChatLogLogicGen(ctx, svc),
	}
}

// GetChatLog 获取会话记录
func (l *GetChatLogLogic) GetChatLog(in *imRpc.GetChatLogReq) (*imRpc.GetChatLogResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	// 根据id
	if l.req.MsgId != "" {
		chatLog, err := l.svc.ChatLogModel.FindOne(l.ctx, l.req.MsgId)
		if err != nil {
			return nil, l.resd.Error(err)
		}
		return &imRpc.GetChatLogResp{
			List: []*imRpc.ChatLog{
				{
					Id:             chatLog.ID.Hex(),
					SendTime:       chatLog.SendTime,
					SendAtMs:       chatLog.SendAtMs,
					MsgState:       chatLog.MsgState,
					MsgType:        int64(chatLog.MsgType),
					ChatType:       int64(chatLog.ChatType),
					ConversationId: chatLog.ConversationId,
					RecvId:         chatLog.RecvId,
					SendId:         chatLog.SendId,
					MsgContent:     chatLog.MsgContent,
					MsgReads:       chatLog.MsgReads,
				},
			},
		}, nil
	}
	// 根据时间段
	data, err := l.svc.ChatLogModel.ListBySendTime(l.ctx, l.req.ConversationId, l.req.StartSendAt, l.req.EndSendAt, l.req.Count)
	if err != nil {
		return nil, l.resd.Error(err)
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
			MsgReads:       item.MsgReads,
			SendAtMs:       item.SendAtMs,
			MsgState:       item.MsgState,
		})
	}
	return &imRpc.GetChatLogResp{
		List: list,
	}, nil
}
