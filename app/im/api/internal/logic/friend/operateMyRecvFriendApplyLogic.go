package friend

import (
	"context"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/pkg/arrd"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type OperateMyRecvFriendApplyLogic struct {
	*OperateMyRecvFriendApplyLogicGen
}

func NewOperateMyRecvFriendApplyLogic(ctx context.Context, svc *svc.ServiceContext) *OperateMyRecvFriendApplyLogic {
	return &OperateMyRecvFriendApplyLogic{
		OperateMyRecvFriendApplyLogicGen: NewOperateMyRecvFriendApplyLogicGen(ctx, svc),
	}
}
func (l *OperateMyRecvFriendApplyLogic) OperateMyRecvFriendApply(req *types.OperateMyRecvFriendApplyReq) (resp *types.ResultResp, err error) {
	if err = l.init(req); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	//合法性校验
	if !arrd.Contain([]int64{constd.SocialFriendStateEmPass, constd.SocialFriendStateEmReject}, l.ReqOperateStateEm) {
		return nil, resd.NewErrWithTempCtx(l.ctx, "", resd.ReqParamFormatErr1, "stateEm")
	}
	_, err = l.svc.SocialRpc.OperateFriendApply(l.ctx, &socialRpc.OperateFriendApplyReq{
		ApplyId:        l.ReqApplyId,
		OperateStateEm: l.ReqOperateStateEm,
		PlatId:         l.platId,
		OperateUid:     l.userMainInfo.Id,
		OperateMsg:     l.ReqOperateMsg,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	return &types.ResultResp{Result: true}, nil

}
func (l *OperateMyRecvFriendApplyLogic) init(req *types.OperateMyRecvFriendApplyReq) (err error) {
	if err = l.initReq(req); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}
	return nil
}
