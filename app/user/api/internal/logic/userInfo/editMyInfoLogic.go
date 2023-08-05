package userInfo

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/api/internal/biz"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type EditMyInfoLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	platId     int64
	platClasEm int64
}

func NewEditMyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditMyInfoLogic {
	return &EditMyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditMyInfoLogic) EditMyInfo(req *types.EditMyInfoReq) (resp *types.SuccessResp, err error) {
	userInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return nil, resd.NewErr("获取用户信息失败")
	}
	userBiz := biz.NewUserBiz(l.ctx, l.svcCtx)
	editUserInfo := &pb.EditUserInfoReq{}
	utild.Copy(&editUserInfo, req)
	editUserInfo.Id = userInfo.Id
	err = userBiz.EditUserInfo(editUserInfo)
	if err != nil {
		return nil, err
	}
	return &types.SuccessResp{Msg: ""}, nil
}

func (l *EditMyInfoLogic) initPlat() (err error) {
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
