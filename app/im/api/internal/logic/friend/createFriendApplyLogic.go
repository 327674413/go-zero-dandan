package friend

import (
	"context"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"
	"time"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type CreateFriendApplyLogic struct {
	*CreateFriendApplyLogicGen
}

func NewCreateFriendApplyLogic(ctx context.Context, svc *svc.ServiceContext) *CreateFriendApplyLogic {
	return &CreateFriendApplyLogic{
		CreateFriendApplyLogicGen: NewCreateFriendApplyLogicGen(ctx, svc),
	}
}
func (l *CreateFriendApplyLogic) CreateFriendApply(req *types.CreateFriendApplyReq) (resp *types.ResultResp, err error) {
	if err = l.init(req); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	_, err = l.svc.SocialRpc.CreateFriendApply(l.ctx, &socialRpc.CreateFriendApplyReq{
		PlatId:    l.platId,
		UserId:    l.userMainInfo.Id,
		FriendUid: l.ReqFriendUid,
		ApplyMsg:  l.ReqApplyMsg,
		SourceEm:  l.ReqSourceEm,
		ApplyAt:   time.Now().Unix(),
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	return &types.ResultResp{Result: true}, nil
}
func (l *CreateFriendApplyLogic) init(req *types.CreateFriendApplyReq) (err error) {
	if err = l.initReq(req); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}
	return nil
}
