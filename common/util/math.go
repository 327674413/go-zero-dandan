package utils

import (
	"math/rand"
	"time"
)

func Rand(min int, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min+1) + min
}
