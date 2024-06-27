package friend

import (
	"context"
	"fmt"
	"strings"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type OperateMyRecvFriendApplyLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
}

func NewOperateMyRecvFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateMyRecvFriendApplyLogic {
	return &OperateMyRecvFriendApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateMyRecvFriendApplyLogic) OperateMyRecvFriendApply(req *types.OperateMyRecvFriendApplyReq) (resp *types.ResultResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if req.ApplyId == nil {
		return nil, resd.NewErrWithTempCtx(l.ctx, "缺少参数applyId", resd.ReqFieldRequired1, "applyId")
	}
	applyId := strings.TrimSpace(*req.ApplyId)
	if applyId == "" {
		return nil, resd.NewErrWithTempCtx(l.ctx, "缺少参数applyId", resd.ReqFieldEmpty1, "applyId")
	}
	if req.StateEm == nil {
		return nil, resd.NewErrWithTempCtx(l.ctx, "缺少参数stateEm", resd.ReqFieldRequired1, "stateEm")
	}
	StateEm := *req.StateEm

	fmt.Println(StateEm)
	return
}

func (l *OperateMyRecvFriendApplyLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *OperateMyRecvFriendApplyLogic) initPlat() (err error) {
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
