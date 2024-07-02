package main

import (
	"go-zero-dandan/common/fmtd"
)

const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorReset   = "\033[0m"
)

func main() {
	//fmt.Println(ColorRed + "红色文字" + ColorReset)
	fmtd.Info("aaa")
}
