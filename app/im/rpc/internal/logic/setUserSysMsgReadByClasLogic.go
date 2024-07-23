package logic

import (
	"context"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
)

type SetUserSysMsgReadByClasLogic struct {
	*SetUserSysMsgReadByClasLogicGen
}

func NewSetUserSysMsgReadByClasLogic(ctx context.Context, svc *svc.ServiceContext) *SetUserSysMsgReadByClasLogic {
	return &SetUserSysMsgReadByClasLogic{
		SetUserSysMsgReadByClasLogicGen: NewSetUserSysMsgReadByClasLogicGen(ctx, svc),
	}
}

func (l *SetUserSysMsgReadByClasLogic) SetUserSysMsgReadByClas(in *imRpc.SetUserSysMsgReadByClasReq) (*imRpc.ResultResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	//设置每条消息已读
	_, err := l.svc.SysMsgLogModel.SetSysMsgReadByMsgClas(l.ctx, l.meta.UserId, l.req.ClasList)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	//更新未读消息数
	if _, err := l.svc.SysMsgStatModel.SetZeroSysMsgUnreadNum(l.ctx, l.meta.UserId, l.req.ClasList); err != nil {
		return nil, l.resd.Error(err)
	}
	return &imRpc.ResultResp{}, nil
}
