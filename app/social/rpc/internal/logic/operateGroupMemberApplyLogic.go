package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateGroupMemberApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperateGroupMemberApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateGroupMemberApplyLogic {
	return &OperateGroupMemberApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OperateGroupMemberApplyLogic) OperateGroupMemberApply(in *socialRpc.OperateGroupMemberApplyReq) (*socialRpc.ResultResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &socialRpc.ResultResp{}, nil
}
func (l *OperateGroupMemberApplyLogic) checkReqParams(in *socialRpc.OperateGroupMemberApplyReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
