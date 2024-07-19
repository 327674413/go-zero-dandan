package websocketd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"net/http"
	"runtime/debug"
	"sync"
	"time"
)

// Server websocket核心服务，管理路由、连接信息等
type Server struct {
	sync.RWMutex                                 //读写锁，允许读并发，虽然开销大一点，当绝大多数是读操作时，但比互斥锁Mutex性能更好
	*threading.TaskRunner                        //zero封装的，多协程并发执行的方法
	opt                   *serverOption          //服务启动的配置参数
	authenication         Authentication         //用户连接的授权接口，要实现根据http请求获取，返回是否允许执行和获取用户id方法
	routes                map[string]HandlerFunc //方法路由，通过请求中的method方法来调用具体路由对应的执行方法
	addr                  string                 // ws启动的端口地址
	upgrader              websocket.Upgrader     //将http升级到socket协议使用
	logx.Logger                                  // zero封装的日志方法
	connToUser            map[*Conn]string       // 管理连接，以conn连接为key
	userToConn            map[string]*Conn       // 管理连接，以用户id为key
	patten                string                 //监听路径，也就是地址:ip后面部分的路径地址，如这里实际用的是/ws
}

func NewServer(addr string, opts ...ServerOptions) *Server {
	opt := newServerOptions(opts...)
	return &Server{
		routes: make(map[string]HandlerFunc),
		addr:   addr,
		patten: opt.patten,
		opt:    &opt,
		upgrader: websocket.Upgrader{
			//设置允许跨域
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Logger:        logx.WithContext(context.Background()),
		connToUser:    make(map[*Conn]string),
		userToConn:    make(map[string]*Conn),
		authenication: opt.Authentication,
		TaskRunner:    threading.NewTaskRunner(opt.concurrency),
	}
}

// ServerWs 处理ws请求的方法，调用时通过http.HandleFunc(patten,ServerWs)完成服务绑定和启动
func (t *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	// 捕获ws服务中的panic，避免程序中断
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("server handler ws recover err：%v \n,%s", r, debug.Stack())
		}
	}()
	//将用户的http请求连接升级为websocket的conn连接
	conn := NewConn(t, w, r)
	if conn == nil {
		return
	}
	//校验用户的连接权限，比如是否登陆用户等
	if err := t.authenication.Auth(w, r); err != nil {
		//无权限
		t.Send(&Message{FrameType: FrameErr, Data: conn.ErrMsgData(err)}, conn)
		//主动断开链接
		conn.Close()
		return
	}
	//记录链接
	t.addConn(conn, r)
	//通过协程，开启每个连接的消息处理
	go t.handlerConn(conn)
}
func (t *Server) addConn(conn *Conn, req *http.Request) {
	//根据授权，获取用户id
	uid := t.authenication.UserId(req)
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()
	//判断用户是否已经登陆过
	if c := t.userToConn[uid]; c != nil {
		//先不做支持重复登陆，关闭之前的链接
		// todo::支持多设备登陆使用
		logx.Infof("用户%s已登陆，踢下线", uid)
		c.Close()
	}
	logx.Infof("用户%s连接成功", uid)
	t.connToUser[conn] = uid
	conn.Uid = uid
	t.userToConn[uid] = conn
	logx.Info("当前总连接数量：", len(t.userToConn))
	//todo::需要发送上线消息给kafka，更新上线用户缓存和给好友发上线通知
}

// GetConn 根据用户id获取conn连接
func (t *Server) GetConn(uid string) *Conn {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()
	return t.userToConn[uid]
}

// GetConns 根据多个用户id，获取多个conn连接
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

// GetUser 根据conn连接获取用户id
func (t *Server) GetUser(conn *Conn) string {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()
	return t.connToUser[conn]
}

// GetUsers 根据多个conn连接获取多个用户id
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

// Close 关闭conn连接
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

