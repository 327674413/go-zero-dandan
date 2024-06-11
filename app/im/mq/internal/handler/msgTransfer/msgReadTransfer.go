package msgTransfer

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/mq/internal/svc"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/pkg/bitmapd"
	"sync"
	"time"
)

var (
	GroupMsgReadRecordDelayTime  = time.Second
	GroupMsgReadRecordDelayCount = 10
)

const (
	GroupMsgReadHandlerAtTransfer = iota
	GroupMsgReadHandlerDelayTransfer
)

type MsgReadTransfer struct {
	*baseMsgTransfer                          //基础消息转化器
	mu               sync.Mutex               //读写锁
	groupMsgs        map[string]*groupMsgRead //存储群消息处理
	push             chan *websocketd.Push    //消息推送通道

}

func NewMsgReadTransfer(svcCtx *svc.ServiceContext) kq.ConsumeHandler {
	if svcCtx.Config.MsgReadHandler.GroupMsgReadHandler != GroupMsgReadHandlerAtTransfer {
		if svcCtx.Config.MsgReadHandler.GroupMsgReadRecordDelayCount > 0 {
			GroupMsgReadRecordDelayCount = svcCtx.Config.MsgReadHandler.GroupMsgReadRecordDelayCount
		}
		if svcCtx.Config.MsgReadHandler.GroupMsgReadRecordDelayTime > 0 {
			GroupMsgReadRecordDelayTime = time.Duration(svcCtx.Config.MsgReadHandler.GroupMsgReadRecordDelayTime) * time.Second
		}
	}

	m := &MsgReadTransfer{
		baseMsgTransfer: NewBaseMsgTransfer(svcCtx),
		groupMsgs:       make(map[string]*groupMsgRead),
		push:            make(chan *websocketd.Push, 1),
	}
	go m.transfer() //启动专门的读取状态通知协程
	return m
}
func (t *MsgReadTransfer) Consume(key, value string) error {
	t.Info("MsgReadTransfer", value)
	var (
		data *kafkad.MsgMarkRead
		ctx  = context.Background()
	)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	//业务处理
	readRecords, err := t.UpdateChatLogRead(ctx, data)
	if err != nil {
		return err
	}
	logx.Infof("消息的recvId:%v", data.RecvId)
	push := &websocketd.Push{
		ConversationId: data.ConversationId,
		ChatType:       data.ChatType,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ContentType:    websocketd.ContentMakeRead,
		ReadRecords:    readRecords,
	}
	switch data.ChatType {
	case websocketd.SingleChatType:
		logx.Infof("私聊消息发送%v：", push.ConversationId)
		//私聊，直接推送
		t.push <- push
	case websocketd.GroupChatType:
		logx.Infof("群聊消息发送%v：", push.ConversationId)
		// 判断是否开合并推送
		if t.svc.Config.MsgReadHandler.GroupMsgReadHandler == GroupMsgReadHandlerAtTransfer {
			logx.Infof("没开启延迟发送，直接发%v：", push.ConversationId)
			//未开启，直接发送
			t.push <- push
		} else {
			logx.Infof("开启了延迟发送%v：", push.ConversationId)
			//这里等超时后就被锁住了，没有执行下去
			t.mu.Lock()
			logx.Infof("没有锁住，处理中%v：", push.ConversationId)
			defer t.mu.Unlock()
			if _, ok := t.groupMsgs[push.ConversationId]; ok {
				logx.Infof("合并消息%v：", push.ConversationId)
				t.groupMsgs[push.ConversationId].mergePush(push)
			} else {
				logx.Infof("创建消息延迟队列%v：", push.ConversationId)
				t.groupMsgs[push.ConversationId] = newGroupMsgRead(push, t.push)
			}
		}
	}
	return nil
}

func (t *MsgReadTransfer) UpdateChatLogRead(ctx context.Context, data *kafkad.MsgMarkRead) (map[string]string, error) {
	res := make(map[string]string)
	chatLogs, err := t.svc.ChatLogModel.ListByMsgIds(ctx, data.MsgIds)
	if err != nil {
		return nil, err
	}
	// 处理已读
	for _, chatLog := range chatLogs {
		switch chatLog.ChatType {
		case websocketd.SingleChatType:
			chatLog.ReadRecords = []byte{1}
		case websocketd.GroupChatType:
			readRecords := bitmapd.Load(chatLog.ReadRecords)
			readRecords.SetId(data.SendId)
			chatLog.ReadRecords = readRecords.Export()
		}
		//为了保证精度，转成base64返回给前端判断已读状态
		res[chatLog.ID.Hex()] = base64.StdEncoding.EncodeToString(chatLog.ReadRecords)
		err = t.svc.ChatLogModel.UpdateMakeRead(ctx, chatLog.ID, chatLog.ReadRecords)
		if err != nil {
			return nil, err
		}
	}
	return res, err
}

func (t *MsgReadTransfer) transfer() {
	for push := range t.push {
		logx.Infof("消息转化：推送进来了%v", push.ConversationId)
		if push.RecvId != 0 || len(push.RecvIds) > 0 {
			// 问题排查到这里了，好像消息是发了，但是没收到，应该在这里！！！！！！再试一下非群聊是不是能收到？？？？？
			if err := t.Transfer(context.Background(), push); err != nil {
				logx.Errorf("推送消息失败：%v", err)
			}
		}
		logx.Infof("消息转化：没有发送对象%v", push.ConversationId)
		//私聊
		if push.ChatType == websocketd.SingleChatType {
			logx.Info("消息转化：私聊世界结束")
			continue
		}
		//及时处理
		if t.svc.Config.MsgReadHandler.GroupMsgReadHandler == GroupMsgReadHandlerAtTransfer {
			logx.Info("消息转化：即使发送直接结束")
			continue
		}
		//清空数据
		t.mu.Lock()
		if _, ok := t.groupMsgs[push.ConversationId]; ok && t.groupMsgs[push.ConversationId].IsIdleWithMuLock() {
			logx.Info("消息转化：空闲，释放掉")
			t.groupMsgs[push.ConversationId].clear()
			delete(t.groupMsgs, push.ConversationId)
		}
		t.mu.Unlock()
	}
}
