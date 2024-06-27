package friend

import (
	"context"
	"go-zero-dandan/app/social/rpc/types/socialRpc"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type GetMyFriendApplyRecvPageLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
}

func NewGetMyFriendApplyRecvPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFriendApplyRecvPageLogic {
	return &GetMyFriendApplyRecvPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyFriendApplyRecvPageLogic) GetMyFriendApplyRecvPage(req *types.GetMyFriendApplyRecvPageReq) (resp *types.FriendApplyListResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	rpcReq := &socialRpc.GetUserRecvFriendApplyPageReq{
		UserId: l.userMainInfo.Id,
		PlatId: l.platId,
	}
	if req.Page != nil {
		rpcReq.Page = *req.Page
	}
	if req.Size != nil {
		rpcReq.Size = *req.Size
	}
	res, err := l.svcCtx.SocialRpc.GetUserRecvFriendApplyPage(l.ctx, rpcReq)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	resp = &types.FriendApplyListResp{List: make([]*types.FriendApply, len(res.List))}
	for k, v := range res.List {
		resp.List[k] = &types.FriendApply{
			Id:            v.Id,
			UserId:        v.UserId,
			FriendUid:     v.FriendUid,
			ApplyLastMsg:  v.ApplyLastMsg,
			ApplyLastAt:   v.ApplyLastAt,
			OperateMsg:    v.OperateMsg,
			OperateAt:     v.OperateAt,
			StateEm:       v.StateEm,
			UserName:      v.UserName,
			UserSex:       v.UserSex,
			UserAvatarImg: v.UserAvatarImg,
		}
	}
	return
}

func (l *GetMyFriendApplyRecvPageLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *GetMyFriendApplyRecvPageLogic) initPlat() (err error) {
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
