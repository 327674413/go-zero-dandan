package logic

import (
	"context"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"

	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"

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

func (l *GetUserCronyListLogic) GetUserCronyList(in *pb.IdReq) (*pb.UserCronyList, error) {
	if in.PlatId == 0 {
		return l.rpcFail(resd.NewErrCtx(l.ctx, "未传入应用标识"))
	}
	model := model.NewUserCronyModel(l.svcCtx.SqlConn, in.PlatId)
	data, _ := model.Ctx(l.ctx).Where("owner_user_id = ? AND type_em = ?", in.Id, constd.UserCronyTypeEmNormal).Select()
	list := make([]*pb.UserCronyInfo, 0)
	// todo::为什么数据库查出来后，不能直接赋值给&pb.UserCronyInfo{}呢，导致这里要人工循环转结构
	for _, v := range data {
		d := &pb.UserCronyInfo{
			OwnerUserId:  v.OwnerUserId,
			TargetUserId: v.TargetUserId,
			TypeEm:       v.TypeEm,
			CreateAt:     v.CreateAt,
			OwnerRemark:  v.OwnerRemark,
		}
		list = append(list, d)
	}
	return &pb.UserCronyList{List: list}, nil
}

func (l *GetUserCronyListLogic) rpcFail(err error) (*pb.UserCronyList, error) {
	return nil, resd.RpcErrEncode(err)
}
