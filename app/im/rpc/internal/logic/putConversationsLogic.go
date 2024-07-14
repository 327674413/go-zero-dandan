package logic

import (
	"context"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/im/ws/websocketd"
)

type PutConversationsLogic struct {
	*PutConversationsLogicGen
}

func NewPutConversationsLogic(ctx context.Context, svc *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		PutConversationsLogicGen: NewPutConversationsLogicGen(ctx, svc),
	}
}

// PutConversations 更新会话
func (l *PutConversationsLogic) PutConversations(in *imRpc.PutConversationsReq) (*imRpc.PutConversationsResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	//获取用户的会话列表
	data, err := l.svc.ConversationsModel.FindByUserId(l.ctx, l.req.UserId)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	//如果用户无会话列表，初始化创建一个空列表
	if data.ConversationList == nil {
		data.ConversationList = make(map[string]*modelMongo.Conversation)
	}
	//遍历需要更新的会话列表
	for k, v := range in.ConversationList {
		var oldTotal int64
		//如果已经存在原先的会话列表，则用原先的会话列表的total值
		if data.ConversationList[k] != nil {
			oldTotal = data.ConversationList[k].Total
		}
		//更新会话信息
		data.ConversationList[k] = &modelMongo.Conversation{
			ConversationId: v.ConversationId,
			ChatType:       websocketd.ChatType(v.ChatType),
			IsShow:         v.IsShow,
			Total:          v.Read + oldTotal, //会话总数 = 本次读数 + 原先的会话总数
			Seq:            v.Seq,
		}
	}
	_, err = l.svc.ConversationsModel.Update(l.ctx, data)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	return &imRpc.PutConversationsResp{}, nil
}
