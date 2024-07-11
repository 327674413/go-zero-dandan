package logic

import (
	"context"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"
)

type GetFriendOnlineLogic struct {
	*GetFriendOnlineLogicGen
}

func NewGetFriendOnlineLogic(ctx context.Context, svc *svc.ServiceContext) *GetFriendOnlineLogic {
	return &GetFriendOnlineLogic{
		GetFriendOnlineLogicGen: NewGetFriendOnlineLogicGen(ctx, svc),
	}
}

func (l *GetFriendOnlineLogic) GetFriendOnline(in *socialRpc.GetFriendOnlineReq) (*socialRpc.FriendOnlineResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	data, err := NewGetUserFriendListLogic(l.ctx, l.svc).GetUserFriendList(&socialRpc.GetUserFriendListReq{
		UserId: in.UserId,
		PlatId: in.PlatId,
	})
	if err != nil {
		return nil, err
	}
	uids := make([]string, 0, len(data.List))
	for _, v := range data.List {
		uids = append(uids, v.FriendUid)
	}
	onlines, err := l.svc.Redis.HgetallCtx(l.ctx, websocketd.RedisOnlineUser)
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	onlineMap := make(map[string]bool, len(uids))
	for _, uid := range uids {
		if _, ok := onlines[uid]; ok {
			onlineMap[uid] = true
		} else {
			onlineMap[uid] = false
		}
	}
	return &socialRpc.FriendOnlineResp{
		OnlineUser: onlineMap,
	}, nil

}
func (l *GetFriendOnlineLogic) checkReqParams(in *socialRpc.GetFriendOnlineReq) error {

	return nil
}
