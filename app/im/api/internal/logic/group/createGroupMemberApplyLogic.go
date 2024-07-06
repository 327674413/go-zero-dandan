package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type CreateGroupMemberApplyLogic struct {
	*CreateGroupMemberApplyLogicGen
}

func NewCreateGroupMemberApplyLogic(ctx context.Context, svc *svc.ServiceContext) *CreateGroupMemberApplyLogic {
	return &CreateGroupMemberApplyLogic{
		CreateGroupMemberApplyLogicGen: NewCreateGroupMemberApplyLogicGen(ctx, svc),
	}
}

func (l *CreateGroupMemberApplyLogic) CreateGroupMemberApply(req *types.CreateGroupMemberApplyReq) (resp *types.CreateGroupMemberApplyResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	return
}

func (l *CreateGroupMemberApplyLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *CreateGroupMemberApplyLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platId, _ := l.ctx.Value("platId").(string)
	if platId == "" {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platId
	l.platClasEm = platClasEm
	return nil
}