// handlerConn 每个连接的ws处理方法，一个conn就会有一个协程
func (t *Server) handlerConn(conn *Conn) {
	// 专门再起一个写消息的协程，本方法金接收消息，判断ack或直接处理
	go t.messageHandler(conn)
	// 根据是否启用ack，开启ack的处理协层，传nil就会走启动服务时的opt参数
	if t.isAck(nil) {
		go t.readAck(conn)
	}
	// 阻塞，直到有客户端法消息进来才会开始处理消息
	for {
		//这里读到消息后，并没有直接处理业务，而是判断类型后，放到ack待确认里，或放到管道中，专门另一个协程来处理
		_, msg, err := conn.ReadMessage()
		//备注：如果msg中的数据有int64那种很长的数字，似乎会出现精度丢失问题，在这里收到就和发出的不一样了，所以最好长数字要用字符串传输
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				t.Infof("用户%s下线", conn.Uid)
			} else {
				t.Infof("用户%s离线", conn.Uid)
			}
			t.Close(conn)
			return
		}
		t.Infof("用户%s消息来了", conn.Uid)
		//消息转化，json解析成消息结构体
		var message *Message
		if err = json.Unmarshal(msg, &message); err != nil {
			t.Errorf("消息json解析失败：%v ，msg：%s", err, string(msg))
			t.Close(conn)
			return
		}
		if message.FrameType == FramePing {
			//如果是ping消息，放置客户端被断开用，可不回复
			continue
		}
		//根据消息机制进行处理
		if t.isAck(message) {
			t.Infof("ACK机制的消息来了 %v", message)
			//进行ack处理机制，写入待处理的队列，因为每次只处理一条，直接用数组，不确定用管道会不会更好
			conn.appendMsgMq(message)
		} else {
			//非ack，直接放入消息发送的管道，该管道有写消息的协程执行，会解析消息内容进行发送
			conn.readMessageChan <- message
		}

	}
}

func (t *Server) isAck(message *Message) bool {
	if message == nil {
		return t.opt.ack != AckTypeNoAck
	}
	return t.opt.ack != AckTypeNoAck && message.FrameType != FrameNoAck
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
		case AckTypeOnlyAck:
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
			conn.readMessageChan <- ackMessage
		case AckTypeRigorAck:
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
				conn.readMessageChan <- ackMessage
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

// messageHandler 处理收到消息的协程方法
func (t *Server) messageHandler(conn *Conn) {
	for {
		select {
		case <-conn.done:
			//主动关闭调用conn.Close时会关闭，触发进入该方法，然后结束这个conn的写协程
			return
		case message := <-conn.readMessageChan:
			//无需ack或已经确认完的消息进来，进行消息处理
			switch message.FrameType {
			case FramePing:
				//过来的是ping类型的消息，这里好像没有回复gong，也是给该客户端回复的ping？？？还没搞懂是不是通过ping来做心跳
				t.Send(&Message{FrameType: FramePing}, conn)
			case FrameData:
				//data类型的消息，根据消息中的method参数，找到对应的路由.方法，并执行
				if handler, ok := t.routes[message.Method]; ok {
					handler(t, conn, message)
				} else {
					//未提供方法时，报错
					t.Send(&Message{FrameType: FrameData, Data: fmt.Sprintf("不存在的ws写入方法：%v", message.Method)}, conn)
				}
			}
			if t.isAck(message) {
				//消息读取处理完毕后，如果是ack机制的，需要移除ack序号的map
				conn.messageMu.Lock()
				delete(conn.readMessageSeqMap, message.Id)
				conn.messageMu.Unlock()
			}
		}
	}
}

// AddRoutes 注册路由方法
func (t *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		t.routes[r.Method] = r.Handler
	}
}

// Start 启动ws的连接服务
func (t *Server) Start() {
	http.HandleFunc(t.patten, t.ServerWs)
	t.Info(http.ListenAndServe(t.addr, nil))
}

// Stop 停止ws的连接服务，暂未实现
func (t *Server) Stop() {
	fmt.Println("停止服务")
}

// SendByUserId 根据用户id发送消息
func (t *Server) SendByUserId(msg interface{}, userIds ...string) error {
	if len(userIds) == 0 {
		return nil
	}
	return t.Send(msg, t.GetConns(userIds...)...)
}

// Send 根据conn连接发送消息
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
