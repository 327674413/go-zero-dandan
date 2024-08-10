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
	conv, err := l.svc.ConversationModel.FindByConvId(l.ctx, l.req.GroupId)
	//有报错
	if err == nil {
		return nil, l.resd.Error(err)
	}
	//没报错且有数据
	if conv != nil {
		return res, nil
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
