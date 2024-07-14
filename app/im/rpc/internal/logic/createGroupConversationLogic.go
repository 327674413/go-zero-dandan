package logic

import (
	"context"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/resd"
)

type CreateGroupConversationLogic struct {
	*CreateGroupConversationLogicGen
}

func NewCreateGroupConversationLogic(ctx context.Context, svc *svc.ServiceContext) *CreateGroupConversationLogic {
	return &CreateGroupConversationLogic{
		CreateGroupConversationLogicGen: NewCreateGroupConversationLogicGen(ctx, svc),
	}
}

func (l *CreateGroupConversationLogic) CreateGroupConversation(in *imRpc.CreateGroupConversationReq) (*imRpc.CreateGroupConversationResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	res := &imRpc.CreateGroupConversationResp{}
	_, err := l.svc.ConversationModel.FindByCode(l.ctx, l.req.GroupId)
	//未报错则有数据，已存在
	if err == nil {
		return res, nil
	}
	//不是未找到，是报错，则返回报错
	if err != modelMongo.ErrNotFound {
		return nil, l.resd.Error(err)
	}
	//未找到，创建
	err = l.svc.ConversationModel.Insert(l.ctx, &modelMongo.Conversation{
		ConversationId: l.req.GroupId,
		ChatType:       websocketd.ChatTypeGroup,
	})
	chatType := int64(websocketd.ChatTypeGroup)
	//创建群后，需要创建用户的会话列表
	_, err = NewSetUpUserConversationLogic(l.ctx, l.svc).SetUpUserConversation(&imRpc.SetUpUserConversationReq{
		SendId:   &l.req.CreateId,
		RecvId:   &l.req.GroupId,
		ChatType: &chatType,
	})
	if err != nil {
		return nil, l.resd.NewErr(resd.ErrCreateConversation)
	}
	return res, nil
}
