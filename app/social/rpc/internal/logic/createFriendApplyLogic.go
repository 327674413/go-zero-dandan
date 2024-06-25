package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/social/model"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/pb"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"time"
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

type applyContent struct {
	UserId string `json:"userId"`
	TimeAt int64  `json:"timeAt"`
	Text   string `json:"text"`
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

	//查看是否存在过申请
	applyModel := model.NewSocialFriendApplyModel(l.svcCtx.SqlConn, in.PlatId)
	existApply, err := applyModel.Where("friend_uid = ? AND user_id = ?", in.FriendUid, in.UserId).Find()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	var applyId string
	if existApply != nil {
		//存在申请，更新与添加申请消息
		applyId = existApply.Id
		msgList := make([]*applyContent, 0)
		if existApply.Content.String != "" {
			_ = json.Unmarshal([]byte(existApply.Content.String), &msgList)
		}
		msgList = append([]*applyContent{
			{
				UserId: in.UserId,
				TimeAt: time.Now().Unix(),
				Text:   in.ApplyMsg,
			},
		}, msgList...)
		newContent, _ := json.Marshal(msgList)
		_, err = applyModel.WhereId(applyId).Update(map[dao.TableField]any{
			model.SocialFriendApply_ApplyLastMsg: in.ApplyMsg,
			model.SocialFriendApply_ApplyLastAt:  time.Now().Unix(),
			model.SocialFriendApply_IsRead:       0,
			model.SocialFriendApply_Content:      string(newContent),
		})
		if err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, "添加好友申请失败", resd.MysqlUpdateErr)
		}
	} else {
		//不存在申请，创建新申请
		applyId = utild.MakeId()
		msgList := []*applyContent{
			{
				UserId: in.UserId,
				TimeAt: time.Now().Unix(),
				Text:   in.ApplyMsg,
			},
		}
		newContent, _ := json.Marshal(msgList)
		_, err = applyModel.Insert(&model.SocialFriendApply{
			Id:           applyId,
			UserId:       in.UserId,
			FriendUid:    in.FriendUid,
			ApplyLastMsg: in.ApplyMsg,
			ApplyLastAt:  time.Now().Unix(),
			ApplyStartAt: time.Now().Unix(),
			Content: sql.NullString{
				String: string(newContent),
				Valid:  true,
			},
		})
		if err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, "创建好友申请失败", resd.MysqlInsertErr)
		}
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
