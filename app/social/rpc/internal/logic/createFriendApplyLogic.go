package logic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/logx"
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
	ctx context.Context
	svc *svc.ServiceContext
	logx.Logger
}

func NewCreateFriendApplyLogic(ctx context.Context, svc *svc.ServiceContext) *CreateFriendApplyLogic {
	return &CreateFriendApplyLogic{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFriendApplyLogic) CreateFriendApply(in *socialRpc.CreateFriendApplyReq) (*socialRpc.CreateFriendApplyResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	friendModel := model.NewSocialFriendModel(l.ctx, l.svc.SqlConn, in.PlatId)
	//我申请别人为好友，UserId是我，FriendUid是好友
	//查询我-对方好友数据
	existSelf, err := friendModel.Where("user_id = ? and friend_uid = ?", in.UserId, in.FriendUid).Find()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	//查询对方-我好友数据
	existFriend, err := friendModel.Where("friend_uid = ? and user_id = ?", in.UserId, in.FriendUid).Find()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	//分析我-对方的关系
	if existSelf != nil {
		//如果我-对方是好友 并且 对方-我也是好友
		if existSelf.StateEm == constd.SocialFriendStateEmPass && existFriend != nil && existFriend.StateEm == constd.SocialFriendStateEmPass {
			return nil, resd.NewRpcErrCtx(l.ctx, "已经是好友", resd.SocialAlreadyFriend)
		}
	}
	//分析对方-我的关系
	if existFriend != nil {
		//对方拉黑了我，不允许申请
		if existFriend.StateEm == constd.SocialFriendStateEmBlack {
			return nil, resd.NewRpcErrCtx(l.ctx, "对方把你拉黑", resd.SocialAlreadyBlackMe)
		}
		//对方还是我的好友，而我不是对方好友
		if existFriend.StateEm == constd.SocialFriendStateEmPass {
			//先当作还是要对方点同意来弄，所以先不用处理
		}
	}
	//基础校验完成，可以添加，开启事务
	tx, err := dao.StartTrans(l.svc.SqlConn, l.ctx)
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlStartTransErr)
	}
	//确保有我 - 对方的关系
	if existSelf == nil {
		// 不存在，创建
		if _, err = friendModel.TxInsert(tx, &model.SocialFriend{
			Id:        utild.MakeId(),
			UserId:    in.UserId,
			FriendUid: in.FriendUid,
			StateEm:   constd.SocialFriendStateEmApply,
		}); err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, "创建我的好友申请失败", resd.MysqlInsertErr)
		}
	} else {
		//存在，更新我 - 对方的状态
		if _, err = friendModel.WhereId(existSelf.Id).TxUpdate(tx, map[dao.TableField]any{
			model.SocialFriend_StateEm: constd.SocialFriendStateEmApply,
		}); err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, "更新我的好友表失败", resd.MysqlUpdateErr)
		}
	}
	//确保有对方 - 我的关系
	if existFriend == nil {
		//不存在，创建
		if _, err = friendModel.TxInsert(tx, &model.SocialFriend{
			Id:        utild.MakeId(),
			UserId:    in.FriendUid,
			FriendUid: in.UserId,
			StateEm:   constd.SocialFriendStateEmApply,
		}); err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, "创建对方申请失败", resd.MysqlInsertErr)
		}
	} else {
		//存在，更新 对我 - 我的状态
		if _, err = friendModel.WhereId(existFriend.Id).TxUpdate(tx, map[dao.TableField]any{
			model.SocialFriend_StateEm: constd.SocialFriendStateEmApply,
		}); err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, "更新对方的好友表失败", resd.MysqlUpdateErr)
		}
	}

	//查看是否存在过申请
	applyModel := model.NewSocialFriendApplyModel(l.ctx, l.svc.SqlConn, in.PlatId)
	existApply, err := applyModel.Where("friend_uid = ? AND user_id = ?", in.FriendUid, in.UserId).Find()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	var applyId string
	if existApply != nil {
		//存在申请，更新与添加申请消息
		applyId = existApply.Id
		newContent := biz.AddApplyRecord(existApply.Content.String, in.UserId, in.ApplyMsg, in.SourceEm, constd.SocialFriendApplyRecordTypeEmApply)
		if _, err = applyModel.WhereId(applyId).TxUpdate(tx, map[dao.TableField]any{
			model.SocialFriendApply_ApplyLastMsg: in.ApplyMsg,
			model.SocialFriendApply_ApplyLastAt:  time.Now().Unix(),
			model.SocialFriendApply_IsRead:       0,
			model.SocialFriendApply_Content:      newContent,
			model.SocialFriend_StateEm:           constd.SocialFriendStateEmApply,
			model.SocialFriend_SourceEm:          in.SourceEm,
		}); err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, "添加好友申请失败", resd.MysqlUpdateErr)
		}

	} else {
		//不存在申请，创建新申请
		applyId = utild.MakeId()
		newContent := biz.AddApplyRecord("", in.UserId, in.ApplyMsg, in.SourceEm, constd.SocialFriendApplyRecordTypeEmApply)
		if _, err = applyModel.TxInsert(tx, &model.SocialFriendApply{
			Id:           applyId,
			UserId:       in.UserId,
			FriendUid:    in.FriendUid,
			ApplyLastMsg: in.ApplyMsg,
			ApplyLastAt:  time.Now().Unix(),
			ApplyStartAt: time.Now().Unix(),
			StateEm:      constd.SocialFriendStateEmApply,
			SourceEm:     in.SourceEm,
			Content: sql.NullString{
				String: newContent,
				Valid:  true,
			},
		}); err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, "创建好友申请失败", resd.MysqlInsertErr)
		}

	}
	if err := tx.Commit(); err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlCommitErr)
	}
	asyncCtx := contextx.ValueOnlyFrom(l.ctx)
	threading.GoSafe(func() {
		_, err := l.svc.ImRpc.SendSysMsg(asyncCtx, &imRpc.SendSysMsgReq{
			PlatId:    in.PlatId,
			UserId:    in.FriendUid,
			SendTime:  utild.NowTime(),
			MsgClasEm: constd.MsgClasEmFriendApplyNew,
		})
		if err != nil {
			logx.Error(err)
		} else {
			logx.Debug("推送新好友通知")
		}
	})
	return &socialRpc.CreateFriendApplyResp{
		ApplyId: applyId,
	}, nil
}
func (l *CreateFriendApplyLogic) checkReqParams(in *socialRpc.CreateFriendApplyReq) error {
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
