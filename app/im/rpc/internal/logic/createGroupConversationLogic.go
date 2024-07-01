package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/rpc/internal/svc"
)

type CreateGroupConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupConversationLogic {
	return &CreateGroupConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupConversationLogic) CreateGroupConversation(in *imRpc.CreateGroupConversationReq) (*imRpc.CreateGroupConversationResp, error) {
	res := &imRpc.CreateGroupConversationResp{}
	_, err := l.svcCtx.ConversationModel.FindByCode(l.ctx, in.GroupId)
	//未报错则有数据，已存在
	if err == nil {
		return res, nil
	}
	//不是未找到，是报错，则返回报错
	if err != modelMongo.ErrNotFound {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
	}
	//未找到，创建
	err = l.svcCtx.ConversationModel.Insert(l.ctx, &modelMongo.Conversation{
		ConversationId: in.GroupId,
		ChatType:       websocketd.ChatTypeGroup,
	})
	//创建群后，需要创建用户的会话列表
	_, err = NewSetUpUserConversationLogic(l.ctx, l.svcCtx).SetUpUserConversation(&imRpc.SetUpUserConversationReq{
		SendId:   in.CreateId,
		RecvId:   in.GroupId,
		ChatType: int64(websocketd.ChatTypeGroup),
	})
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, fmt.Sprintf("创建会话失败：%v", err))
	}
	return res, nil
}
