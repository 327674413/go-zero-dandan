package types

type Message struct {
	FromId   int64  //发送者
	ToUserId int64  //接收者
	Type     string //发送类型：私聊、群聊、广播
	Media    int    //消息类型：文字、图片、因屏
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他统计

}
