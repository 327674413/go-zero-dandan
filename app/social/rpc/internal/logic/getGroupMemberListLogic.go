package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberListLogic {
	return &GetGroupMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMemberListLogic) GetGroupMemberList(in *socialRpc.GetGroupMemberListReq) (*socialRpc.GroupMemberListResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	m := model.NewSocialGroupMemberModel(l.svcCtx.SqlConn, in.PlatId)
	list, err := m.Where("group_id = ?", in.GroupId).Select()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	resp := &socialRpc.GroupMemberListResp{List: make([]*socialRpc.GroupMember, 0, len(list))}
	for _, item := range list {
		resp.List = append(resp.List, &socialRpc.GroupMember{
			Id:           item.Id,
			GroupId:      item.GroupId,
			UserId:       item.UserId,
			RoleLevel:    item.RoleLevel,
			JoinAt:       item.JoinAt,
			JoinSourceEm: item.JoinSourceEm,
			InviteUid:    item.InviteUid,
			OperateUid:   item.OperateUid,
			PlatId:       item.PlatId,
		})
	}
	return &socialRpc.GroupMemberListResp{}, nil
}
func (l *GetGroupMemberListLogic) checkReqParams(in *socialRpc.GetGroupMemberListReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	if in.GroupId == "" {
		return resd.NewErrWithTempCtx(l.ctx, "缺少groupId", resd.ReqFieldRequired1, "groupId")
	}
	return nil
}
