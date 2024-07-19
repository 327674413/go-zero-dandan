package websocketd

import (
	"time"
)

type FrameType uint8

const (
	FrameData      FrameType = 0x0
	FramePing      FrameType = 0x1
	FrameAck       FrameType = 0x2
	FrameNoAck     FrameType = 0x3
	FrameErr       FrameType = 0x9
	FrameTranspond FrameType = 0x6
)

type Message struct {
	Id        string             `json:"id"` // 消息ID
	FrameType `json:"frameType"` // 数据帧类型
	AckSeq    int                `json:"ackSeq"`   //ack确认
	ackTime   time.Time          `json:"ackTime"`  //ack时间
	errCount  int                `json:"errCount"` //错误次数，用于重发
	Method    string             `json:"method"`   //消息执行的方法
	FromUid   string             `json:"fromUid"`  //发送人
	Data      any                `json:"data"`     //消息具体内容
}

func NewMessage(FromUid string, data interface{}) *Message {
	return &Message{
		FrameType: FrameData,
		FromUid:   FromUid,
		Data:      data,
	}
}
func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameErr,
		Data:      err.Error(),
	}
}
