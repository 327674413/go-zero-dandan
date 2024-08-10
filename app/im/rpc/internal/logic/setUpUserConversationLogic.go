package logic

import (
	"context"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
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
	convData := &modelMongo.Conversation{}
	conversationId := ""
	//获取会话id
	if l.req.ChatType == constd.ChatTypeSingle {
		//私聊根据双方id组装
		conversationId = utild.CombineId(l.req.SendId, l.req.RecvId)
	} else if l.req.ChatType == constd.ChatTypeGroup {
		//群聊就是群id，也就是接收消息方id
		conversationId = l.req.RecvId
	} else {
		return nil, l.resd.NewErrWithTemp(resd.ErrReqParamFormat1, "chatType")
	}
	//验证是否建立过会话
	conversationRes, err := l.svc.ConversationModel.FindByConvId(l.ctx, conversationId)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	if conversationRes != nil {
		//有查到数据，已经建立过了，更新消息序号、窗口信息
		convData = conversationRes
		convData.ReadSeq = convData.Total
		convData.LastAtMs = utild.GetTimeMs()
		//更新操作人的会话
		_, err = l.svc.ConversationModel.Update(l.ctx, convData)
		if err != nil {
			return nil, l.resd.Error(err)
		}
		//更新操作人的会话列表
		err = l.setUpConversationList(convData, l.req.SendId, l.req.RecvId, websocketd.ChatType(l.req.ChatType), true)
		if err != nil {
			return nil, l.resd.Error(err)
		}
		return &imRpc.SetUpUserConversationResp{}, nil
	} else {
		//没查到数据，建立会话
		convData = &modelMongo.Conversation{
			ConversationId: conversationId,
			ChatType:       websocketd.ChatTypeSingle,
			LastAtMs:       utild.GetTimeMs(),
			PlatId:         l.meta.PlatId,
		}
		err := l.svc.ConversationModel.Insert(l.ctx, convData)
		if err != nil {
			return nil, l.resd.Error(err)
		}
	}
	if l.req.ChatType == constd.ChatTypeSingle {
		//建立两者关系的会话
		err = l.setUpConversationList(convData, l.req.SendId, l.req.RecvId, websocketd.ChatType(l.req.ChatType), true)
		if err != nil {
			return nil, l.resd.Error(err)
		}
		//对接收方来说还不需要展开会话
		err = l.setUpConversationList(convData, l.req.RecvId, l.req.SendId, websocketd.ChatType(l.req.ChatType), false)
		if err != nil {
			return nil, l.resd.Error(err)
		}
	} else if l.req.ChatType == constd.ChatTypeGroup {
		//群聊则应该肯定是有会话的，直接处理激活好了
		err := l.setUpConversationList(convData, l.req.SendId, l.req.RecvId, websocketd.ChatType(l.req.ChatType), true)
		if err != nil {
			return nil, l.resd.Error(err)
		}
	}
	return &imRpc.SetUpUserConversationResp{}, nil
}

func (l *SetUpUserConversationLogic) setUpConversationList(convData *modelMongo.Conversation, userId, recvId string, chatType websocketd.ChatType, isShow bool) error {
	//用户的会话列表
	conversations, err := l.svc.ConversationsModel.FindByUserId(l.ctx, userId)
	if err != nil {
		return l.resd.Error(err)
	}
	if conversations == nil {
		//没有会话列表，建立会话
		conversations = &modelMongo.Conversations{
			ID:               primitive.NewObjectID(),
			UserId:           userId,
			ConversationList: make(map[string]*modelMongo.Conversation),
		}
	}
	convData.IsShow = isShow
	convData.TargetId = recvId
	convData.UserId = userId
	//既然用save了，应该不用这里了？
	//更新会话记录
	//if _, ok := conversations.ConversationList[convData.ConversationId]; ok {
	//	conversations.ConversationList[convData.ConversationId] = convData
	//	_, err = l.svc.ConversationsModel.Update(l.ctx, conversations)
	//	if err != nil {
	//		return l.resd.Error(err)
	//	}
	//	return nil
	//}
	//添加会话记录
	conversations.ConversationList[convData.ConversationId] = convData
	//更新
	err = l.svc.ConversationsModel.Save(l.ctx, conversations)
	if err != nil {
		return l.resd.Error(err)
	}
	return nil
}
