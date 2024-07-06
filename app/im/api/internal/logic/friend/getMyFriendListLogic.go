package friend

import (
	"context"
	"go-zero-dandan/app/social/rpc/social"
	"go-zero-dandan/common/utild/copier"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
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
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	friends, err := l.svc.SocialRpc.GetUserFriendList(l.ctx, &social.GetUserFriendListReq{
		UserId: l.userMainInfo.Id,
		PlatId: l.platId,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.FriendInfo, 0)
	copier.Copy(&list, friends.List)
	return &types.FriendListResp{List: list}, nil
}

func (l *GetMyFriendListLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *GetMyFriendListLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platId, _ := l.ctx.Value("platId").(string)
	if platId == "" {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platId
	l.platClasEm = platClasEm
	return nil
}
