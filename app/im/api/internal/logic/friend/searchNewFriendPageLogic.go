package friend

import (
	"context"
	"go-zero-dandan/app/user/rpc/types/pb"
	"strings"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type SearchNewFriendPageLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
}

func NewSearchNewFriendPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchNewFriendPageLogic {
	return &SearchNewFriendPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchNewFriendPageLogic) SearchNewFriendPage(req *types.SearchNewFriendReq) (resp *types.SearchNewFriendResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if req.Keyword == nil {
		return nil, resd.NewErrWithTempCtx(l.ctx, "keyword不得为空", resd.ReqFieldRequired1, "keyword")
	}
	keyword := strings.TrimSpace(*req.Keyword)
	if keyword == "" {
		return nil, resd.NewErrWithTempCtx(l.ctx, "keyword不得为空", resd.ReqFieldRequired1, "keyword")
	}
	searchRes, err := l.svcCtx.UserRpc.GetUserPage(l.ctx, &pb.GetUserPageReq{
		PlatId: l.platId,
		Match: map[string]*pb.MatchField{
			"phone": {Str: &keyword},
		},
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	list := make([]*types.NewFriendInfo, len(searchRes.List))
	ids := make([]string, 0)
	for _, v := range searchRes.List {
		list = append(list, &types.NewFriendInfo{
			Id:        v.Id,
			Nickname:  v.Nickname,
			AvatarImg: v.AvatarImg,
			Signature: v.Signature,
		})
		ids = append(ids, v.Id)
	}
	//relatsRes, err := l.svcCtx.SocialRpc.GetUserRelation(l.ctx)
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
