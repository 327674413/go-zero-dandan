package logic

import (
	"context"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
)

type GetUserSysMsgUnreadNumLogic struct {
	*GetUserSysMsgUnreadNumLogicGen
}

func NewGetUserSysMsgUnreadNumLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserSysMsgUnreadNumLogic {
	return &GetUserSysMsgUnreadNumLogic{
		GetUserSysMsgUnreadNumLogicGen: NewGetUserSysMsgUnreadNumLogicGen(ctx, svc),
	}
}

func (l *GetUserSysMsgUnreadNumLogic) GetUserSysMsgUnreadNum(in *imRpc.GetUserSysMsgUnreadNumReq) (*imRpc.GetUserSysMsgUnreadNumResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	list, err := l.svc.SysMsgStatModel.GetUserUnread(l.ctx, l.req.UserId, l.req.MsgClasEm)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	unread := make(map[int64]int64)
	for _, v := range list {
		unread[int64(v.MsgClasEm)] = v.UnreadNum
	}
	return &imRpc.GetUserSysMsgUnreadNumResp{Unread: unread}, nil
}
