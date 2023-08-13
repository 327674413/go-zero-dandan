package logic

import (
	"context"

	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"

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

func (l *GetUserFriendListLogic) GetUserFriendList(in *pb.IdReq) (*pb.GetUserFriendList, error) {

	return &pb.GetUserFriendList{List: []*pb.UserRelationInfo{}}, nil
}
