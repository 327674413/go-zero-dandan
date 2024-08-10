package logic

import (
	"context"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"
)

type GetConversationsLogic struct {
	*GetConversationsLogicGen
}

func NewGetConversationsLogic(ctx context.Context, svc *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		GetConversationsLogicGen: NewGetConversationsLogicGen(ctx, svc),
	}
}

// GetConversations 获取会话
func (l *GetConversationsLogic) GetConversations(in *imRpc.GetConversationsReq) (*imRpc.GetConversationsResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	//在用户会话关系表中，查询用户所有会话
	data, err := l.svc.ConversationsModel.FindByUserId(l.ctx, l.req.UserId)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	if data == nil {
		//没有会话记录
		return &imRpc.GetConversationsResp{}, nil
	}
	res := &imRpc.GetConversationsResp{}
	err = copier.Copy(&res, data)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	ids := make([]string, 0, len(data.ConversationList))
	//获取所有会话id
	for _, v := range data.ConversationList {
		ids = append(ids, v.ConversationId)
	}
	// 通过会话ID列表，获取会话详细信息
	conversations, err := l.svc.ConversationModel.ListByConversationIds(l.ctx, ids)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	err = copier.Copy(&res.ConversationList, conversations)
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrCopier)
	}
	//循环会话，根据会话详情判断消息状态
	for _, item := range conversations {
		//如果会话详情的id不在用户的会话列表中，跳过（讲道理不会触发啊？根据关系id查到的总集合，怎会不在？）
		if _, ok := res.ConversationList[item.ConversationId]; !ok {
			continue
		}
		if item.LastMsg != nil {
			res.ConversationList[item.ConversationId].LastMsg = &imRpc.ChatLog{
				Id:             item.LastMsg.ID.Hex(),
				ConversationId: item.LastMsg.ConversationId,
				SendId:         item.LastMsg.SendId,
				RecvId:         item.LastMsg.RecvId,
				MsgType:        int64(item.LastMsg.MsgType),
				MsgContent:     item.LastMsg.MsgContent,
				ChatType:       int64(item.LastMsg.ChatType),
				SendTime:       item.LastMsg.SendTime,
				ReadRecords:    item.LastMsg.ReadRecords,
			}
		}

		//用户会话关系中的消息总数
		total := res.ConversationList[item.ConversationId].Total
		//判断会话关系中消息总数 小于 会话详情中总数，则代表该会话有新消息
		if total < item.Total {
			//更新返回会话列表的消息总数，用会话详情的总数
			res.ConversationList[item.ConversationId].Total = item.Total
			//会话详情总数 - 用户会话关系中的总数，就是未读消息数
			res.ConversationList[item.ConversationId].Unread = item.Total - total
			//更改当前会话为显示状态（有可能用户会把会话删除，删除了就是false）
			res.ConversationList[item.ConversationId].IsShow = true
		}
	}
	return res, nil
}
