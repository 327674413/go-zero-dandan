package logic

import (
	"context"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCronyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCronyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCronyListLogic {
	return &GetUserCronyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCronyListLogic) GetUserCronyList(in *pb.GetUserCronyListReq) (*pb.GetUserCronyListResp, error) {
	if in.PlatId == nil || *in.PlatId == 0 {
		return l.rpcFail(resd.NewErrCtx(l.ctx, "未传入应用标识"))
	}
	if in.OwnerUserId == nil || *in.OwnerUserId == 0 {
		return l.rpcFail(resd.NewErrCtx(l.ctx, "ownerUserId不得为空"))
	}
	userCronyModel := model.NewUserCronyModel(l.svcCtx.SqlConn, *in.PlatId)
	var err error
	var data []*model.UserCrony
	var total int64
	if in.IsNeedTotal != nil && *in.IsNeedTotal == 1 {
		data, total, err = userCronyModel.Ctx(l.ctx).Where("owner_user_id = ? AND type_em = ?", *in.OwnerUserId, constd.UserCronyTypeEmNormal).SelectWithTotal()
	} else {
		data, err = userCronyModel.Ctx(l.ctx).Where("owner_user_id = ? AND type_em = ?", *in.OwnerUserId, constd.UserCronyTypeEmNormal).Select()
	}
	if err != nil {
		return l.rpcFail(resd.ErrorCtx(l.ctx, err))
	}
	resp := &pb.GetUserCronyListResp{}
	err = copier.Copy(&resp.List, data)
	if err != nil {
		return l.rpcFail(resd.ErrorCtx(l.ctx, err))
	}
	if in.IsNeedTotal != nil && *in.IsNeedTotal == 1 {
		resp.Total = &total
	}
	return resp, nil
}

func (l *GetUserCronyListLogic) rpcFail(err error) (*pb.GetUserCronyListResp, error) {
	return nil, resd.RpcErrEncode(err)
}
