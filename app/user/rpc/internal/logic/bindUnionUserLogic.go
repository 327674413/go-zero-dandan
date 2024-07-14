package logic

import (
	"context"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
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

func (l *BindUnionUserLogic) BindUnionUser(in *userRpc.BindUnionUserReq) (*userRpc.BindUnionUserResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &userRpc.BindUnionUserResp{}, nil
}
func (l *BindUnionUserLogic) checkReqParams(in *userRpc.BindUnionUserReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ErrReqFieldRequired1, "platId")
	}
	return nil
}
