package conversation

import (
	"context"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/common/resd"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type SetUpMyConversationLogic struct {
	*SetUpMyConversationLogicGen
}

func NewSetUpMyConversationLogic(ctx context.Context, svc *svc.ServiceContext) *SetUpMyConversationLogic {
	return &SetUpMyConversationLogic{
		SetUpMyConversationLogicGen: NewSetUpMyConversationLogicGen(ctx, svc),
	}
}
func (l *SetUpMyConversationLogic) SetUpMyConversation(in *types.SetUpMyConversationReq) (resp *types.ResultResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	if l.req.ChatType != 1 && l.req.ChatType != 2 {
		return nil, l.resd.NewErrWithTemp(resd.ErrReqParamFormat1, "chatType")
	}
	_, err = l.svc.ImRpc.SetUpUserConversation(l.ctx, &imRpc.SetUpUserConversationReq{
		SendId:   &l.meta.UserId,
		RecvId:   &l.req.RecvId,
		ChatType: &l.req.ChatType,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	return
}
