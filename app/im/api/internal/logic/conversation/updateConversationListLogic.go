package conversation

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type UpdateConversationListLogic struct {
	*UpdateConversationListLogicGen
}

func NewUpdateConversationListLogic(ctx context.Context, svc *svc.ServiceContext) *UpdateConversationListLogic {
	return &UpdateConversationListLogic{
		UpdateConversationListLogicGen: NewUpdateConversationListLogicGen(ctx, svc),
	}
}
func (l *UpdateConversationListLogic) UpdateConversationList(in *types.UpdateConversationListReq) (resp *types.ResultResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	return
}
