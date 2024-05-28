package crony

import (
	"context"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/common/utild/copier"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type GetMyCronyListLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       int64
	platClasEm   int64
}

func NewGetMyCronyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyCronyListLogic {
	return &GetMyCronyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyCronyListLogic) GetMyCronyList(req *types.GetUserCronyListReq) (resp *types.GetUserCronyListResp, err error) {
	if err = l.initPlatAndUsr(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	data, err := l.svcCtx.UserRpc.GetUserCronyList(l.ctx, &user.GetUserCronyListReq{
		PlatId:        &l.platId,
		OwnerUserId:   &l.userMainInfo.Id,
		OwnerUserName: req.OwnerUserName,
		GroupId:       req.GroupId,
		TypeEms:       req.TypeEms,
		AddStartTime:  req.AddStartTime,
		AddEndTime:    req.AddEndTime,
		IsNeedTotal:   req.IsNeedTotal,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	/*list := make([]*types.UserCronyInfo, 0)
	for _, v := range data.List {
		list = append(list, &types.UserCronyInfo{
			Id:               &v.Id,
			OwnerUserId:      &v.OwnerUserId,
			TargetUserId:     &v.TargetUserId,
			TargetUserName:   &v.TargetUserName,
			TargetUserAvatar: &v.TargetUserAvatar,
			Remark:           &v.Remark,
		})
	}*/
	resp = &types.GetUserCronyListResp{}
	err = copier.Copy(&resp.List, data.List)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if data.Total != nil {
		resp.Total = data.Total
	}
	return resp, nil
}

func (l *GetMyCronyListLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
func (l *GetMyCronyListLogic) initPlatAndUsr() (err error) {
	if err = l.initPlat(); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}
