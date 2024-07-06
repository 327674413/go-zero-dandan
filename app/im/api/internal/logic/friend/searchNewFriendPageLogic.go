package friend

import (
	"context"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"strings"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type SearchNewFriendPageLogic struct {
	*SearchNewFriendPageLogicGen
}

func NewSearchNewFriendPageLogic(ctx context.Context, svc *svc.ServiceContext) *SearchNewFriendPageLogic {
	return &SearchNewFriendPageLogic{
		SearchNewFriendPageLogicGen: NewSearchNewFriendPageLogicGen(ctx, svc),
	}
}

func (l *SearchNewFriendPageLogic) SearchNewFriendPage(req *types.SearchNewFriendReq) (resp *types.SearchNewFriendResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if req.Keyword == nil {
		return nil, resd.NewErrWithTempCtx(l.ctx, "keyword不得为空", resd.ReqFieldRequired1, "keyword")
	}
	keyword := strings.TrimSpace(*req.Keyword)
	if keyword == "" {
		return nil, resd.NewErrWithTempCtx(l.ctx, "keyword不得为空", resd.ReqFieldRequired1, "keyword")
	}
	searchRes, err := l.svc.UserRpc.GetUserPage(l.ctx, &userRpc.GetUserPageReq{
		PlatId: l.platId,
		Match: map[string]*userRpc.MatchField{
			"phone": {Str: &keyword},
		},
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
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
			UserId:     l.userMainInfo.Id,
			FriendUids: ids,
			PlatId:     l.platId,
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

func (l *SearchNewFriendPageLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *SearchNewFriendPageLogic) initPlat() (err error) {
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
