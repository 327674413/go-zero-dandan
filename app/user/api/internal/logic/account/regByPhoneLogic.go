package account

import (
	"context"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
)

type RegByPhoneLogic struct {
	*RegByPhoneLogicGen
}

func NewRegByPhoneLogic(ctx context.Context, svc *svc.ServiceContext) *RegByPhoneLogic {
	return &RegByPhoneLogic{
		RegByPhoneLogicGen: NewRegByPhoneLogicGen(ctx, svc),
	}
}

func (l *RegByPhoneLogic) RegByPhone(in *types.RegByPhoneReq) (resp *types.UserInfoResp, err error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	return
}
