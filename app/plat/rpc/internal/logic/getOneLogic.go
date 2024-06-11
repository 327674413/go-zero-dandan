package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/plat/model"
	"go-zero-dandan/app/plat/rpc/internal/svc"
	"go-zero-dandan/app/plat/rpc/types/pb"
	"go-zero-dandan/common/resd"
)

type GetOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneLogic) GetOne(in *pb.IdReq) (*pb.PlatInfo, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	platModel := model.NewPlatMainModel(l.svcCtx.SqlConn)
	platMain, err := platModel.Ctx(l.ctx).WhereId(in.Id).Find()
	if platMain == nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	res := &pb.PlatInfo{
		Id:     platMain.Id,
		ClasEm: platMain.ClasEm,
	}
	return res, nil
}
func (l *GetOneLogic) checkReqParams(in *pb.IdReq) error {
	if in.Id == 0 {
		return resd.NewErrWithTempCtx(l.ctx, "参数缺少id", resd.ReqFieldRequired1, "id")
	}
	return nil
}
