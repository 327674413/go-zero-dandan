package msgTransfer

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/ws/websocketd"
	"sync"
	"time"
)

type groupMsgRead struct {
	mu             sync.Mutex
	push           *websocketd.Push //用户推送的消息
	pushCh         chan *websocketd.Push
	count          int
	conversationId string
	pushTime       time.Time //上次推送时间
	done           chan struct{}
}

func newGroupMsgRead(push *websocketd.Push, pushCh chan *websocketd.Push) *groupMsgRead {
	m := &groupMsgRead{
		conversationId: push.ConversationId,
		push:           push,
		pushCh:         pushCh,
		count:          1,
		pushTime:       time.Now(),
		done:           make(chan struct{}),
	}
	go m.transfer()
	return m
}
func (t *groupMsgRead) transfer() {
	//超时发送
	timer := time.NewTimer(GroupMsgReadRecordDelayTime / 2)
	defer timer.Stop()
	for {
		select {
		case <-t.done:
			logx.Info("触发结束了")
			return
		case <-timer.C:
			logx.Info("延迟发送定时器到时间了")
			t.mu.Lock()
			pushTime := t.pushTime
			val := GroupMsgReadRecordDelayTime - time.Since(pushTime)
			push := t.push
			if val > 0 && t.count < GroupMsgReadRecordDelayCount || push == nil {
				if val > 0 {
					timer.Reset(val)
				}
				t.mu.Unlock()
				continue
			}
			t.pushTime = time.Now()
			t.push = nil
			t.count = 0
			timer.Reset(GroupMsgReadRecordDelayTime / 2)
			t.mu.Unlock()
			//推送
			logx.Infof("超过合并延迟时间，发送：%v", push.ConversationId)
			t.pushCh <- push
		default:
			//超量发送
			t.mu.Lock()
			if t.count >= GroupMsgReadRecordDelayCount {
				push := t.push
				t.push = nil
				t.count = 0
				t.mu.Unlock()
				logx.Infof("超过合并条数，发送：%v", push.ConversationId)
				t.pushCh <- push
				continue
			}
			if t.IsIdle() {
				logx.Info("延迟发送器是空闲状态，释放掉")
				t.mu.Unlock()
				// 使得msgReadTransfer释放
				t.pushCh <- &websocketd.Push{
					ChatType:       websocketd.ChatTypeGroup,
					ConversationId: t.conversationId,
				}
				continue
			}
			t.mu.Unlock()
			tempDelay := GroupMsgReadRecordDelayTime / 4
			if tempDelay > time.Second {
				tempDelay = time.Second
			}
			logx.Infof("继续延迟%v", tempDelay)
			time.Sleep(tempDelay)
		}
	}

}

// 合并推送，同个消息id的就合并了
func (t *groupMsgRead) mergePush(push *websocketd.Push) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.count++
	for msgId, read := range push.MsgReads {
		t.push.MsgReads[msgId] = read
	}
}

// IsIdleWithMuLock 判断是否为活跃状态，会锁mu
func (t *groupMsgRead) IsIdleWithMuLock() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.IsIdle()
}

// IsIdle 判断是否为活跃状态
func (t *groupMsgRead) IsIdle() bool {
	pushTime := t.pushTime
	val := GroupMsgReadRecordDelayTime*2 - time.Since(pushTime) //可能阻塞还没来得及处理，多给一些时间所以乘以2
	if val <= 0 && t.push == nil && t.count == 0 {
		return true
	}
	return false
}

// clear 清理方法
func (t *groupMsgRead) clear() {
	select {
	case <-t.done: //好像是优先用done方法，没有就进默认的关闭
	default:
		close(t.done)
	}
	t.push = nil
}
