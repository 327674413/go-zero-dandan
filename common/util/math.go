package util

import (
	"math/rand"
	"time"
)

// globalRand 使用了一个全局变量 globalRand 来保存随机数生成器，这样可以确保在程序运行期间只有一个随机数生成器实例存在，并且多个协程可以共享这个实例，保证了并发安全
var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Rand(min, max int) int {
	return globalRand.Intn(max-min+1) + min
}
