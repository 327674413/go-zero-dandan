package wsTypes

type Message struct {
	FromId   int64  //发送者
	ToUserId int64  //接收者
	Type     string //发送类型：私聊、群聊、广播、红包、视频、语音、抖动
	Media    int    //消息类型：文字、图片、语音、视频、文件、导航地址、h5分享链接
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他统计

}

type ConnectionReq struct {
	Token      string `form:"token"`
	PlatformEm int64  `form:"platformEm"`
}

type ConnectionRes struct {
	UserId int64 `json:"userId,string"`
}
