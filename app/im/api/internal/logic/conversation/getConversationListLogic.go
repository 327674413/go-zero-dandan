package conversation

import (
	"context"
	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"
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
	convMap := make(map[string]*types.Conversation)
	if err = copier.Copy(&convMap, data.ConversationList); err != nil {
		return nil, l.resd.Error(err, resd.ErrCopier)
	}
	resp.Conversations = convMap
	return resp, nil

}
