package logic

import (
	"context"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupUserOnlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupUserOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupUserOnlineLogic {
	return &GetGroupUserOnlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupUserOnlineLogic) GetGroupUserOnline(in *pb.GetGroupUserOnlineReq) (*pb.GroupUserOnlineResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	data, err := NewGetGroupMemberListLogic(l.ctx, l.svcCtx).GetGroupMemberList(&pb.GetGroupMemberListReq{
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
	return &pb.GroupUserOnlineResp{
		OnlineUser: onlineMap,
	}, nil
}
func (l *GetGroupUserOnlineLogic) checkReqParams(in *pb.GetGroupUserOnlineReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
