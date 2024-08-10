package logic

import (
	"context"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
)

type DeleteUserConversationLogic struct {
	*DeleteUserConversationLogicGen
}

func NewDeleteUserConversationLogic(ctx context.Context, svc *svc.ServiceContext) *DeleteUserConversationLogic {
	return &DeleteUserConversationLogic{
		DeleteUserConversationLogicGen: NewDeleteUserConversationLogicGen(ctx, svc),
	}
}

func (l *DeleteUserConversationLogic) DeleteUserConversation(in *imRpc.DeleteUserConversationReq) (*imRpc.ResultResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	userConv, err := l.svc.ConversationsModel.FindByUserId(l.ctx, l.req.UserId)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	if userConv == nil {
		return &imRpc.ResultResp{
			Code:    1,
			Content: "用户暂无会话列表",
		}, nil
	}
	if _, ok := userConv.ConversationList[l.req.ConversationId]; ok {
		userConv.ConversationList[l.req.ConversationId].DeleteSeq = userConv.ConversationList[l.req.ConversationId].Total
		userConv.ConversationList[l.req.ConversationId].IsShow = false
		if err = l.svc.ConversationsModel.Save(l.ctx, userConv); err != nil {
			return nil, l.resd.Error(err)
		}
	} else {
		return &imRpc.ResultResp{
			Code:    2,
			Content: "用户会话列表总没有该会话",
		}, nil
	}
	return &imRpc.ResultResp{}, nil
}
