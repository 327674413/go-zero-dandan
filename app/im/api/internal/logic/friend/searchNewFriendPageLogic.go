package friend

import (
	"context"
	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/app/user/rpc/types/userRpc"

	"go-zero-dandan/common/resd"
)

type SearchNewFriendPageLogic struct {
	*SearchNewFriendPageLogicGen
}

func NewSearchNewFriendPageLogic(ctx context.Context, svc *svc.ServiceContext) *SearchNewFriendPageLogic {
	return &SearchNewFriendPageLogic{
		SearchNewFriendPageLogicGen: NewSearchNewFriendPageLogicGen(ctx, svc),
	}
}

func (l *SearchNewFriendPageLogic) SearchNewFriendPage(in *types.SearchNewFriendReq) (resp *types.SearchNewFriendResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	searchRes, err := l.svc.UserRpc.GetUserPage(l.ctx, &userRpc.GetUserPageReq{
		PlatId: &l.meta.PlatId,
		Match: map[string]*userRpc.MatchField{
			"phone": {Str: &l.req.Keyword},
		},
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	list := make([]*types.NewFriendInfo, len(searchRes.List))
	ids := make([]string, 0)
	for k, v := range searchRes.List {
		list[k] = &types.NewFriendInfo{
			Id:        v.Id,
			Nickname:  v.Nickname,
			AvatarImg: v.AvatarImg,
			Signature: v.Signature,
		}
		ids = append(ids, v.Id)
	}
	if len(list) > 0 {
		relatsRes, err := l.svc.SocialRpc.GetUserRelation(l.ctx, &socialRpc.GetUserRelationReq{
			UserId:     &l.meta.UserId,
			FriendUids: ids,
			PlatId:     &l.meta.PlatId,
		})
		if err != nil {
			return nil, resd.ErrorCtx(l.ctx, err)
		}
		for i, v := range list {
			if state, ok := relatsRes.Relats[v.Id]; ok {
				list[i].StateEm = state
			}
		}
	}

	return &types.SearchNewFriendResp{
		List: list,
	}, nil
}
