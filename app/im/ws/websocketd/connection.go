package websocketd

import (
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"net/http"
	"sync"
	"time"
)

// Conn 定义了一个包含 websocket 连接和相关信息的结构体，一个客户端连接一个conn
type Conn struct {
	Uid               string        //该连接的用户id
	PlatId            string        //所属应用id
	PlatClasEm        int64         //所属应用的应用类型
	PortEm            int64         //连接的端口
	Lang              string        //国际化当前语言
	Resd              *resd.Resp    //国际化统一返回
	*websocket.Conn                 // 嵌入的 websocket 连接
	idleMu            sync.Mutex    // 用于保护共享资源的互斥锁
	s                 *Server       // 指向关联的服务器实例
	idle              time.Time     // 上次活动的时间戳
	maxConnectionIdle time.Duration // 最大允许的连接空闲时间
	done              chan struct{} // 用于通知连接结束的通道
	messageMu         sync.Mutex
	readMessageAckMq  []*Message          // 本客户端需要ack确认消息的队列，保证按顺序确认用
	readMessageSeqMap map[string]*Message //如果是要应答的ack消息，用消息id作为索引管理应答确认秦光
	readMessageChan   chan *Message       // 用来异步处理收到消息的管道，当无需ack确认或ack确认后，丢给该管道处理
}

// NewConn 创建一个新的 websocket 连接并返回 Conn 结构体的实例
func NewConn(s *Server, w http.ResponseWriter, r *http.Request) *Conn {
	var responseHeader http.Header
	//教程学需要加这个才不会出现1006 eof错误，但我自己实际好像没碰到这个问题，先加着
	if protocol := r.Header.Get("Sec-WebSocket-Protocol"); protocol != "" {
		responseHeader = http.Header{
			"Sec-Websocket-Protocol": []string{protocol},
		}
	}

	c, err := s.upgrader.Upgrade(w, r, responseHeader) // 升级 HTTP 连接为 websocket 连接
	if err != nil {                                    // 如果升级失败，记录错误并返回 nil
		s.Errorf("upgrade err %v:", err)
		return nil
	}
	conn := &Conn{ // 初始化并返回新的 Conn 实例
		Conn:              c,                       // 保存 websocket 连接
		s:                 s,                       // 关联服务器实例
		idle:              time.Now(),              // 记录当前时间为上次活动时间
		maxConnectionIdle: s.opt.maxConnectionIdle, // 设置最大连接空闲时间
		done:              make(chan struct{}),     // 初始化通知通道
		Lang:              r.FormValue("lang"),
		Resd:              resd.NewResp(r.Context(), r.FormValue("lang")),
		readMessageAckMq:  make([]*Message, 0, 2),
		readMessageSeqMap: make(map[string]*Message),
		readMessageChan:   make(chan *Message, 1), //给初始容量，防止阻塞，同时1个容量保证顺序
	}
	go conn.keepalive()
	return conn
}

// 写入ack的队列
func (t *Conn) appendMsgMq(msg *Message) {
	t.messageMu.Lock()
	defer t.messageMu.Unlock()
	//根据消息id，判断是否存在ack确认记录
	logx.Info("mq操作开始")
	if m, ok := t.readMessageSeqMap[msg.Id]; ok {
		logx.Info("有该消息的ack记录")
		//存在ack确认记录，先判断当前是否有待发送消息
		if len(t.readMessageAckMq) == 0 {
			//消息队列中没有消息，应该是已经发送了，推出
			logx.Info("确认ack，但队列中没消息，退出")
			return
		}
		// 消息队列中有消息，判断ack的序号
		if m.AckSeq >= msg.AckSeq {
			//过来的消息序号比记录的序号小，代表重复或过时的记录
			return
		}
		// 更新ack确认记录，确保队列中消息的ack确认序号是最新的
		t.readMessageSeqMap[msg.Id] = msg
		return
	}
	//还没有ack确认，先避免客户端重复发送ack确认
	if msg.FrameType == FrameAck {
		logx.Info("没有消息ack待确认，并且当前发的是第二次确认，应该是重复发，过滤")
		return
	}
	//ack确认，将消息放入消息队列数组
	logx.Info("需要ack确认，推入待确认消息列表")
	t.readMessageAckMq = append(t.readMessageAckMq, msg)
	//ack的确认序号中增加该消息id，代表等待确认
	t.readMessageSeqMap[msg.Id] = msg
}

// ReadMessage 读取消息
func (t *Conn) ReadMessage() (messageType int, p []byte, err error) {
	messageType, p, err = t.Conn.ReadMessage()
	//三方库中的读写是不安全的，需要锁
	t.idleMu.Lock()
	defer t.idleMu.Unlock()
	t.idle = time.Now() //设置上次活跃时间为当前时间
	return
}

// WriteMessage 发送消息
func (t *Conn) WriteMessage(messageType int, data []byte) error {
	//三方库中的读写是不安全的，需要锁
	t.idleMu.Lock()
	defer t.idleMu.Unlock()
	err := t.Conn.WriteMessage(messageType, data)
	t.idle = time.Now() //设置上次活跃时间为当前时间
	return err

}

// keepalive 维护连接活跃状态，如果连接空闲超时则关闭连接
func (t *Conn) keepalive() {
	idleTimer := time.NewTimer(t.maxConnectionIdle) // 创建一个定时器，初始时间为最大连接空闲时间
	defer func() {                                  // 确保定时器在函数退出时停止
		idleTimer.Stop()
	}()
	for {
		select {
		case <-idleTimer.C: // 定时器触发
			t.idleMu.Lock()    // 加锁，保护共享资源
			idle := t.idle     // 读取上次活动时间
			if idle.IsZero() { // 如果上次活动时间未设置
				t.idleMu.Unlock()                    // 解锁
				idleTimer.Reset(t.maxConnectionIdle) // 重置定时器
				continue                             // 继续循环
			}
			val := t.maxConnectionIdle - time.Since(idle) // 计算剩余空闲时间
			t.idleMu.Unlock()                             // 解锁
			if val <= 0 {                                 // 如果剩余时间小于等于零
				t.s.Close(t) // 关闭连接
				return       // 结束 keepalive 协程
			}
			idleTimer.Reset(val) // 重置定时器为剩余空闲时间
		case <-t.done: // 如果收到连接结束的通知
			return // 结束 keepalive 协程
		}
	}
}

// Close 主动关闭客户端连接
func (t *Conn) Close() error {
	//确保通道只关闭一次，避免重复关闭通道可能导致的 panic
	select {
	case <-t.done: // 如果 t.done 通道已经关闭（即已经有值可以接收），什么都不做
	default:
		close(t.done) // 如果 t.done 通道还没有关闭，关闭它
	}
	return t.Conn.Close()
}

func (t *Conn) ErrMsgData(err error, errCode ...int) any {
	code := resd.ErrSys
	danErr, ok := resd.AssertErr(err)
	if ok {
		if len(errCode) > 0 {
			code = errCode[0]
		} else {
			code = danErr.Code
		}
		return map[string]any{
			"code": code,
			"msg":  t.Resd.Msg(danErr.Code, danErr.GetTemps()...),
		}
	} else {
		if len(errCode) > 0 {
			code = errCode[0]
		}
		return map[string]any{
			"code": code,
			"msg":  err.Error(),
		}
	}
}
