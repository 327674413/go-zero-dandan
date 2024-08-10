package conversation

import (
	"context"
	"go-zero-dandan/app/im/rpc/types/imRpc"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type DeleteConversationLogic struct {
	*DeleteConversationLogicGen
}

func NewDeleteConversationLogic(ctx context.Context, svc *svc.ServiceContext) *DeleteConversationLogic {
	return &DeleteConversationLogic{
		DeleteConversationLogicGen: NewDeleteConversationLogicGen(ctx, svc),
	}
}
func (l *DeleteConversationLogic) DeleteConversation(in *types.IdReq) (resp *types.ResultResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	_, err = l.svc.ImRpc.DeleteUserConversation(l.ctx, &imRpc.DeleteUserConversationReq{
		UserId:         &l.meta.UserId,
		ConversationId: &l.req.Id,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	return
}
