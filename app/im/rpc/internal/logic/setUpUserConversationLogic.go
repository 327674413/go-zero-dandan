package logic

import (
	"context"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/utild"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SetUpUserConversationLogic struct {
	*SetUpUserConversationLogicGen
}

func NewSetUpUserConversationLogic(ctx context.Context, svc *svc.ServiceContext) *SetUpUserConversationLogic {
	return &SetUpUserConversationLogic{
		SetUpUserConversationLogicGen: NewSetUpUserConversationLogicGen(ctx, svc),
	}
}

// SetUpUserConversation 建立会话: 群聊, 私聊
func (l *SetUpUserConversationLogic) SetUpUserConversation(in *imRpc.SetUpUserConversationReq) (*imRpc.SetUpUserConversationResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	switch websocketd.ChatType(l.req.ChatType) {
	case websocketd.ChatTypeSingle:
		//生成会话id
		conversationId := utild.CombineId(l.req.SendId, l.req.RecvId)
		//验证是否建立过会话
		conversationRes, err := l.svc.ConversationModel.FindByCode(l.ctx, conversationId)
		if conversationRes != nil {
			//有查到数据，已经建立过了
			return &imRpc.SetUpUserConversationResp{}, nil
		} else if err != nil && err == modelMongo.ErrNotFound {
			//没查到数据，建立会话
			err := l.svc.ConversationModel.Insert(l.ctx, &modelMongo.Conversation{
				ConversationId: conversationId,
				ChatType:       websocketd.ChatTypeSingle,
				LastAt:         utild.GetStamp(),
				PlatId:         l.meta.PlatId,
			})
			if err != nil {
				return nil, l.resd.Error(err)
			}
		} else {
			//查询报错
			return nil, l.resd.Error(err)
		}
		//建立两者关系的会话
		err = l.setupUserConversation(conversationId, l.req.SendId, l.req.RecvId, websocketd.ChatType(l.req.ChatType), true)
		if err != nil {
			return nil, l.resd.Error(err)
		}
		//对接收方来说还不需要展开会话
		err = l.setupUserConversation(conversationId, l.req.RecvId, l.req.SendId, websocketd.ChatType(l.req.ChatType), false)
		if err != nil {
			return nil, l.resd.Error(err)
		}
	case websocketd.ChatTypeGroup:
		err := l.setupUserConversation(l.req.RecvId, l.req.SendId, l.req.RecvId, websocketd.ChatType(l.req.ChatType), true)
		if err != nil {
			return nil, l.resd.Error(err)
		}
	}
	return &imRpc.SetUpUserConversationResp{}, nil
}

func (l *SetUpUserConversationLogic) setupUserConversation(conversationId string, userId, recvId string, chatType websocketd.ChatType, isShow bool) error {
	//用户的会话列表
	conversations, err := l.svc.ConversationsModel.FindByUserId(l.ctx, userId)
	if err != nil {
		if err == modelMongo.ErrNotFound {
			//没有会话列表，建立会话
			conversations = &modelMongo.Conversations{
				ID:               primitive.NewObjectID(),
				UserId:           userId,
				ConversationList: make(map[string]*modelMongo.Conversation),
			}
		} else {
			return err
		}
	}
	//更新会话记录
	if _, ok := conversations.ConversationList[conversationId]; ok {
		return nil
	}
	//添加会话记录
	conversations.ConversationList[conversationId] = &modelMongo.Conversation{
		ConversationId: conversationId,
		ChatType:       chatType,
		IsShow:         isShow,
		TargetId:       recvId,
		LastAt:         utild.GetStamp(),
	}
	//更新
	err = l.svc.ConversationsModel.Save(l.ctx, conversations)
	if err != nil {
		return l.resd.Error(err)
	}
	return nil
}
