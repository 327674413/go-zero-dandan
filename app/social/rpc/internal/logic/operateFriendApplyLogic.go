package logic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/threading"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
)

type OperateFriendApplyLogic struct {
	*OperateFriendApplyLogicGen
}

func NewOperateFriendApplyLogic(ctx context.Context, svc *svc.ServiceContext) *OperateFriendApplyLogic {
	return &OperateFriendApplyLogic{
		OperateFriendApplyLogicGen: NewOperateFriendApplyLogicGen(ctx, svc),
	}
}

// 处理申请，根据处理人，将申请状态，来源等数据更新到好友表上

func (l *OperateFriendApplyLogic) OperateFriendApply(in *socialRpc.OperateFriendApplyReq) (resp *socialRpc.ResultResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	//查询申请信息
	applyModel := model.NewSocialFriendApplyModel(l.ctx, l.svc.SqlConn, l.req.PlatId)
	apply, err := applyModel.WhereId(l.req.ApplyId).Find()
	if err != nil {
		return nil, l.resd.Error(err)
	}
	//权限判定
	if l.req.SysRoleEm == constd.SysRoleEmUser && l.req.SysRoleUid != apply.UserId {
		//用户操作，如用户处理申请的api调用
		return nil, l.resd.NewErr(resd.ErrAuthOperateUser)
	}
	//操作状态校验
	if apply.StateEm != constd.SocialFriendStateEmApply {
		//非申请中不可操作
		return nil, l.resd.NewErr(resd.ErrAuthOperateState)
	}
	err = dao.WithTrans(l.ctx, l.svc.SqlConn, func(tx *sql.Tx) error {
		if l.req.OperateStateEm == constd.SocialFriendStateEmReject {
			//拒绝，更新申请状态
			_, err = applyModel.WhereId(apply.Id).TxUpdate(tx, map[dao.TableField]any{
				model.SocialFriendApply_StateEm:    l.req.OperateStateEm,
				model.SocialFriendApply_OperateMsg: l.req.OperateMsg,
			})
		} else if l.req.OperateStateEm == constd.SocialFriendStateEmPass {
			// 通过，更新申请状态
			_, err = applyModel.WhereId(apply.Id).TxUpdate(tx, map[dao.TableField]any{
				model.SocialFriendApply_StateEm:    l.req.OperateStateEm,
				model.SocialFriendApply_OperateMsg: l.req.OperateMsg,
			})
			if err != nil {
				return l.resd.Error(err)
			}
		} else {
			//暂不支持其他类型处理，如撤销等
			return l.resd.NewErr(resd.ErrAuthOperateState)
		}
		// 查询好友数据
		friendModel := model.NewSocialFriendModel(l.ctx, l.svc.SqlConn)
		// 查询 对方 - 我， 处理加我申请，所以FriendUid是我， UserId是对方
		friendRelat, err := friendModel.Where("user_id = ? and friend_uid = ?", apply.UserId, apply.FriendUid).Find()
		if err != nil {
			return l.resd.Error(err)
		}
		//没找到对方 - 我数据
		if friendRelat == nil {
			return l.resd.NewErr(resd.ErrDataBiz)
		}
		//查询 我 - 对方信息
		myRelat, err := friendModel.Where("user_id = ? and friend_uid = ?", apply.FriendUid, apply.UserId).Find()
		if err != nil {
			return l.resd.Error(err)
		}
		//没找到我 - 对方数据
		if myRelat == nil {
			return l.resd.NewErr(resd.ErrDataBiz)
		}
		//获取对方用户信息
		//暂时决定不冗余了，这里线不用获取两边人的信息了
		//friendInfo, err := l.svc.UserRpc.GetUserById(l.ctx, &userRpc.IdReq{
		//	Id: &apply.UserId,
		//})
		//if err != nil {
		//	return l.resd.Error(err)
		//}
		////获取我的用户信息
		//myInfo, err := l.svc.UserRpc.GetUserById(l.ctx, &userRpc.IdReq{
		//	Id: &apply.FriendUid,
		//})
		//if err != nil {
		//	return l.resd.Error(err)
		//}
		// 更新对方 - 我的好友关系
		_, err = friendModel.WhereId(friendRelat.Id).TxUpdate(tx, map[dao.TableField]any{
			model.SocialFriend_StateEm:  l.req.OperateStateEm,
			model.SocialFriend_SourceEm: apply.SourceEm,
		})
		if err != nil {
			return resd.ErrorCtx(l.ctx, err)
		}
		// 更新我 - 对方好友关系
		_, err = friendModel.WhereId(myRelat.Id).TxUpdate(tx, map[dao.TableField]any{
			model.SocialFriend_StateEm:  l.req.OperateStateEm,
			model.SocialFriend_SourceEm: -1 * apply.SourceEm,
		})
		if err != nil {
			return l.resd.Error(err)
		}
		return nil
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	//异步发送好友通过的通知

	asyncCtx := contextx.ValueOnlyFrom(l.ctx)
	msgClasEm := int64(constd.MsgClasEmFriendApplyOperated)
	threading.GoSafe(func() {
		l.svc.ImRpc.SendSysMsg(asyncCtx, &imRpc.SendSysMsgReq{
			UserId:     &apply.UserId,
			MsgClasEm:  &msgClasEm,
			MsgContent: &apply.Id,
		})
	})
	return &socialRpc.ResultResp{Result: true}, nil
}
func (l *OperateFriendApplyLogic) checkReqParams(in *socialRpc.OperateFriendApplyReq) error {
	return nil
}
