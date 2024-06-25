package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFriendApplyPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFriendApplyPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFriendApplyPageLogic {
	return &GetUserFriendApplyPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFriendApplyPageLogic) GetUserFriendApplyPage(in *pb.GetUserFriendApplyPageReq) (*pb.FriendApplyPageResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &pb.FriendApplyPageResp{}, nil
}
func (l *GetUserFriendApplyPageLogic) checkReqParams(in *pb.GetUserFriendApplyPageReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
