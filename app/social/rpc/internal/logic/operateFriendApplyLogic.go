package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateFriendApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperateFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateFriendApplyLogic {
	return &OperateFriendApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OperateFriendApplyLogic) OperateFriendApply(in *socialRpc.OperateFriendApplyReq) (*socialRpc.ResultResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &socialRpc.ResultResp{}, nil
}
func (l *OperateFriendApplyLogic) checkReqParams(in *socialRpc.OperateFriendApplyReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
