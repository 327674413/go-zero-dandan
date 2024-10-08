package friend

import (
	"context"
	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"
)

type GetMyFriendApplyRecvPageLogic struct {
	*GetMyFriendApplyRecvPageLogicGen
}

func NewGetMyFriendApplyRecvPageLogic(ctx context.Context, svc *svc.ServiceContext) *GetMyFriendApplyRecvPageLogic {
	return &GetMyFriendApplyRecvPageLogic{
		GetMyFriendApplyRecvPageLogicGen: NewGetMyFriendApplyRecvPageLogicGen(ctx, svc),
	}
}

func (l *GetMyFriendApplyRecvPageLogic) GetMyFriendApplyRecvPage(in *types.GetMyFriendApplyRecvPageReq) (resp *types.FriendApplyListResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, err
	}
	rpcReq := &socialRpc.GetUserRecvFriendApplyPageReq{
		UserId: &l.meta.UserId,
		PlatId: &l.meta.PlatId,
	}
	if !l.hasReq.Page {
		rpcReq.Page = &l.req.Page
	}
	if !l.hasReq.Size {
		rpcReq.Size = &l.req.Size
	}
	res, err := l.svc.SocialRpc.GetUserRecvFriendApplyPage(l.ctx, rpcReq)
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
