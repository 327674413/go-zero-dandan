package websocketd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	sync.RWMutex          //读写锁，允许读并发，虽然开销大一点，当绝大多数是读操作时，但比互斥锁Mutex性能更好
	*threading.TaskRunner //并发处理
	opt                   *serverOption
	authenication         Authentication
	routes                map[string]HandlerFunc
	addr                  string
	upgrader              websocket.Upgrader
	logx.Logger
	connToUser map[*Conn]string
	userToConn map[string]*Conn
	patten     string
}

func NewServer(addr string, opts ...ServerOptions) *Server {
	opt := newServerOptions(opts...)
	return &Server{
		routes:        make(map[string]HandlerFunc),
		addr:          addr,
		patten:        opt.patten,
		opt:           &opt,
		upgrader:      websocket.Upgrader{},
		Logger:        logx.WithContext(context.Background()),
		connToUser:    make(map[*Conn]string),
		userToConn:    make(map[string]*Conn),
		authenication: opt.Authentication,
		TaskRunner:    threading.NewTaskRunner(opt.concurrency),
	}
}
func (t *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("server handler ws recover err：%v", r)
		}
	}()
	conn := NewConn(t, w, r)
	if conn == nil {
		return
	}
	if !t.authenication.Auth(w, r) {
		//提示权限
		t.Send(&Message{FrameType: FrameData, Data: "不具备访问权限"}, conn)
		//主动断开链接
		conn.Close()
		return
	}
	//记录链接
	t.addConn(conn, r)
	//处理链接
	go t.handlerConn(conn)
}
func (t *Server) addConn(conn *Conn, req *http.Request) {
	uid := t.authenication.UserId(req)
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()
	//判断用户是否已经登陆过
	if c := t.userToConn[uid]; c != nil {
		//先不做支持重复登陆，关闭之前的链接
		logx.Infof("%s用户已登陆，踢下线", uid)
		c.Close()
	}
	logx.Infof("%s用户连接成功", uid)
	t.connToUser[conn] = uid
	t.userToConn[uid] = conn
	logx.Info("当前总连接数量：", len(t.userToConn))
}
func (t *Server) GetConn(uid string) *Conn {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()
	return t.userToConn[uid]
}
func (t *Server) GetConns(uids ...string) []*Conn {
	if len(uids) == 0 {
		return nil
	}
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()
	res := make([]*Conn, 0, len(uids))
	for _, uid := range uids {
		res = append(res, t.userToConn[uid])
	}
	return res
}
func (t *Server) GetUser(conn *Conn) string {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()
	return t.connToUser[conn]
}
func (t *Server) GetUsers(conns ...*Conn) []string {
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()
	var res []string
	if len(conns) == 0 {
		res = make([]string, 0, len(t.userToConn))
		for _, uid := range t.connToUser {
			res = append(res, uid)
		}
	} else {
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, t.connToUser[conn])
		}
	}
	return res
}
func (t *Server) Close(conn *Conn) {
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()

	uid := t.connToUser[conn]
	if uid == "" {
		// 已被关闭了
		return
	}
	delete(t.connToUser, conn)
	delete(t.userToConn, uid)
	conn.Close()
}

// 根据链接对象进行任务处理
func (t *Server) handlerConn(conn *Conn) {
	//获取用户id
	conn.Uid = t.GetUser(conn)
	// 写入处理任务
	go t.handlerWrite(conn)
	//判断是否开启ack
	if t.isAck(nil) {
		go t.readAck(conn)
	}
	for {
		_, msg, err := conn.ReadMessage()
		//能确定，用int64精度丢失问题，在这里不管怎么解析，解析出来的都是错的数字
		logx.Info("消息来了")
		if err != nil {
			logx.Errorf("websocket conn read message err：%v", err)
			t.Close(conn)
			return
		}
		//消息转化
		var message *Message
		if err = json.Unmarshal(msg, &message); err != nil {
			logx.Errorf("json unmarshal err：%v ，msg：%s", err, string(msg))
			t.Close(conn)
			return
		}
		//根据消息机制进行处理
		if t.isAck(message) {
			logx.Infof("ACK机制的消息来了 %v", message)
			//进行ack处理机制，写入处理的队列
			conn.appendMsgMq(message)
		} else {
			//非ack，直接放入消息推送的通道，发送消息
			conn.sendMessageChan <- message
		}

	}
}
func (t *Server) isAck(message *Message) bool {
	if message == nil {
		return t.opt.ack != NoAck
	}
	return t.opt.ack != NoAck && message.FrameType != FrameNoAck
}

