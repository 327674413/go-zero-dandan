package logic

import (
	"context"

	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	BindClasEmPhone   = 1
	BindClasEmAccount = 2
)

type BindUnionUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindUnionUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUnionUserLogic {
	return &BindUnionUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindUnionUserLogic) BindUnionUser(in *pb.BindUnionUserReq) (*pb.BindUnionUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.BindUnionUserResp{}, nil
}
