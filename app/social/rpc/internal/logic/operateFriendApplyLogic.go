package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
)

type OperateFriendApplyLogic struct {
	ctx context.Context
	svc *svc.ServiceContext
	logx.Logger
}

func NewOperateFriendApplyLogic(ctx context.Context, svc *svc.ServiceContext) *OperateFriendApplyLogic {
	return &OperateFriendApplyLogic{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
	}
}

// 处理申请，根据处理人，将申请状态，来源等数据更新到好友表上
func (l *OperateFriendApplyLogic) OperateFriendApply(in *socialRpc.OperateFriendApplyReq) (*socialRpc.ResultResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	//查询申请信息
	applyModel := model.NewSocialFriendApplyModel(l.svc.SqlConn, in.PlatId)
	apply, err := applyModel.WhereId(in.ApplyId).Find()
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	//权限判定
	if in.SysRoleEm == constd.SysRoleEmUser && in.SysRoleUid != apply.UserId {
		//用户操作，如用户处理申请的api调用
		return nil, resd.NewRpcErrCtx(l.ctx, "", resd.AuthOperateUserErr)
	}
	//操作状态校验
	if apply.StateEm != constd.SocialFriendStateEmApply {
		//非申请中不可操作
		return nil, resd.NewRpcErrCtx(l.ctx, "", resd.AuthOperateStateErr)
	}
	tx, err := dao.StartTrans(l.svc.SqlConn, l.ctx)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if in.OperateStateEm == constd.SocialFriendStateEmReject {
		//拒绝，更新申请状态
		_, err = applyModel.WhereId(apply.Id).TxUpdate(tx, map[dao.TableField]any{
			model.SocialFriendApply_StateEm:    apply.StateEm,
			model.SocialFriendApply_OperateMsg: in.OperateMsg,
		})
	} else if in.OperateStateEm == constd.SocialFriendStateEmPass {
		// 通过，更新申请状态
		_, err = applyModel.WhereId(apply.Id).TxUpdate(tx, map[dao.TableField]any{
			model.SocialFriendApply_StateEm:    in.OperateStateEm,
			model.SocialFriendApply_OperateMsg: in.OperateMsg,
		})
		if err != nil {
			return nil, resd.ErrorCtx(l.ctx, err, resd.MysqlUpdateErr)
		}
	} else {
		//暂不支持其他类型处理，如撤销等
		return nil, resd.NewRpcErrCtx(l.ctx, "", resd.AuthOperateStateErr)
	}
	// 查询好友数据
	friendModel := model.NewSocialFriendModel(l.svc.SqlConn)
	// 查询 对方 - 我， 处理加我申请，所以FriendUid是我， UserId是对方
	friendRelat, err := friendModel.Where("user_id = ? and friend_uid = ?", apply.UserId, apply.FriendUid).Find()
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	//没找到对方 - 我数据
	if friendRelat == nil {
		return nil, resd.NewErrCtx(l.ctx, "未找到好友关系1", resd.DataBizErr)
	}

	//查询 我 - 对方信息
	myRelat, err := friendModel.Where("user_id = ? and friend_uid = ?", apply.FriendUid, apply.UserId).Find()
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	//没找到我 - 对方数据
	if myRelat == nil {
		return nil, resd.NewErrCtx(l.ctx, "未找到好友关系2", resd.DataBizErr)
	}
	//获取对方用户信息
	friendInfo, err := l.svc.UserRpc.GetUserById(l.ctx, &userRpc.IdReq{
		PlatId: in.PlatId,
		Id:     apply.UserId,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	//获取我的用户信息
	myInfo, err := l.svc.UserRpc.GetUserById(l.ctx, &userRpc.IdReq{
		PlatId: in.PlatId,
		Id:     apply.FriendUid,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	// 更新对方 - 我的好友关系
	_, err = friendModel.WhereId(friendRelat.Id).TxUpdate(tx, map[dao.TableField]any{
		model.SocialFriend_StateEm:    in.OperateStateEm,
		model.SocialFriend_FriendName: myInfo.Nickname,
		model.SocialFriend_FriendIcon: myInfo.AvatarImg,
		model.SocialFriend_SourceEm:   apply.SourceEm,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	// 更新我 - 对方好友关系
	_, err = friendModel.WhereId(myRelat.Id).TxUpdate(tx, map[dao.TableField]any{
		model.SocialFriend_StateEm:    in.OperateStateEm,
		model.SocialFriend_FriendName: friendInfo.Nickname,
		model.SocialFriend_FriendIcon: friendInfo.AvatarImg,
		model.SocialFriend_SourceEm:   -1 * apply.SourceEm,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err := tx.Commit(); err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, "", resd.MysqlCommitErr)
	}
	//todo::给对方发消息，通知对方通过
	return &socialRpc.ResultResp{Result: true}, nil
}
func (l *OperateFriendApplyLogic) checkReqParams(in *socialRpc.OperateFriendApplyReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}