package userInfo

import (
	"context"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
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
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
}

func NewEditMyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditMyInfoLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	return &EditMyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

func (l *EditMyInfoLogic) EditMyInfo(req *types.EditMyInfoReq) (resp *types.SuccessResp, err error) {
	userInfo := l.ctx.Value("userInfoRpc").(*user.UserInfoRpcResp)
	userBiz := biz.NewUserBiz(l.ctx, l.svcCtx)
	editUserInfo := &pb.EditUserInfoReq{}
	utild.Copy(&editUserInfo, req)
	editUserInfo.Id = userInfo.Id
	fmt.Println(editUserInfo)
	err = userBiz.EditUserInfo(editUserInfo)
	if err != nil {
		return nil, err
	}
	return &types.SuccessResp{Msg: resd.Msg(l.lang, resd.Ok)}, nil
}

func (l *EditMyInfoLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.FailCode(l.lang, resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.FailCode(l.lang, resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
