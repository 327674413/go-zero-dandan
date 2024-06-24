package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupMemberApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserGroupMemberApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupMemberApplyListLogic {
	return &GetUserGroupMemberApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserGroupMemberApplyListLogic) GetUserGroupMemberApplyList(in *pb.GetUserGroupMemberApplyListReq) (*pb.GroupMemberApplyListResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &pb.GroupMemberApplyListResp{}, nil
}
func (l *GetUserGroupMemberApplyListLogic) checkReqParams(in *pb.GetUserGroupMemberApplyListReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