// 读取消息ack处理的写协程
func (t *Server) readAck(conn *Conn) {
	//处理发送失败场景
	send := func(msg *Message, conn *Conn) error {
		err := t.Send(msg, conn)
		if err == nil {
			return nil
		}
		t.Errorf("message ack  send err：%v ,msg：", err, msg)
		conn.readMessageAckMq[0].errCount++
		conn.messageMu.Unlock()
		tempDelay := time.Duration(200*conn.readMessageAckMq[0].errCount) * time.Microsecond
		if max := 1 * time.Second; tempDelay > max {
			tempDelay = max
		}
		time.Sleep(tempDelay)
		return err
	}
	for {
		select {
		case <-conn.done:
			t.Infof("close message ack uid %v", conn.Uid)
			return
		default:

		}

		//非阻塞的，没关闭前，会一直执行下面的代码
		conn.messageMu.Lock()
		//判断是否有待处理的ack消息
		if len(conn.readMessageAckMq) == 0 {
			//没有需要处理ack的消息，增加睡眠，确保协程工作切换
			conn.messageMu.Unlock()
			time.Sleep(100 * time.Microsecond)
			continue
		}
		// 读取ack队列第一条
		ackMessage := conn.readMessageAckMq[0]
		if ackMessage.errCount > t.opt.sendErrCount {
			logx.Errorf("conn send fail,message:%v,ackType:%v,maxSendErrCount:%v", ackMessage, t.opt.ack.ToString(), t.opt.sendErrCount)
			conn.messageMu.Unlock()
			//因为多次发送错误，这里选择放弃消息，如果有业务需要处理则再调整，比如加钩子之类的
			delete(conn.readMessageSeqMap, ackMessage.Id)
			conn.readMessageAckMq = conn.readMessageAckMq[1:]
			continue
		}
		//判断当前ack方式
		switch t.opt.ack {
		case OnlyAck:
			logx.Info("单次ack机制处理，则直接确认，发送消息")
			//单次ack确认，即可以发送成功，回复该发送者，服务器已处理消息
			if err := send(&Message{
				Id:        ackMessage.Id,
				FrameType: FrameAck,
				AckSeq:    ackMessage.AckSeq + 1,
			}, conn); err != nil {
				continue
			}
			// 将待处理ack消息移除
			conn.readMessageAckMq = conn.readMessageAckMq[1:]
			conn.messageMu.Unlock()
			// 将消息放入发送队列，发送给接收者
			conn.sendMessageChan <- ackMessage
		case RigorAck:
			//应答模式
			//判断是否是未确认过的消息
			if ackMessage.AckSeq == 0 {
				//序号为0则，还未确认过，服务端进行序号递增，回复给发送着
				conn.readMessageAckMq[0].AckSeq++
				conn.readMessageAckMq[0].ackTime = time.Now()
				t.Send(&Message{
					Id:        ackMessage.Id,
					FrameType: FrameAck,
					AckSeq:    ackMessage.AckSeq,
				}, conn)
				logx.Infof("应答模式消息首次确认，消息id：%v, seq %v,time %v", ackMessage.Id, ackMessage.AckSeq, ackMessage.ackTime)
				conn.messageMu.Unlock()
				continue
			}
			//序号不是0，代表是确认过的消息，进行再次确认
			msgCurrSeq := conn.readMessageSeqMap[ackMessage.Id] //这里不会出现key不存在的异常请求吗？
			// ack队列中的序号是否大于待处理消息队列中的序号
			if msgCurrSeq.AckSeq > ackMessage.AckSeq {
				//大于，则本次收到的消息序号更大，确认成功，发送消息
				conn.readMessageAckMq = conn.readMessageAckMq[1:]
				conn.messageMu.Unlock()
				conn.sendMessageChan <- ackMessage
				logx.Info("应答机制消息确认成功 mid: %v", ackMessage.Id)
				continue
			}
			//序号没有大于，则客户端还没有确认
			//允许的超时时间 减去 距离上一次确认的时间
			val := t.opt.ackTimeout - time.Since(ackMessage.ackTime)
			// 上次确认时间不为空，且最大允许超时时间小，则超时
			if !ackMessage.ackTime.IsZero() && val <= 0 {
				//超时了，暂时先选择放弃消息， 如果实际业务有需要特殊粗粝，再考虑加钩子之类的
				delete(conn.readMessageSeqMap, ackMessage.Id)
				conn.readMessageAckMq = conn.readMessageAckMq[1:]
				conn.messageMu.Unlock()
				logx.Info("ack确认超时，不再重发")
				continue

			}
			//未超过，重新发送
			conn.messageMu.Unlock()
			logx.Info("ack未确认，且没超时，重发")
			t.Send(&Message{FrameType: FrameAck, AckSeq: ackMessage.AckSeq, Id: ackMessage.Id}, conn)
			//睡眠一定时间
			time.Sleep(3 * time.Second)
		}
	}
}

// ack处理
func (t *Server) handlerWrite(conn *Conn) {
	for {
		select {
		case <-conn.done:
			//连接关闭
			return
		case message := <-conn.sendMessageChan:
			switch message.FrameType {
			case FramePing:
				t.Send(&Message{FrameType: FramePing}, conn)
			case FrameData:
				//根据请求的method分发路由并执行
				if handler, ok := t.routes[message.Method]; ok {
					handler(t, conn, message)
				} else {
					t.Send(&Message{FrameType: FrameData, Data: fmt.Sprintf("不存在的ws写入方法：%v", message.Method)}, conn)
				}
			}
			if t.isAck(message) {
				//消息发送成功后，如果是ack机制的，需要移除ack序号的map
				conn.messageMu.Lock()
				delete(conn.readMessageSeqMap, message.Id)
				conn.messageMu.Unlock()
			}
		}
	}
}
func (t *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		t.routes[r.Method] = r.Handler
	}
}
func (t *Server) Start() {
	http.HandleFunc(t.patten, t.ServerWs)
	t.Info(http.ListenAndServe(t.addr, nil))
}
func (t *Server) Stop() {
	fmt.Println("停止服务")
}

func (t *Server) SendByUserId(msg interface{}, userIds ...string) error {
	if len(userIds) == 0 {
		return nil
	}
	return t.Send(msg, t.GetConns(userIds...)...)
}
func (t *Server) Send(msg interface{}, conns ...*Conn) error {
	if len(conns) == 0 {
		return nil
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}
	return nil
}
