package friend

import (
	"context"
	"go-zero-dandan/app/social/rpc/types/pb"
	"strings"
	"time"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type CreateFriendApplyLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
	friendUid    string
	applyMsg     string
}

func NewCreateFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFriendApplyLogic {
	return &CreateFriendApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFriendApplyLogic) CreateFriendApply(req *types.CreateFriendApplyReq) (resp *types.ResultResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.checkParam(req); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	_, err = l.svcCtx.SocialRpc.CreateFriendApply(l.ctx, &pb.CreateFriendApplyReq{
		PlatId:    l.platId,
		UserId:    l.userMainInfo.Id,
		FriendUid: l.friendUid,
		ApplyMsg:  l.applyMsg,
		ApplyAt:   time.Now().Unix(),
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	return &types.ResultResp{Result: true}, nil
}
func (l *CreateFriendApplyLogic) checkParam(req *types.CreateFriendApplyReq) error {
	if req.FriendUid == nil {
		return resd.NewErrWithTempCtx(l.ctx, "", resd.ReqFieldRequired1, "friendUid")
	}
	l.friendUid = strings.TrimSpace(*req.FriendUid)
	if l.friendUid == "" {
		return resd.NewErrWithTempCtx(l.ctx, "", resd.ReqFieldEmpty1, "friendUid")
	}
	if req.ApplyMsg != nil {
		l.applyMsg = strings.TrimSpace(*req.ApplyMsg)
	}

	return nil
}

func (l *CreateFriendApplyLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *CreateFriendApplyLogic) initPlat() (err error) {
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
