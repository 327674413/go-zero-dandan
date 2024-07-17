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
func (l *CreateFriendApplyLogic) CreateFriendApply(in *types.CreateFriendApplyReq) (resp *types.ResultResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	stamp := time.Now().Unix()
	_, err = l.svc.SocialRpc.CreateFriendApply(l.ctx, &socialRpc.CreateFriendApplyReq{
		PlatId:    &l.meta.PlatId,
		UserId:    &l.meta.UserId,
		FriendUid: &l.req.FriendUid,
		ApplyMsg:  &l.req.ApplyMsg,
		SourceEm:  &l.req.SourceEm,
		ApplyAt:   &stamp,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	return &types.ResultResp{Result: true}, nil
}
