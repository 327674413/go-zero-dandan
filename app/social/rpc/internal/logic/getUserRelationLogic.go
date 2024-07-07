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

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRelationLogic {
	return &GetUserRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRelationLogic) GetUserRelation(in *socialRpc.GetUserRelationReq) (*socialRpc.GetUserRelationResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	if len(in.FriendUids) == 0 {
		return nil, resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少friendUids", resd.ReqFieldEmpty1, "friendUids")
	}
	if in.UserId == "" {
		return nil, resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少userId", resd.ReqFieldRequired1, "userId")
	}
	placeholders := make([]string, len(in.FriendUids))
	friendUids := make([]any, len(in.FriendUids))
	relats := make(map[string]int64)
	for i := range in.FriendUids {
		placeholders[i] = "?"
		friendUids[i] = in.FriendUids[i]
		if in.FriendUids[i] == in.UserId {
			relats[in.FriendUids[i]] = constd.SocialFriendStateEmSelf
		} else {
			relats[in.FriendUids[i]] = constd.SocialFriendStateEmNoRelat
		}

	}
	friendModel := model.NewSocialFriendModel(l.ctx, l.svcCtx.SqlConn, in.PlatId)
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
	friendApplyModel := model.NewSocialFriendApplyModel(l.ctx, l.svcCtx.SqlConn, in.PlatId)
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
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
