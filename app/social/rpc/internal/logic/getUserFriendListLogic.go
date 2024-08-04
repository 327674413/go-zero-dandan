package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/constd"
)

type GetUserFriendListLogic struct {
	*GetUserFriendListLogicGen
}

func NewGetUserFriendListLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserFriendListLogic {
	return &GetUserFriendListLogic{
		GetUserFriendListLogicGen: NewGetUserFriendListLogicGen(ctx, svc),
	}
}

func (l *GetUserFriendListLogic) GetUserFriendList(in *socialRpc.GetUserFriendListReq) (*socialRpc.FriendListResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	if err := l.checkReqParams(in); err != nil {
		return nil, l.resd.Error(err)
	}
	m := model.NewSocialFriendModel(l.ctx, l.svc.SqlConn, l.req.PlatId)

	list, err := m.Where("user_id = ? AND state_em=?", l.req.UserId, constd.SocialFriendStateEmPass).Select()

	if err != nil {
		return nil, l.resd.Error(err)
	}
	resp := &socialRpc.FriendListResp{
		List: make([]*socialRpc.FriendInfo, 0),
	}
	userIds := make([]string, len(list))
	for k, v := range list {
		userIds[k] = v.FriendUid
	}
	userInfos, err := l.svc.UserRpc.GetUserNormalInfo(l.ctx, &userRpc.GetUserInfoReq{
		Ids: userIds,
	})
	for _, v := range list {
		friendInfo := &socialRpc.FriendInfo{
			Id:           v.Id,
			UserId:       v.UserId,
			FriendRemark: v.FriendRemark,
			SourceEm:     v.SourceEm,
			FriendUid:    v.FriendUid,
			PlatId:       v.PlatId,
		}
		if userInfo, ok := userInfos.Users[v.FriendUid]; ok {
			friendInfo.FriendSexEm = userInfo.MainInfo.SexEm
			friendInfo.FriendName = userInfo.MainInfo.Nickname
			friendInfo.FriendIcon = userInfo.MainInfo.AvatarImg
		}
		resp.List = append(resp.List, friendInfo)
	}

	return resp, nil
}
func (l *GetUserFriendListLogic) checkReqParams(in *socialRpc.GetUserFriendListReq) error {
	return nil
}
