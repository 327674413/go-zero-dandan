package types

type Message struct {
	FromUserId int64 //发送人
	ToUserId   int64 //接收人
	Type       string
	Media      int
	Content    string
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他统计

}
