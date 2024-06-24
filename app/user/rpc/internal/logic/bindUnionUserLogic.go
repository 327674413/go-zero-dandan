package logic

import (
	"context"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
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
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &pb.BindUnionUserResp{}, nil
}
func (l *BindUnionUserLogic) checkReqParams(in *pb.BindUnionUserReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
