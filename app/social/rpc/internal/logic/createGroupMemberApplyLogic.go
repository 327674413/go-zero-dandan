package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupMemberApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupMemberApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupMemberApplyLogic {
	return &CreateGroupMemberApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupMemberApplyLogic) CreateGroupMemberApply(in *socialRpc.CreateGroupMemberApplyReq) (*socialRpc.CreateGroupMemberApplyResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &socialRpc.CreateGroupMemberApplyResp{}, nil
}
func (l *CreateGroupMemberApplyLogic) checkReqParams(in *socialRpc.CreateGroupMemberApplyReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}