package logic

import (
	"context"

	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinListLogic {
	return &GroupPutinListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutinListLogic) GroupPutinList(in *pb.GroupPutinListReq) (*pb.GroupPutinListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GroupPutinListResp{}, nil
}
