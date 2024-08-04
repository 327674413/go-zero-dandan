package conversation

import (
	"context"

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
	return
}
