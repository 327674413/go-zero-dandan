package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupListLogic {
	return &GetUserGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserGroupListLogic) GetUserGroupList(in *pb.GetUserGroupListReq) (*pb.GroupListResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &pb.GroupListResp{}, nil
}
func (l *GetUserGroupListLogic) checkReqParams(in *pb.GetUserGroupListReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
