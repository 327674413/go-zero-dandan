package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/fmtd"
	"go-zero-dandan/common/resd"
)

type GetUserRecvFriendApplyPageLogic struct {
	*GetUserRecvFriendApplyPageLogicGen
}

func NewGetUserRecvFriendApplyPageLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserRecvFriendApplyPageLogic {
	return &GetUserRecvFriendApplyPageLogic{
		GetUserRecvFriendApplyPageLogicGen: NewGetUserRecvFriendApplyPageLogicGen(ctx, svc),
	}
}

func (l *GetUserRecvFriendApplyPageLogic) GetUserRecvFriendApplyPage(in *socialRpc.GetUserRecvFriendApplyPageReq) (*socialRpc.FriendApplyPageResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	applyModel := model.NewSocialFriendApplyModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	list, err := applyModel.Where("friend_uid = ?", in.UserId).Except("content").Order("apply_last_at DESC").Page(l.req.Page, l.req.Size).Select()
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
	fmtd.Info(l.meta)
	userInfos, err := l.svc.UserRpc.GetUserNormalInfo(l.ctx, &userRpc.GetUserInfoReq{
		Ids:    userIds,
		PlatId: l.meta.PlatId,
	})
	for k, v := range resp.List {
		if _, ok := userInfos.Users[v.UserId]; ok {
			resp.List[k].UserSex = userInfos.Users[v.UserId].MainInfo.SexEm
			resp.List[k].UserName = userInfos.Users[v.UserId].MainInfo.Nickname
			resp.List[k].UserAvatarImg = userInfos.Users[v.UserId].MainInfo.AvatarImg
		}
	}
	if l.req.IsNeedTotal == 1 {
		total, err := applyModel.Where("friend_uid = ?", in.UserId).Total()
		if err != nil {
			return nil, resd.NewRpcErrWithTempCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
		}
		resp.Total = total
	}
	return resp, nil
}
func (l *GetUserRecvFriendApplyPageLogic) checkReqParams(in *socialRpc.GetUserRecvFriendApplyPageReq) error {
	return nil
}
