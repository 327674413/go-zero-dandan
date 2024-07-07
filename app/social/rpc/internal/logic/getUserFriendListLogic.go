package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFriendListLogic {
	return &GetUserFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFriendListLogic) GetUserFriendList(in *socialRpc.GetUserFriendListReq) (*socialRpc.FriendListResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	if in.UserId == "" {
		return nil, resd.NewRpcErrWithTempCtx(l.ctx, "缺少UserId", resd.ReqFieldRequired1, "*UserId")
	}
	m := model.NewSocialFriendModel(l.ctx, l.svcCtx.SqlConn, in.PlatId)

	list, err := m.Where("user_id = ?", in.UserId).Select()

	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	resp := &socialRpc.FriendListResp{
		List: make([]*socialRpc.FriendInfo, 0),
	}
	if err = copier.Copy(&resp.List, list); err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	return resp, nil
	return &socialRpc.FriendListResp{}, nil
}
func (l *GetUserFriendListLogic) checkReqParams(in *socialRpc.GetUserFriendListReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
