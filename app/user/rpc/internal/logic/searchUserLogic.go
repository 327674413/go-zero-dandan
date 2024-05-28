package logic

import (
	"context"

	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLogic) SearchUser(in *pb.SearchUserReq) (*pb.SearchUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserResp{}, nil
}
