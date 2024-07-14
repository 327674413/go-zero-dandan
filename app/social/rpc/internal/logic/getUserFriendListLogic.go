package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"
)

type GetUserFriendListLogic struct {
	*GetUserFriendListLogicGen
}

func NewGetUserFriendListLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserFriendListLogic {
	return &GetUserFriendListLogic{
		GetUserFriendListLogicGen: NewGetUserFriendListLogicGen(ctx, svc),
	}
}

func (l *GetUserFriendListLogic) GetUserFriendList(in *socialRpc.GetUserFriendListReq) (*socialRpc.FriendListResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	m := model.NewSocialFriendModel(l.ctx, l.svc.SqlConn, l.req.PlatId)

	list, err := m.Where("user_id = ?", in.UserId).Select()

	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	resp := &socialRpc.FriendListResp{
		List: make([]*socialRpc.FriendInfo, 0),
	}
	if err = copier.Copy(&resp.List, list); err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	return resp, nil
	return &socialRpc.FriendListResp{}, nil
}
func (l *GetUserFriendListLogic) checkReqParams(in *socialRpc.GetUserFriendListReq) error {
	return nil
}
