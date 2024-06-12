package websocketd

import "time"

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
	Id        string `json:"id"`
	FrameType `json:"frameType"`
	AckSeq    int       `json:"ackSeq"`
	ackTime   time.Time `json:"ackTime"`
	errCount  int       `json:"errCount"`
	Method    string    `json:"method"`
	FormCode  string    `json:"formCode"`
	Data      any       `json:"data"`
}

func NewMessage(formCode string, data interface{}) *Message {
	return &Message{
		FrameType: FrameData,
		FormCode:  formCode,
		Data:      data,
	}
}
func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameErr,
		Data:      err.Error(),
	}
}
