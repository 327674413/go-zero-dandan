// Code generated by goctl. DO NOT EDIT.
package friend

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type CreateFriendApplyLogicGen struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
	hasUserInfo  bool
	mustUserInfo bool
	ReqApplyMsg  string `json:"applyMsg,optional"`
	ReqFriendUid string `json:"friendUid,optional" check:"required"`
	ReqSourceEm  int64  `json:"sourceEm,optional" check:"required"`
	HasReq       struct {
		ApplyMsg  bool
		FriendUid bool
		SourceEm  bool
	}
}

func NewCreateFriendApplyLogicGen(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFriendApplyLogicGen {
	return &CreateFriendApplyLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFriendApplyLogicGen) initReq(req *types.CreateFriendApplyReq) error {
	var err error
	if err = l.initPlat(); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}

	if req.ApplyMsg != nil {
		l.ReqApplyMsg = *req.ApplyMsg
		l.HasReq.ApplyMsg = true
	} else {
		l.HasReq.ApplyMsg = false
	}

	if req.FriendUid != nil {
		l.ReqFriendUid = *req.FriendUid
		l.HasReq.FriendUid = true
	} else {
		l.HasReq.FriendUid = false
	}

	if l.HasReq.FriendUid == false {
		return resd.NewErrWithTempCtx(l.ctx, "缺少参数FriendUid", resd.ReqFieldRequired1, "*FriendUid")
	}

	if l.ReqFriendUid == "" {
		return resd.NewErrWithTempCtx(l.ctx, "FriendUid不得为空", resd.ReqFieldEmpty1, "*FriendUid")
	}

	if req.SourceEm != nil {
		l.ReqSourceEm = *req.SourceEm
		l.HasReq.SourceEm = true
	} else {
		l.HasReq.SourceEm = false
	}

	if l.HasReq.SourceEm == false {
		return resd.NewErrWithTempCtx(l.ctx, "缺少参数SourceEm", resd.ReqFieldRequired1, "*SourceEm")
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}

func (l *CreateFriendApplyLogicGen) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *CreateFriendApplyLogicGen) initPlat() (err error) {
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