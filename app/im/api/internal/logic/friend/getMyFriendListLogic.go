package friend

import (
	"context"
	"go-zero-dandan/app/social/rpc/social"
	"go-zero-dandan/common/utild/copier"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type GetMyFriendListLogic struct {
	*GetMyFriendListLogicGen
}

func NewGetMyFriendListLogic(ctx context.Context, svc *svc.ServiceContext) *GetMyFriendListLogic {
	return &GetMyFriendListLogic{
		GetMyFriendListLogicGen: NewGetMyFriendListLogicGen(ctx, svc),
	}
}

func (l *GetMyFriendListLogic) GetMyFriendList() (resp *types.FriendListResp, err error) {
	if err = l.initReq(); err != nil {
		return nil, l.resd.Error(err)
	}
	friends, err := l.svc.SocialRpc.GetUserFriendList(l.ctx, &social.GetUserFriendListReq{
		UserId: &l.meta.UserId,
		PlatId: &l.meta.PlatId,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	list := make([]*types.FriendInfo, 0)
	copier.Copy(&list, friends.List)
	return &types.FriendListResp{List: list}, nil
}
