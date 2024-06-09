package websocketd

import (
	"math"
	"time"
)

const (
	defaultMaxConnectionIdle = time.Duration(math.MaxInt64) // 最大允许的连接空闲时间
	defaultAckTimeout        = 30 * time.Second
	defaultSendErrCount      = 3 //发送失败次数
)
