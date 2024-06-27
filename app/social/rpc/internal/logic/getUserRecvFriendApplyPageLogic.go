package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRecvFriendApplyPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRecvFriendApplyPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRecvFriendApplyPageLogic {
	return &GetUserRecvFriendApplyPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRecvFriendApplyPageLogic) GetUserRecvFriendApplyPage(in *socialRpc.GetUserRecvFriendApplyPageReq) (*socialRpc.FriendApplyPageResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	applyModel := model.NewSocialFriendApplyModel(l.svcCtx.SqlConn, in.PlatId)
	list, err := applyModel.Where("friend_uid = ?", in.UserId).Except("content").Order("apply_last_at DESC").Page(in.Page, in.Size).Select()
	if err != nil {
		return nil, resd.NewRpcErrWithTempCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	resp := &socialRpc.FriendApplyPageResp{
		List: make([]*socialRpc.FriendApply, len(list)),
	}
	userIds := make([]string, len(list))
	for k, v := range list {
		userIds[k] = v.UserId
		resp.List[k] = &socialRpc.FriendApply{
			Id:           v.Id,
			UserId:       v.UserId,
			FriendUid:    v.FriendUid,
			ApplyLastMsg: v.ApplyLastMsg,
			ApplyLastAt:  v.ApplyLastAt,
			OperateMsg:   v.OperateMsg,
			OperateAt:    v.OperateAt,
			StateEm:      v.StateEm,
			PlatId:       v.PlatId,
		}
	}
	userInfos, err := l.svcCtx.UserRpc.GetUserNormalInfo(l.ctx, &userRpc.GetUserInfoReq{
		Ids:    userIds,
		PlatId: in.PlatId,
	})
	for k, v := range resp.List {
		if _, ok := userInfos.Users[v.UserId]; ok {
			resp.List[k].UserSex = userInfos.Users[v.UserId].MainInfo.SexEm
			resp.List[k].UserName = userInfos.Users[v.UserId].MainInfo.Nickname
			resp.List[k].UserAvatarImg = userInfos.Users[v.UserId].MainInfo.AvatarImg
		}
	}
	if in.IsNeedTotal == 1 {
		total, err := applyModel.Where("friend_uid = ?", in.UserId).Total()
		if err != nil {
			return nil, resd.NewRpcErrWithTempCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
		}
		resp.Total = &total
	}
	return resp, nil
}
func (l *GetUserRecvFriendApplyPageLogic) checkReqParams(in *socialRpc.GetUserRecvFriendApplyPageReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
