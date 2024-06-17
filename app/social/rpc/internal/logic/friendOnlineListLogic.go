package logic

import (
	"context"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendOnlineListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendOnlineListLogic {
	return &FriendOnlineListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendOnlineListLogic) FriendOnlineList(in *pb.FriendListReq) (*pb.FriendOnlineResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	data, err := NewFriendListLogic(l.ctx, l.svcCtx).FriendList(&pb.FriendListReq{
		UserId: in.UserId,
		PlatId: in.PlatId,
	})
	if err != nil {
		return nil, err
	}
	uids := make([]string, 0, len(data.List))
	for _, v := range data.List {
		uids = append(uids, v.FriendId)
	}
	onlines, err := l.svcCtx.Redis.HgetallCtx(l.ctx, websocketd.RedisOnlineUser)
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
	return &pb.FriendOnlineResp{
		OnlineUser: onlineMap,
	}, nil
}
func (l *FriendOnlineListLogic) checkReqParams(in *pb.FriendListReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
