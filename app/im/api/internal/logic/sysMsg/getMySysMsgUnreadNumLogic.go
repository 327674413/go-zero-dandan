package sysMsg

import (
	"context"
	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
	"go-zero-dandan/app/im/rpc/types/imRpc"
)

type GetMySysMsgUnreadNumLogic struct {
	*GetMySysMsgUnreadNumLogicGen
}

func NewGetMySysMsgUnreadNumLogic(ctx context.Context, svc *svc.ServiceContext) *GetMySysMsgUnreadNumLogic {
	return &GetMySysMsgUnreadNumLogic{
		GetMySysMsgUnreadNumLogicGen: NewGetMySysMsgUnreadNumLogicGen(ctx, svc),
	}
}
func (l *GetMySysMsgUnreadNumLogic) GetMySysMsgUnreadNum(in *types.GetMySysMsgUnreadNumReq) (resp *types.GetMySysMsgUnreadNumResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	res, err := l.svc.ImRpc.GetUserSysMsgUnreadNum(l.ctx, &imRpc.GetUserSysMsgUnreadNumReq{
		UserId:    &l.meta.UserId,
		MsgClasEm: in.MsgClasEm,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	return &types.GetMySysMsgUnreadNumResp{Unread: res.Unread}, nil
}
