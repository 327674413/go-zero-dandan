package logic

import (
	"context"
	"go-zero-dandan/app/plat/model"
	"go-zero-dandan/app/plat/rpc/internal/svc"
	"go-zero-dandan/app/plat/rpc/types/platRpc"
)

type GetOneLogic struct {
	*GetOneLogicGen
}

func NewGetOneLogic(ctx context.Context, svc *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		GetOneLogicGen: NewGetOneLogicGen(ctx, svc),
	}
}

func (l *GetOneLogic) GetOne(in *platRpc.IdReq) (*platRpc.PlatInfo, error) {
	if err := l.init(in); err != nil {
		return nil, l.resd.Error(err)
	}

	platModel := model.NewPlatMainModel(l.ctx, l.svc.SqlConn)
	platMain, err := platModel.WhereId(l.req.Id).Find()
	if platMain == nil {
		return nil, l.resd.Error(err)
	}
	res := &platRpc.PlatInfo{
		Id:     platMain.Id,
		ClasEm: platMain.ClasEm,
	}
	return res, nil
}

func (l *GetOneLogic) init(req *platRpc.IdReq) (err error) {
	if err = l.initReq(req); err != nil {
		return l.resd.Error(err)
	}
	return nil
}
