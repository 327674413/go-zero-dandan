package logic

import (
	"context"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
)

type GetGroupUserOnlineLogic struct {
	*GetGroupUserOnlineLogicGen
}

func NewGetGroupUserOnlineLogic(ctx context.Context, svc *svc.ServiceContext) *GetGroupUserOnlineLogic {
	return &GetGroupUserOnlineLogic{
		GetGroupUserOnlineLogicGen: NewGetGroupUserOnlineLogicGen(ctx, svc),
	}
}

func (l *GetGroupUserOnlineLogic) GetGroupUserOnline(in *socialRpc.GetGroupUserOnlineReq) (*socialRpc.GroupUserOnlineResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	data, err := NewGetGroupMemberListLogic(l.ctx, l.svc).GetGroupMemberList(&socialRpc.GetGroupMemberListReq{
		GroupId: in.GroupId,
		PlatId:  in.PlatId,
	})
	if err != nil {
		return nil, err
	}
	uids := make([]string, 0, len(data.List))
	for _, v := range data.List {
		uids = append(uids, v.UserId)
	}
	onlines, err := l.svc.Redis.HgetallCtx(l.ctx, websocketd.RedisOnlineUser)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	onlineMap := make(map[string]bool, len(uids))
	for _, uid := range uids {
		if _, ok := onlines[uid]; ok {
			onlineMap[uid] = true
		} else {
			onlineMap[uid] = false
		}
	}
	return &socialRpc.GroupUserOnlineResp{
		OnlineUser: onlineMap,
	}, nil
}
func (l *GetGroupUserOnlineLogic) checkReqParams(in *socialRpc.GetGroupUserOnlineReq) error {
	return nil
}
