package conversation

import (
	"context"
	"go-zero-dandan/app/im/rpc/types/imRpc"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type GetConversationListLogic struct {
	*GetConversationListLogicGen
}

func NewGetConversationListLogic(ctx context.Context, svc *svc.ServiceContext) *GetConversationListLogic {
	return &GetConversationListLogic{
		GetConversationListLogicGen: NewGetConversationListLogicGen(ctx, svc),
	}
}
func (l *GetConversationListLogic) GetConversationList() (resp *types.GetConversationListResp, err error) {
	if err = l.initReq(); err != nil {
		return nil, l.resd.Error(err)
	}
	data, err := l.svc.ImRpc.GetConversations(l.ctx, &imRpc.GetConversationsReq{UserId: &l.meta.UserId})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	resp = &types.GetConversationListResp{
		Conversations: make(map[string]*types.Conversation),
	}
	for _, v := range data.ConversationList {
		conv := &types.Conversation{
			ConversationId: v.ConversationId,
			ChatType:       v.ChatType,
			TargetId:       v.TargetId,
			IsShow:         v.IsShow,
			ReadSeq:        v.ReadSeq,
			Unread:         v.Unread,
			Total:          v.Total,
			LastAt:         v.LastAt,
			DeleteSeq:      v.DeleteSeq,
		}
		if v.LastMsg != nil {
			conv.LastMsg = &types.ChatLog{
				Id:             v.LastMsg.Id,
				ConversationId: v.LastMsg.ConversationId,
				SendId:         v.LastMsg.SendId,
				RecvId:         v.LastMsg.RecvId,
				MsgType:        v.LastMsg.MsgType,
				MsgContent:     v.LastMsg.MsgContent,
				ChatType:       v.LastMsg.ChatType,
				SendTime:       v.LastMsg.SendTime,
			}
		}
		resp.Conversations[v.ConversationId] = conv
	}
	return
}
