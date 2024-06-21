package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *pb.FriendListReq) (*pb.FriendListResp, error) {
	if in.UserId == "" {
		return nil, resd.NewRpcErrWithTempCtx(l.ctx, "缺少UserId", resd.ReqFieldRequired1, "*UserId")
	}
	m := model.NewSocialFriendModel(l.svcCtx.SqlConn, in.PlatId)

	list, err := m.Where("user_id = ?", in.UserId).Select()

	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	resp := &pb.FriendListResp{
		List: make([]*pb.Friends, 0),
	}
	if err = copier.Copy(&resp.List, list); err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	return resp, nil

}
