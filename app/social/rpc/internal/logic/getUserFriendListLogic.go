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

type GetUserFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFriendListLogic {
	return &GetUserFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFriendListLogic) GetUserFriendList(in *pb.GetUserFriendListReq) (*pb.FriendListResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	if in.UserId == "" {
		return nil, resd.NewRpcErrWithTempCtx(l.ctx, "缺少UserId", resd.ReqFieldRequired1, "*UserId")
	}
	m := model.NewSocialFriendModel(l.svcCtx.SqlConn, in.PlatId)

	list, err := m.Where("user_id = ?", in.UserId).Select()

	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	resp := &pb.FriendListResp{
		List: make([]*pb.FriendInfo, 0),
	}
	if err = copier.Copy(&resp.List, list); err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	return resp, nil
	return &pb.FriendListResp{}, nil
}
func (l *GetUserFriendListLogic) checkReqParams(in *pb.GetUserFriendListReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
