package conversation

import (
	"context"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/common/utild/copier"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type GetChatLogLogic struct {
	*GetChatLogLogicGen
}

func NewGetChatLogLogic(ctx context.Context, svc *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		GetChatLogLogicGen: NewGetChatLogLogicGen(ctx, svc),
	}
}
func (l *GetChatLogLogic) GetChatLog(in *types.GetChatLogReq) (resp *types.GetChatLogResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	data, err := l.svc.ImRpc.GetChatLog(l.ctx, &imRpc.GetChatLogReq{
		ConversationId: &l.req.ConversationId,
		StartSendAt:    &l.req.StartSendAt,
		EndSendAt:      &l.req.EndSendAt,
		Count:          &l.req.Count,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}

	var res types.GetChatLogResp
	copier.Copy(&res, data)

	return &res, err
}
