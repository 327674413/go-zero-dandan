package logic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/threading"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/biz"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"time"
)

type CreateFriendApplyLogic struct {
	*CreateFriendApplyLogicGen
}

func NewCreateFriendApplyLogic(ctx context.Context, svc *svc.ServiceContext) *CreateFriendApplyLogic {
	return &CreateFriendApplyLogic{
		CreateFriendApplyLogicGen: NewCreateFriendApplyLogicGen(ctx, svc),
	}
}

func (l *CreateFriendApplyLogic) CreateFriendApply(in *socialRpc.CreateFriendApplyReq) (*socialRpc.CreateFriendApplyResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	if err := l.checkReqParams(); err != nil {
		return nil, l.resd.Error(err)
	}
	friendModel := model.NewSocialFriendModel(l.ctx, l.svc.SqlConn, l.req.PlatId)
	//我申请别人为好友，UserId是我，FriendUid是好友
	//查询我-对方好友数据
	existSelf, err := friendModel.Where("user_id = ? and friend_uid = ?", l.req.UserId, l.req.FriendUid).Find()
	if err != nil {
		return nil, l.resd.Error(err)
	}
	//查询对方-我好友数据
	existFriend, err := friendModel.Where("friend_uid = ? and user_id = ?", l.req.UserId, l.req.FriendUid).Find()
	if err != nil {
		return nil, l.resd.Error(err)
	}
	//分析我-对方的关系
	if existSelf != nil {
		//如果我-对方是好友 并且 对方-我也是好友
		if existSelf.StateEm == constd.SocialFriendStateEmPass && existFriend != nil && existFriend.StateEm == constd.SocialFriendStateEmPass {
			return nil, l.resd.NewErr(resd.ErrSocialAlreadyFriend)
		}
	}
	//分析对方-我的关系
	if existFriend != nil {
		//对方拉黑了我，不允许申请
		if existFriend.StateEm == constd.SocialFriendStateEmBlack {
			return nil, l.resd.NewErr(resd.ErrSocialAlreadyBlackMe)
		}
		//对方还是我的好友，而我不是对方好友
		if existFriend.StateEm == constd.SocialFriendStateEmPass {
			//先当作还是要对方点同意来弄，所以先不用处理
		}
	}
	var applyId string
	//基础校验完成，可以添加，开启事务
	err = dao.WithTrans(l.ctx, l.svc.SqlConn, func(tx *sql.Tx) error {
		//确保有我 - 对方的关系
		if existSelf == nil {
			// 不存在，创建
			if _, err = friendModel.TxInsert(tx, &model.SocialFriend{
				Id:        utild.MakeId(),
				UserId:    l.req.UserId,
				FriendUid: l.req.FriendUid,
				StateEm:   constd.SocialFriendStateEmApply,
			}); err != nil {
				return l.resd.Error(err)
			}
		} else {
			//存在，更新我 - 对方的状态
			if _, err = friendModel.WhereId(existSelf.Id).TxUpdate(tx, map[dao.TableField]any{
				model.SocialFriend_StateEm: constd.SocialFriendStateEmApply,
			}); err != nil {
				return l.resd.Error(err)
			}
		}
		//确保有对方 - 我的关系
		if existFriend == nil {
			//不存在，创建
			if _, err = friendModel.TxInsert(tx, &model.SocialFriend{
				Id:        utild.MakeId(),
				UserId:    l.req.FriendUid,
				FriendUid: l.req.UserId,
				StateEm:   constd.SocialFriendStateEmApply,
			}); err != nil {
				return l.resd.Error(err)
			}
		} else {
			//存在，更新 对我 - 我的状态
			if _, err = friendModel.WhereId(existFriend.Id).TxUpdate(tx, map[dao.TableField]any{
				model.SocialFriend_StateEm: constd.SocialFriendStateEmApply,
			}); err != nil {
				return l.resd.Error(err)
			}
		}

		//查看是否存在过申请
		applyModel := model.NewSocialFriendApplyModel(l.ctx, l.svc.SqlConn, l.req.PlatId)
		existApply, err := applyModel.Where("friend_uid = ? AND user_id = ?", in.FriendUid, in.UserId).Find()
		if err != nil {
			return l.resd.Error(err)
		}

		if existApply != nil {
			//存在申请，更新与添加申请消息
			applyId = existApply.Id
			newContent := biz.AddApplyRecord(existApply.Content.String, l.req.UserId, l.req.ApplyMsg, l.req.SourceEm, constd.SocialFriendApplyRecordTypeEmApply)
			if _, err = applyModel.WhereId(applyId).TxUpdate(tx, map[dao.TableField]any{
				model.SocialFriendApply_ApplyLastMsg: in.ApplyMsg,
				model.SocialFriendApply_ApplyLastAt:  time.Now().Unix(),
				model.SocialFriendApply_IsRead:       0,
				model.SocialFriendApply_Content:      newContent,
				model.SocialFriend_StateEm:           constd.SocialFriendStateEmApply,
				model.SocialFriend_SourceEm:          in.SourceEm,
			}); err != nil {
				return l.resd.Error(err)
			}

		} else {
			//不存在申请，创建新申请
			applyId = utild.MakeId()
			newContent := biz.AddApplyRecord("", l.req.UserId, l.req.ApplyMsg, l.req.SourceEm, constd.SocialFriendApplyRecordTypeEmApply)
			if _, err = applyModel.TxInsert(tx, &model.SocialFriendApply{
				Id:           applyId,
				UserId:       l.req.UserId,
				FriendUid:    l.req.FriendUid,
				ApplyLastMsg: l.req.ApplyMsg,
				ApplyLastAt:  time.Now().Unix(),
				ApplyStartAt: time.Now().Unix(),
				StateEm:      constd.SocialFriendStateEmApply,
				SourceEm:     l.req.SourceEm,
				Content: sql.NullString{
					String: newContent,
					Valid:  true,
				},
			}); err != nil {
				return l.resd.Error(err)
			}

		}
		return nil
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	sendTime := utild.NowTime()
	msgClasEm := int64(constd.MsgClasEmFriendApplyNew)
	asyncCtx := contextx.ValueOnlyFrom(l.ctx)
	threading.GoSafe(func() {
		_, err := l.svc.ImRpc.SendSysMsg(asyncCtx, &imRpc.SendSysMsgReq{
			UserId:    &l.req.FriendUid,
			SendTime:  &sendTime,
			MsgClasEm: &msgClasEm,
		})
		if err != nil {
			resd.ErrorCtx(asyncCtx, err)
		}
	})
	return &socialRpc.CreateFriendApplyResp{
		ApplyId: applyId,
	}, nil
}
func (l *CreateFriendApplyLogic) checkReqParams() error {
	if l.req.FriendUid == l.req.UserId {
		return l.resd.NewErr(resd.ErrSocialNotAddSelf)
	}
	return nil
}
