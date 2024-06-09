package logic

import (
	"context"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/pb"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/pkg/numd"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUpUserConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUpUserConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUpUserConversationLogic {
	return &SetUpUserConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SetUpUserConversation 建立会话: 群聊, 私聊
func (l *SetUpUserConversationLogic) SetUpUserConversation(in *pb.SetUpUserConversationReq) (*pb.SetUpUserConversationResp, error) {
	switch websocketd.ChatType(in.ChatType) {
	case websocketd.SingleChatType:
		//生成会话id
		conversationId := numd.CombineInt64(in.SendId, in.RecvId)
		logx.Info(conversationId)
		//验证是否建立过会话
		conversationRes, err := l.svcCtx.ConversationModel.FindByCode(l.ctx, conversationId)
		if conversationRes != nil {
			//有查到数据，已经建立过了
			return nil, nil
		} else if err != nil && err == modelMongo.ErrNotFound {
			//没查到数据，建立会话
			err := l.svcCtx.ConversationModel.Insert(l.ctx, &modelMongo.Conversation{
				ConversationId: conversationId,
				ChatType:       websocketd.SingleChatType,
			})
			if err != nil {
				return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
			}
		} else {
			//查询报错
			return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
		}
		//建立两者关系的会话
		err = l.setupUserConversation(conversationId, in.SendId, in.RecvId, websocketd.ChatType(in.ChatType), true)
		if err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
		}
		//对接收方来说还不需要展开会话
		err = l.setupUserConversation(conversationId, in.RecvId, in.SendId, websocketd.ChatType(in.ChatType), false)
		if err != nil {
			return nil, resd.NewRpcErrCtx(l.ctx, err.Error())
		}
	}

	return &pb.SetUpUserConversationResp{}, nil
}

func (l *SetUpUserConversationLogic) setupUserConversation(conversationId string, userId, recvId int64, chatType websocketd.ChatType, isShow bool) error {
	//用户的会话列表
	conversations, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, userId)
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
	}
	//更新
	err = l.svcCtx.ConversationsModel.Save(l.ctx, conversations)
	return err
}
