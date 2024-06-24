package logic

import (
	"context"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFriendApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFriendApplyLogic {
	return &CreateFriendApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFriendApplyLogic) CreateFriendApply(in *pb.CreateFriendApplyReq) (*pb.CreateFriendApplyResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	friendModel := model.NewSocialFriendModel(l.svcCtx.SqlConn, in.PlatId)
	//查询我加对方
	existSelf, err := friendModel.Where("user_id = ? and friend_uid = ?", in.UserId, in.FriendUid).Find()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	//查询对方加我
	existFriend, err := friendModel.Where("friend_uid = ? and user_id = ?", in.UserId, in.FriendUid).Find()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	//如果我加了对方
	if existSelf != nil {
		//如果我加对方是好友， 对方加我也是好友
		if existSelf.StateEm == constd.SocialFriendStateEmPass && existFriend != nil && existFriend.StateEm == constd.SocialFriendStateEmPass {
			return nil, resd.NewRpcErrCtx(l.ctx, "已经是好友", resd.SocialAlreadyFriend)
		}
	} else if existFriend != nil {
		//我没有加对方，但对方有加我
		//我被拉黑
		if existFriend.StateEm == constd.SocialFriendStateEmBlack {
			return nil, resd.NewRpcErrCtx(l.ctx, "对方把你拉黑", resd.SocialAlreadyBlackMe)
		}
		//我还是对方好友
		if existFriend.StateEm == constd.SocialFriendStateEmPass {
			//先当作还是要对方点同意来弄，所以先不用处理
		}
	}

	applyModel := model.NewSocialFriendApplyModel(l.svcCtx.SqlConn, in.PlatId)
	applyId := utild.MakeId()
	addData := &model.SocialFriendApply{
		Id:       applyId,
		UserId:   in.UserId,
		ApplyUid: in.FriendUid,
		ApplyMsg: in.ApplyMsg,
		ApplyAt:  time.Now().Unix(),
	}
	_, err = applyModel.Insert(addData)
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, "创建好友申请失败", resd.MysqlInsertErr)
	}
	return &pb.CreateFriendApplyResp{
		ApplyId: applyId,
	}, nil
}
func (l *CreateFriendApplyLogic) checkReqParams(in *pb.CreateFriendApplyReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少PlatId", resd.ReqFieldRequired1, "PlatId")
	}
	if in.UserId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少UserId", resd.ReqFieldRequired1, "UserId")
	}
	if in.FriendUid == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少FriendUid", resd.ReqFieldRequired1, "FriendUid")
	}
	if in.FriendUid == in.UserId {
		return resd.NewRpcErrCtx(l.ctx, "不能添加自己", resd.SocialNotAddSelf)
	}
	return nil
}
