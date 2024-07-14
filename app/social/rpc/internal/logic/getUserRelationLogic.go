package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"strings"
)

type GetUserRelationLogic struct {
	*GetUserRelationLogicGen
}

func NewGetUserRelationLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserRelationLogic {
	return &GetUserRelationLogic{
		GetUserRelationLogicGen: NewGetUserRelationLogicGen(ctx, svc),
	}
}

func (l *GetUserRelationLogic) GetUserRelation(in *socialRpc.GetUserRelationReq) (*socialRpc.GetUserRelationResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	if len(in.FriendUids) == 0 {
		return nil, resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少friendUids", resd.ErrReqFieldEmpty1, "friendUids")
	}
	placeholders := make([]string, len(in.FriendUids))
	friendUids := make([]any, len(in.FriendUids))
	relats := make(map[string]int64)
	for i := range in.FriendUids {
		placeholders[i] = "?"
		friendUids[i] = in.FriendUids[i]
		if l.req.FriendUids[i] == l.req.UserId {
			relats[in.FriendUids[i]] = constd.SocialFriendStateEmSelf
		} else {
			relats[in.FriendUids[i]] = constd.SocialFriendStateEmNoRelat
		}

	}
	friendModel := model.NewSocialFriendModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	whereData := make([]any, 0)
	whereData = append(whereData, in.UserId)
	whereData = append(whereData, friendUids...)

	friends, err := friendModel.Where(fmt.Sprintf("user_id = ? and friend_uid in (%s)", strings.Join(placeholders, ",")), whereData...).Select()
	if err != nil {
		return nil, resd.NewErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	for _, v := range friends {
		relats[v.FriendUid] = v.StateEm
	}
	friendApplyModel := model.NewSocialFriendApplyModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	applys, err := friendApplyModel.Except("content").Where(fmt.Sprintf("user_id = ? and friend_uid in (%s)", strings.Join(placeholders, ",")), whereData...).Select()
	if err != nil {
		return nil, resd.NewErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	for _, v := range applys {
		//无好友关系情况下，走申请表的关系
		if relats[v.FriendUid] == constd.SocialFriendStateEmNoRelat {
			relats[v.FriendUid] = v.StateEm
		}
	}
	return &socialRpc.GetUserRelationResp{
		Relats: relats,
	}, nil
}
func (l *GetUserRelationLogic) checkReqParams(in *socialRpc.GetUserRelationReq) error {
	return nil
}
