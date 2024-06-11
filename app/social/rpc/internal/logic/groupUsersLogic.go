package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUsersLogic {
	return &GroupUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupUsersLogic) GroupUsers(in *pb.GroupUsersReq) (*pb.GroupUsersResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	m := model.NewSocialGroupMemberModel(l.svcCtx.SqlConn, in.PlatId)
	list, err := m.Where("group_id = ?", in.GroupId).Select()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	resp := &pb.GroupUsersResp{List: make([]*pb.GroupMembers, 0, len(list))}
	for _, item := range list {
		resp.List = append(resp.List, &pb.GroupMembers{
			Id:            utild.AnyToInt64(item.Id),
			GroupId:       item.GroupId,
			UserId:        item.UserId,
			RoleLevel:     item.RoleLevel,
			JoinAt:        item.JoinAt,
			JoinSourceEm:  item.JoinSourceEm,
			InviteUserId:  item.InviteUserId,
			OperateUserId: item.OperateUserId,
			PlatId:        item.PlatId,
		})
	}
	return resp, nil
}
func (l *GroupUsersLogic) checkReqParams(in *pb.GroupUsersReq) error {
	if in.PlatId == 0 {
		return resd.NewErrWithTempCtx(l.ctx, "缺少platId", resd.ReqFieldRequired1, "platId")
	}
	if in.GroupId == 0 {
		return resd.NewErrWithTempCtx(l.ctx, "缺少groupId", resd.ReqFieldRequired1, "groupId")
	}
	return nil
}
